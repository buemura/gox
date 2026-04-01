package lsp

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/buemura/gox/pkg/parser"
)

// ComponentInfo holds metadata about a component declaration.
type ComponentInfo struct {
	Name   string
	Params string
	URI    string
	Line   int
	Col    int
}

// ComponentIndex maps component names to their declarations.
type ComponentIndex struct {
	mu         sync.RWMutex
	components map[string]ComponentInfo
}

// NewComponentIndex creates an empty component index.
func NewComponentIndex() *ComponentIndex {
	return &ComponentIndex{components: make(map[string]ComponentInfo)}
}

// Scan walks the root directory, parsing all .gox files and indexing their components.
func (idx *ComponentIndex) Scan(rootDir string) {
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".gox" {
			return nil
		}
		absPath, _ := filepath.Abs(path)
		uri := "file://" + absPath
		src, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		idx.IndexFileContent(uri, string(src))
		return nil
	})
}

// IndexFileContent parses a single file's content and indexes its components.
func (idx *ComponentIndex) IndexFileContent(uri, content string) {
	p := parser.NewParser(content)
	file, err := p.Parse()
	if err != nil || file == nil {
		return
	}
	idx.IndexFile(uri, file)
}

// IndexFile indexes all components from a parsed file.
func (idx *ComponentIndex) IndexFile(uri string, file *parser.File) {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	// Remove old entries for this URI.
	for name, info := range idx.components {
		if info.URI == uri {
			delete(idx.components, name)
		}
	}

	for _, comp := range file.Components {
		idx.components[comp.Name] = ComponentInfo{
			Name:   comp.Name,
			Params: comp.Params,
			URI:    uri,
			Line:   comp.Line,
			Col:    comp.Col,
		}
	}
}

// Lookup returns the component info for a given name, if it exists.
func (idx *ComponentIndex) Lookup(name string) (ComponentInfo, bool) {
	idx.mu.RLock()
	defer idx.mu.RUnlock()
	info, ok := idx.components[name]
	return info, ok
}

// All returns all indexed components.
func (idx *ComponentIndex) All() []ComponentInfo {
	idx.mu.RLock()
	defer idx.mu.RUnlock()
	result := make([]ComponentInfo, 0, len(idx.components))
	for _, info := range idx.components {
		result = append(result, info)
	}
	return result
}
