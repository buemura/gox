package lsp

import (
	"context"
	"errors"

	"github.com/buemura/gox/pkg/parser"
	"go.lsp.dev/protocol"
)

// diagnose parses the document content and publishes diagnostics to the client.
func (s *Server) diagnose(ctx context.Context, uri string, content string) {
	diags := computeDiagnostics(content)
	s.client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
		URI:         protocol.DocumentURI(uri),
		Diagnostics: diags,
	})
}

// computeDiagnostics parses the content and returns LSP diagnostics.
func computeDiagnostics(content string) []protocol.Diagnostic {
	p := parser.NewParser(content)
	_, err := p.Parse()
	if err == nil {
		return []protocol.Diagnostic{}
	}

	var pe *parser.ParseError
	if errors.As(err, &pe) {
		return []protocol.Diagnostic{
			{
				Range:    posRange(pe.Line, pe.Col),
				Severity: protocol.DiagnosticSeverityError,
				Source:   "gox",
				Message:  pe.Message,
			},
		}
	}

	// Unknown error type — report at beginning of file.
	return []protocol.Diagnostic{
		{
			Range:    posRange(1, 1),
			Severity: protocol.DiagnosticSeverityError,
			Source:   "gox",
			Message:  err.Error(),
		},
	}
}

// posRange converts a 1-based parser position to a 0-based LSP range.
// The range covers a single character at the given position.
func posRange(line, col int) protocol.Range {
	l := uint32(0)
	c := uint32(0)
	if line > 0 {
		l = uint32(line - 1)
	}
	if col > 0 {
		c = uint32(col - 1)
	}
	return protocol.Range{
		Start: protocol.Position{Line: l, Character: c},
		End:   protocol.Position{Line: l, Character: c + 1},
	}
}
