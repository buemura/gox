package watcher

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Config holds watcher configuration.
type Config struct {
	Root      string
	OutSuffix string
	Debounce  time.Duration
}

// Watcher watches for .gox file changes and triggers recompilation.
type Watcher struct {
	cfg     Config
	fsw     *fsnotify.Watcher
	compile func(path, outSuffix string) error

	mu     sync.Mutex
	timers map[string]*time.Timer
	events chan string
}

// New creates a Watcher. compileFn is called when a .gox file changes.
func New(cfg Config, compileFn func(path, outSuffix string) error) (*Watcher, error) {
	if cfg.Debounce == 0 {
		cfg.Debounce = 100 * time.Millisecond
	}
	if cfg.OutSuffix == "" {
		cfg.OutSuffix = "_gox.go"
	}

	fsw, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("creating watcher: %w", err)
	}

	w := &Watcher{
		cfg:     cfg,
		fsw:     fsw,
		compile: compileFn,
		timers:  make(map[string]*time.Timer),
		events:  make(chan string, 64),
	}

	if err := w.addRecursive(cfg.Root); err != nil {
		fsw.Close()
		return nil, fmt.Errorf("watching directories: %w", err)
	}

	return w, nil
}

// addRecursive adds all directories under root to the fsnotify watcher.
func (w *Watcher) addRecursive(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return w.fsw.Add(path)
		}
		return nil
	})
}

// isGoxFile returns true if the path is a .gox file (not a generated output).
func (w *Watcher) isGoxFile(path string) bool {
	return filepath.Ext(path) == ".gox" && !strings.HasSuffix(path, w.cfg.OutSuffix)
}

// outputPath returns the generated output path for a .gox file.
func (w *Watcher) outputPath(goxPath string) string {
	return strings.TrimSuffix(goxPath, ".gox") + w.cfg.OutSuffix
}

// debounce schedules a compile for the given path, resetting if already pending.
func (w *Watcher) debounce(path string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if t, ok := w.timers[path]; ok {
		t.Reset(w.cfg.Debounce)
		return
	}

	w.timers[path] = time.AfterFunc(w.cfg.Debounce, func() {
		w.events <- path
		w.mu.Lock()
		delete(w.timers, path)
		w.mu.Unlock()
	})
}

// Watch starts watching and blocks until ctx is cancelled.
func (w *Watcher) Watch(ctx context.Context) error {
	defer w.fsw.Close()

	for {
		select {
		case <-ctx.Done():
			return nil

		case evt, ok := <-w.fsw.Events:
			if !ok {
				return nil
			}
			w.handleEvent(evt)

		case err, ok := <-w.fsw.Errors:
			if !ok {
				return nil
			}
			fmt.Fprintf(os.Stderr, "[watch] error: %v\n", err)

		case path, ok := <-w.events:
			if !ok {
				return nil
			}
			w.compileAndLog(path)
		}
	}
}

// handleEvent processes a single fsnotify event.
func (w *Watcher) handleEvent(evt fsnotify.Event) {
	path := evt.Name

	// Watch new directories.
	if evt.Has(fsnotify.Create) {
		if info, err := os.Stat(path); err == nil && info.IsDir() {
			w.fsw.Add(path)
		}
	}

	if !w.isGoxFile(path) {
		return
	}

	if evt.Has(fsnotify.Remove) || evt.Has(fsnotify.Rename) {
		out := w.outputPath(path)
		if err := os.Remove(out); err == nil {
			fmt.Fprintf(os.Stderr, "[watch] deleted %s\n", out)
		}
		return
	}

	if evt.Has(fsnotify.Create) || evt.Has(fsnotify.Write) {
		w.debounce(path)
	}
}

// compileAndLog runs the compile function and logs the result.
func (w *Watcher) compileAndLog(path string) {
	if err := w.compile(path, w.cfg.OutSuffix); err != nil {
		fmt.Fprintf(os.Stderr, "[watch] error: %v\n", err)
	} else {
		fmt.Fprintf(os.Stderr, "[watch] compiled %s\n", path)
	}
}
