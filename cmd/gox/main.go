package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/buemura/gox/pkg/compiler"
	"github.com/buemura/gox/pkg/formatter"
	"github.com/buemura/gox/pkg/lsp"
	"github.com/buemura/gox/pkg/parser"
	"github.com/buemura/gox/pkg/proxy"
	"github.com/buemura/gox/pkg/watcher"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generate":
		if err := runGenerate(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "watch":
		if err := runWatch(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "fmt":
		if err := runFmt(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "lsp":
		if err := runLSP(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage: gox <command> [flags]

Commands:
  generate    Compile .gox files into Go source code
  watch       Watch .gox files and recompile on changes
  fmt         Format .gox files
  lsp         Start the Language Server Protocol server

Watch flags:
  -proxy      Upstream URL to proxy (e.g. http://localhost:8080)
  -proxyport  Port for the hot-reload proxy (default: 8081)
  -cmd        Command to run/restart after recompilation (e.g. "go run .")
`)
}

func runLSP() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	return lsp.Serve(ctx, os.Stdin, os.Stdout)
}

func runWatch(args []string) error {
	fs := flag.NewFlagSet("watch", flag.ExitOnError)
	dir := fs.String("dir", ".", "root directory to watch")
	out := fs.String("out", "_gox.go", "output file suffix")
	proxyUpstream := fs.String("proxy", "", "upstream URL to proxy (e.g. http://localhost:8080)")
	proxyPort := fs.String("proxyport", "8081", "port for the hot-reload proxy server")
	cmd := fs.String("cmd", "", "command to run/restart after recompilation (e.g. \"go run .\")")
	fs.Parse(args)

	// Initial full generate.
	if err := runGenerate([]string{"-dir", *dir, "-out", *out}); err != nil {
		fmt.Fprintf(os.Stderr, "initial generate: %v\n", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg := watcher.Config{
		Root:      *dir,
		OutSuffix: *out,
		Debounce:  100 * time.Millisecond,
	}

	// Set up hot-reload proxy if --proxy is specified.
	if *proxyUpstream != "" {
		srv, err := proxy.New(*proxyUpstream, ":"+*proxyPort)
		if err != nil {
			return fmt.Errorf("creating proxy: %w", err)
		}
		cfg.OnReload = srv.Reload
		go func() {
			fmt.Printf("proxy listening on http://localhost:%s → %s\n", *proxyPort, *proxyUpstream)
			if err := srv.ListenAndServe(); err != nil {
				fmt.Fprintf(os.Stderr, "proxy error: %v\n", err)
			}
		}()
	}

	// Set up --cmd process management if specified.
	if *cmd != "" {
		var cmdProc *exec.Cmd
		prevOnReload := cfg.OnReload

		startCmd := func() {
			cmdProc = exec.CommandContext(ctx, "sh", "-c", *cmd)
			cmdProc.Stdout = os.Stdout
			cmdProc.Stderr = os.Stderr
			if err := cmdProc.Start(); err != nil {
				fmt.Fprintf(os.Stderr, "[cmd] failed to start: %v\n", err)
			} else {
				fmt.Fprintf(os.Stderr, "[cmd] started: %s (pid %d)\n", *cmd, cmdProc.Process.Pid)
			}
		}

		// Start the command initially.
		startCmd()

		cfg.OnReload = func() {
			// Kill previous process if running.
			if cmdProc != nil && cmdProc.Process != nil {
				cmdProc.Process.Kill()
				cmdProc.Wait()
			}
			// Restart the command.
			startCmd()
			// Also trigger proxy reload if configured.
			if prevOnReload != nil {
				prevOnReload()
			}
		}
	}

	w, err := watcher.New(cfg, compileFile)
	if err != nil {
		return err
	}

	fmt.Printf("watching %s for .gox changes...\n", *dir)
	return w.Watch(ctx)
}

func runFmt(args []string) error {
	fs := flag.NewFlagSet("fmt", flag.ExitOnError)
	dir := fs.String("dir", ".", "root directory to search")
	stdout := fs.Bool("stdout", false, "print to stdout instead of writing in place")
	fs.Parse(args)

	files, err := findGoxFiles(*dir)
	if err != nil {
		return fmt.Errorf("finding .gox files: %w", err)
	}

	if len(files) == 0 {
		fmt.Println("no .gox files found")
		return nil
	}

	for _, f := range files {
		src, err := os.ReadFile(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading %s: %v\n", f, err)
			continue
		}

		formatted, err := formatter.FormatFile(f, string(src))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error formatting %s: %v\n", f, err)
			continue
		}

		if *stdout {
			fmt.Print(formatted)
		} else {
			if err := os.WriteFile(f, []byte(formatted), 0644); err != nil {
				fmt.Fprintf(os.Stderr, "error writing %s: %v\n", f, err)
				continue
			}
			fmt.Printf("  formatted %s\n", f)
		}
	}
	return nil
}

func runGenerate(args []string) error {
	fs := flag.NewFlagSet("generate", flag.ExitOnError)
	dir := fs.String("dir", ".", "root directory to search for .gox files")
	out := fs.String("out", "_gox.go", "output file suffix (replaces .gox extension)")
	fs.Parse(args)

	files, err := findGoxFiles(*dir)
	if err != nil {
		return fmt.Errorf("finding .gox files: %w", err)
	}

	if len(files) == 0 {
		fmt.Println("no .gox files found")
		return nil
	}

	var errs []string
	for _, f := range files {
		if err := compileFile(f, *out); err != nil {
			errs = append(errs, err.Error())
		}
	}

	fmt.Printf("compiled %d file(s)\n", len(files)-len(errs))

	if len(errs) > 0 {
		return fmt.Errorf("failed to compile %d file(s):\n%s", len(errs), strings.Join(errs, "\n"))
	}

	return nil
}

func findGoxFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".gox" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func compileFile(path, outSuffix string) error {
	src, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}

	p := parser.NewParser(string(src))
	ast, err := p.Parse()
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}

	output, err := compiler.Compile(ast)
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}

	outPath := strings.TrimSuffix(path, ".gox") + outSuffix
	if err := os.WriteFile(outPath, []byte(output), 0644); err != nil {
		return fmt.Errorf("%s: writing output: %w", path, err)
	}

	fmt.Printf("  %s → %s\n", path, outPath)
	return nil
}
