package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestEndToEnd_SimpleComponent(t *testing.T) {
	tmp := t.TempDir()

	// Write a .gox file
	goxContent := `package views

import "fmt"

func Hello(name string) {
  <h1>Hello, {{ name }}!</h1>
}
`
	goxPath := filepath.Join(tmp, "hello.gox")
	if err := os.WriteFile(goxPath, []byte(goxContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Run generate
	if err := runGenerate([]string{"-dir", tmp}); err != nil {
		t.Fatalf("generate failed: %v", err)
	}

	// Check output file exists
	outPath := filepath.Join(tmp, "hello_gox.go")
	content, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("output file not found: %v", err)
	}

	src := string(content)
	assertContains(t, src, "package views")
	assertContains(t, src, "func Hello(name string) gox.Component")
	assertContains(t, src, `io.WriteString`)
	assertContains(t, src, `html.EscapeString`)
}

func TestEndToEnd_IfElse(t *testing.T) {
	tmp := t.TempDir()

	goxContent := `package views

func Greeting(loggedIn bool, name string) {
  {{ if loggedIn }}
    <p>Welcome, {{ name }}!</p>
  {{ else }}
    <p>Please log in.</p>
  {{ end }}
}
`
	goxPath := filepath.Join(tmp, "greeting.gox")
	if err := os.WriteFile(goxPath, []byte(goxContent), 0644); err != nil {
		t.Fatal(err)
	}

	if err := runGenerate([]string{"-dir", tmp}); err != nil {
		t.Fatalf("generate failed: %v", err)
	}

	content, err := os.ReadFile(filepath.Join(tmp, "greeting_gox.go"))
	if err != nil {
		t.Fatalf("output file not found: %v", err)
	}

	src := string(content)
	assertContains(t, src, "if loggedIn")
	assertContains(t, src, "Welcome,")
	assertContains(t, src, "Please log in.")
}

func TestEndToEnd_ForLoop(t *testing.T) {
	tmp := t.TempDir()

	goxContent := `package views

func ItemList(items []string) {
  <ul>
    {{ for _, item := range items }}
      <li>{{ item }}</li>
    {{ end }}
  </ul>
}
`
	goxPath := filepath.Join(tmp, "list.gox")
	if err := os.WriteFile(goxPath, []byte(goxContent), 0644); err != nil {
		t.Fatal(err)
	}

	if err := runGenerate([]string{"-dir", tmp}); err != nil {
		t.Fatalf("generate failed: %v", err)
	}

	content, err := os.ReadFile(filepath.Join(tmp, "list_gox.go"))
	if err != nil {
		t.Fatalf("output file not found: %v", err)
	}

	src := string(content)
	assertContains(t, src, "for _, item := range items")
	assertContains(t, src, `"<li"`)
	assertContains(t, src, `"</li>"`)
	assertContains(t, src, `"</ul>"`)

}

func TestEndToEnd_CustomOutSuffix(t *testing.T) {
	tmp := t.TempDir()

	goxContent := `package views

func Simple() {
  <div>hello</div>
}
`
	goxPath := filepath.Join(tmp, "simple.gox")
	if err := os.WriteFile(goxPath, []byte(goxContent), 0644); err != nil {
		t.Fatal(err)
	}

	if err := runGenerate([]string{"-dir", tmp, "-out", ".gen.go"}); err != nil {
		t.Fatalf("generate failed: %v", err)
	}

	outPath := filepath.Join(tmp, "simple.gen.go")
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		t.Fatalf("expected output file %s to exist", outPath)
	}
}

func TestEndToEnd_NoGoxFiles(t *testing.T) {
	tmp := t.TempDir()

	// Should not error with no .gox files
	if err := runGenerate([]string{"-dir", tmp}); err != nil {
		t.Fatalf("expected no error for empty dir, got: %v", err)
	}
}

func TestEndToEnd_CompileAndRun(t *testing.T) {
	// Skip if go toolchain unavailable
	if _, err := exec.LookPath("go"); err != nil {
		t.Skip("go toolchain not available")
	}

	tmp := t.TempDir()

	// Create a Go module with a main package that uses a compiled .gox component
	if err := os.WriteFile(filepath.Join(tmp, "go.mod"), []byte("module testapp\n\ngo 1.24.4\n\nrequire github.com/buemura/gox v0.0.0\n"), 0644); err != nil {
		t.Fatal(err)
	}

	// Get the project root for the replace directive
	projectRoot, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}

	goSum := filepath.Join(tmp, "go.sum")
	os.WriteFile(goSum, []byte(""), 0644)

	// Add replace directive
	cmd := exec.Command("go", "mod", "edit", "-replace=github.com/buemura/gox="+projectRoot)
	cmd.Dir = tmp
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go mod edit failed: %v\n%s", err, out)
	}

	// Create views directory with .gox file
	viewsDir := filepath.Join(tmp, "views")
	os.MkdirAll(viewsDir, 0755)

	goxContent := `package views

func Hello(name string) {
  <h1>Hello, {{ name }}!</h1>
}
`
	if err := os.WriteFile(filepath.Join(viewsDir, "hello.gox"), []byte(goxContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Compile the .gox file
	if err := runGenerate([]string{"-dir", viewsDir}); err != nil {
		t.Fatalf("generate failed: %v", err)
	}

	// Create main.go that uses the compiled component
	mainContent := `package main

import (
	"fmt"
	"github.com/buemura/gox"
	"testapp/views"
)

func main() {
	s, err := gox.RenderToString(views.Hello("World"))
	if err != nil {
		panic(err)
	}
	fmt.Print(s)
}
`
	if err := os.WriteFile(filepath.Join(tmp, "main.go"), []byte(mainContent), 0644); err != nil {
		t.Fatal(err)
	}

	// go mod tidy
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = tmp
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go mod tidy failed: %v\n%s", err, out)
	}

	// Build and run
	cmd = exec.Command("go", "run", ".")
	cmd.Dir = tmp
	out, err := cmd.CombinedOutput()
	if err != nil {
		// Read the generated file for debugging
		gen, _ := os.ReadFile(filepath.Join(viewsDir, "hello_gox.go"))
		t.Fatalf("go run failed: %v\n%s\n\ngenerated code:\n%s", err, out, gen)
	}

	result := strings.TrimSpace(string(out))
	expected := "<h1>Hello, World!</h1>"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func assertContains(t *testing.T, s, substr string) {
	t.Helper()
	if !strings.Contains(s, substr) {
		t.Errorf("expected output to contain %q, but it didn't.\nFull output:\n%s", substr, s)
	}
}
