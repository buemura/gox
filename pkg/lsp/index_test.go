package lsp

import (
	"os"
	"path/filepath"
	"testing"
)

func TestComponentIndex_IndexFileContent(t *testing.T) {
	idx := NewComponentIndex()
	content := `package views

func Header(title string) {
	<h1>{{ title }}</h1>
}

func Footer() {
	<footer>Footer</footer>
}
`
	idx.IndexFileContent("file:///test/views.gox", content)

	header, ok := idx.Lookup("Header")
	if !ok {
		t.Fatal("expected Header component in index")
	}
	if header.Params != "title string" {
		t.Errorf("expected params 'title string', got %q", header.Params)
	}
	if header.URI != "file:///test/views.gox" {
		t.Errorf("expected URI 'file:///test/views.gox', got %q", header.URI)
	}

	footer, ok := idx.Lookup("Footer")
	if !ok {
		t.Fatal("expected Footer component in index")
	}
	if footer.Params != "" {
		t.Errorf("expected empty params, got %q", footer.Params)
	}
}

func TestComponentIndex_UpdateRemovesOldEntries(t *testing.T) {
	idx := NewComponentIndex()
	uri := "file:///test/views.gox"

	idx.IndexFileContent(uri, `package views
func OldComponent() {
	<div>old</div>
}
`)
	if _, ok := idx.Lookup("OldComponent"); !ok {
		t.Fatal("expected OldComponent in index")
	}

	// Re-index same file with different component.
	idx.IndexFileContent(uri, `package views
func NewComponent() {
	<div>new</div>
}
`)
	if _, ok := idx.Lookup("OldComponent"); ok {
		t.Error("OldComponent should have been removed after re-index")
	}
	if _, ok := idx.Lookup("NewComponent"); !ok {
		t.Error("expected NewComponent in index after re-index")
	}
}

func TestComponentIndex_Scan(t *testing.T) {
	dir := t.TempDir()
	content := `package views

func MyComponent() {
	<div>hello</div>
}
`
	err := os.WriteFile(filepath.Join(dir, "test.gox"), []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	idx := NewComponentIndex()
	idx.Scan(dir)

	if _, ok := idx.Lookup("MyComponent"); !ok {
		t.Fatal("expected MyComponent in index after scan")
	}
}

func TestComponentIndex_All(t *testing.T) {
	idx := NewComponentIndex()
	idx.IndexFileContent("file:///a.gox", `package views
func A() {
	<div>a</div>
}
`)
	idx.IndexFileContent("file:///b.gox", `package views
func B() {
	<div>b</div>
}
`)
	all := idx.All()
	if len(all) != 2 {
		t.Errorf("expected 2 components, got %d", len(all))
	}
}

func TestComponentIndex_InvalidFile(t *testing.T) {
	idx := NewComponentIndex()
	// Invalid gox content should not panic.
	idx.IndexFileContent("file:///bad.gox", `this is not valid gox`)
	all := idx.All()
	if len(all) != 0 {
		t.Errorf("expected 0 components for invalid file, got %d", len(all))
	}
}
