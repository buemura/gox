package lsp

import "sync"

// DocumentStore is a thread-safe in-memory store for open document contents.
type DocumentStore struct {
	mu   sync.RWMutex
	docs map[string]string // URI -> content
}

// NewDocumentStore creates an empty document store.
func NewDocumentStore() *DocumentStore {
	return &DocumentStore{docs: make(map[string]string)}
}

// Open stores a newly opened document.
func (s *DocumentStore) Open(uri, content string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.docs[uri] = content
}

// Update replaces the content of an already-open document.
func (s *DocumentStore) Update(uri, content string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.docs[uri] = content
}

// Close removes a document from the store.
func (s *DocumentStore) Close(uri string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.docs, uri)
}

// Get returns the content of an open document and whether it exists.
func (s *DocumentStore) Get(uri string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	content, ok := s.docs[uri]
	return content, ok
}
