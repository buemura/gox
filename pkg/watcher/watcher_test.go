package watcher

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

func TestWatcherCompileOnCreate(t *testing.T) {
	dir := t.TempDir()

	var mu sync.Mutex
	compiled := map[string]bool{}

	compileFn := func(path, outSuffix string) error {
		mu.Lock()
		compiled[path] = true
		mu.Unlock()
		return nil
	}

	cfg := Config{
		Root:      dir,
		OutSuffix: "_gox.go",
		Debounce:  10 * time.Millisecond,
	}

	w, err := New(cfg, compileFn)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go w.Watch(ctx)

	// Create a .gox file.
	goxPath := filepath.Join(dir, "hello.gox")
	if err := os.WriteFile(goxPath, []byte("package views"), 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	// Wait for debounce + processing.
	time.Sleep(200 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()
	if !compiled[goxPath] {
		t.Errorf("expected %s to be compiled", goxPath)
	}
}

func TestWatcherIgnoresNonGoxFiles(t *testing.T) {
	dir := t.TempDir()

	var mu sync.Mutex
	compileCount := 0

	compileFn := func(path, outSuffix string) error {
		mu.Lock()
		compileCount++
		mu.Unlock()
		return nil
	}

	cfg := Config{
		Root:      dir,
		OutSuffix: "_gox.go",
		Debounce:  10 * time.Millisecond,
	}

	w, err := New(cfg, compileFn)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go w.Watch(ctx)

	// Create a non-.gox file.
	if err := os.WriteFile(filepath.Join(dir, "readme.md"), []byte("hello"), 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	time.Sleep(200 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()
	if compileCount != 0 {
		t.Errorf("expected 0 compiles, got %d", compileCount)
	}
}

func TestWatcherDebounce(t *testing.T) {
	dir := t.TempDir()

	var mu sync.Mutex
	compileCount := 0

	compileFn := func(path, outSuffix string) error {
		mu.Lock()
		compileCount++
		mu.Unlock()
		return nil
	}

	cfg := Config{
		Root:      dir,
		OutSuffix: "_gox.go",
		Debounce:  50 * time.Millisecond,
	}

	w, err := New(cfg, compileFn)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go w.Watch(ctx)

	goxPath := filepath.Join(dir, "hello.gox")

	// Write rapidly multiple times.
	for i := 0; i < 5; i++ {
		os.WriteFile(goxPath, []byte("package views"), 0644)
		time.Sleep(10 * time.Millisecond)
	}

	// Wait for debounce to settle.
	time.Sleep(200 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()
	// Debounce should collapse rapid writes into 1-2 compiles.
	if compileCount > 2 {
		t.Errorf("expected debounce to reduce compiles, got %d", compileCount)
	}
	if compileCount == 0 {
		t.Errorf("expected at least 1 compile")
	}
}

func TestWatcherDeleteCleansOutput(t *testing.T) {
	dir := t.TempDir()

	compileFn := func(path, outSuffix string) error { return nil }

	cfg := Config{
		Root:      dir,
		OutSuffix: "_gox.go",
		Debounce:  10 * time.Millisecond,
	}

	// Create a .gox file and its output.
	goxPath := filepath.Join(dir, "hello.gox")
	outPath := filepath.Join(dir, "hello_gox.go")
	os.WriteFile(goxPath, []byte("package views"), 0644)
	os.WriteFile(outPath, []byte("package views"), 0644)

	w, err := New(cfg, compileFn)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go w.Watch(ctx)

	// Delete the .gox file.
	os.Remove(goxPath)

	time.Sleep(200 * time.Millisecond)

	if _, err := os.Stat(outPath); !os.IsNotExist(err) {
		t.Errorf("expected output file %s to be deleted", outPath)
	}
}

func TestWatcherSubdirectory(t *testing.T) {
	dir := t.TempDir()

	var mu sync.Mutex
	compiled := map[string]bool{}

	compileFn := func(path, outSuffix string) error {
		mu.Lock()
		compiled[path] = true
		mu.Unlock()
		return nil
	}

	cfg := Config{
		Root:      dir,
		OutSuffix: "_gox.go",
		Debounce:  10 * time.Millisecond,
	}

	w, err := New(cfg, compileFn)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go w.Watch(ctx)

	// Create a new subdirectory and a .gox file inside it.
	subDir := filepath.Join(dir, "views")
	os.MkdirAll(subDir, 0755)
	time.Sleep(50 * time.Millisecond) // Let watcher pick up new dir.

	goxPath := filepath.Join(subDir, "hello.gox")
	os.WriteFile(goxPath, []byte("package views"), 0644)

	time.Sleep(200 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()
	if !compiled[goxPath] {
		t.Errorf("expected %s to be compiled", goxPath)
	}
}
