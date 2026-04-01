package lsp

import (
	"fmt"

	"go.lsp.dev/protocol"
)

// computeHover returns hover information for the given document position.
func (s *Server) computeHover(uri, content string, pos protocol.Position) *protocol.Hover {
	name := componentNameAtPosition(content, pos)
	if name == "" {
		return nil
	}

	info, ok := s.index.Lookup(name)
	if !ok {
		return nil
	}

	sig := fmt.Sprintf("func %s(%s)", info.Name, info.Params)
	markdown := fmt.Sprintf("```go\n%s\n```", sig)

	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.Markdown,
			Value: markdown,
		},
	}
}
