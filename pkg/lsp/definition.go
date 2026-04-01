package lsp

import (
	"strings"
	"unicode"

	"github.com/buemura/gox/pkg/parser"
	"go.lsp.dev/protocol"
)

// findDefinition returns the location of a component's declaration given a
// cursor position in a document. Returns nil if the cursor is not on a
// component call.
func (s *Server) findDefinition(uri, content string, pos protocol.Position) []protocol.Location {
	name := componentNameAtPosition(content, pos)
	if name == "" {
		return nil
	}

	info, ok := s.index.Lookup(name)
	if !ok {
		return nil
	}

	return []protocol.Location{
		{
			URI:   protocol.DocumentURI(info.URI),
			Range: posRange(info.Line, info.Col),
		},
	}
}

// componentNameAtPosition extracts the component name (uppercase tag) at the
// given cursor position by re-lexing the document and checking if the cursor
// falls within a ComponentCall node's tag name.
func componentNameAtPosition(content string, pos protocol.Position) string {
	// Strategy: parse the file and walk the AST looking for ComponentCall nodes
	// whose position range covers the cursor.
	p := parser.NewParser(content)
	file, err := p.Parse()
	if err != nil || file == nil {
		// Fallback: try to extract from the raw line.
		return componentNameFromLine(content, pos)
	}

	targetLine := int(pos.Line) + 1   // 0-based -> 1-based
	targetCol := int(pos.Character) + 1

	for _, comp := range file.Components {
		if name := findComponentCallInNodes(comp.Body, targetLine, targetCol); name != "" {
			return name
		}
	}
	return ""
}

// findComponentCallInNodes recursively searches for a ComponentCall at the target position.
func findComponentCallInNodes(nodes []parser.Node, line, col int) string {
	for _, node := range nodes {
		switch n := node.(type) {
		case *parser.ComponentCall:
			nodeLine, nodeCol := n.Pos()
			// The tag name starts after '<', so the name occupies columns nodeCol+1 .. nodeCol+len(name)
			nameStart := nodeCol + 1 // skip '<'
			nameEnd := nameStart + len(n.Name)
			if line == nodeLine && col >= nameStart && col <= nameEnd {
				return n.Name
			}
			// Also check children.
			if name := findComponentCallInNodes(n.Children, line, col); name != "" {
				return name
			}
		case *parser.HTMLElement:
			if name := findComponentCallInNodes(n.Children, line, col); name != "" {
				return name
			}
		case *parser.IfNode:
			if name := findComponentCallInNodes(n.Then, line, col); name != "" {
				return name
			}
			if name := findComponentCallInNodes(n.Else, line, col); name != "" {
				return name
			}
		case *parser.ForNode:
			if name := findComponentCallInNodes(n.Body, line, col); name != "" {
				return name
			}
		case *parser.SwitchNode:
			for _, c := range n.Cases {
				if name := findComponentCallInNodes(c.Body, line, col); name != "" {
					return name
				}
			}
		}
	}
	return ""
}

// componentNameFromLine is a fallback that extracts a component name from
// the raw text at the cursor's line when parsing fails.
func componentNameFromLine(content string, pos protocol.Position) string {
	lines := strings.Split(content, "\n")
	lineIdx := int(pos.Line)
	if lineIdx >= len(lines) {
		return ""
	}
	line := lines[lineIdx]
	col := int(pos.Character)

	// Walk backwards to find '<', then forwards to extract the name.
	start := col
	for start > 0 && line[start-1] != '<' {
		if !isIdentChar(rune(line[start-1])) {
			break
		}
		start--
	}
	if start > 0 && line[start-1] == '<' {
		// start is now at the first char of the tag name.
	} else {
		return ""
	}

	end := start
	for end < len(line) && isIdentChar(rune(line[end])) {
		end++
	}

	name := line[start:end]
	if len(name) == 0 {
		return ""
	}
	// Must start with uppercase (component, not HTML element).
	if !unicode.IsUpper(rune(name[0])) {
		return ""
	}
	return name
}

func isIdentChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}
