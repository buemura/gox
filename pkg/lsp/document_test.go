package lsp

import "testing"

func TestDocumentStore(t *testing.T) {
	store := NewDocumentStore()

	// Open.
	store.Open("file:///test.gox", "hello")
	content, ok := store.Get("file:///test.gox")
	if !ok || content != "hello" {
		t.Errorf("expected 'hello', got %q (ok=%v)", content, ok)
	}

	// Update.
	store.Update("file:///test.gox", "world")
	content, ok = store.Get("file:///test.gox")
	if !ok || content != "world" {
		t.Errorf("expected 'world', got %q (ok=%v)", content, ok)
	}

	// Close.
	store.Close("file:///test.gox")
	_, ok = store.Get("file:///test.gox")
	if ok {
		t.Error("expected document to be removed after close")
	}
}

func TestDocumentStore_GetNonexistent(t *testing.T) {
	store := NewDocumentStore()
	_, ok := store.Get("file:///nonexistent.gox")
	if ok {
		t.Error("expected ok=false for nonexistent document")
	}
}
