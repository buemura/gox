package lsp

import (
	"testing"

	"go.lsp.dev/protocol"
)

func TestComponentNameAtPosition_OnComponentCall(t *testing.T) {
	content := `package views

func Page() {
	<div>
		<Header />
	</div>
}
`
	// <Header /> is on line 5 (0-based: 4), the 'H' starts after the tab and '<'
	// In the source: "\t\t<Header />" — '<' is at col 2 (0-based), 'H' at col 3
	pos := protocol.Position{Line: 4, Character: 3}
	name := componentNameAtPosition(content, pos)
	if name != "Header" {
		t.Errorf("expected 'Header', got %q", name)
	}
}

func TestComponentNameAtPosition_OnHTMLElement(t *testing.T) {
	content := `package views

func Page() {
	<div>hello</div>
}
`
	// 'div' is lowercase, should not match as component.
	pos := protocol.Position{Line: 3, Character: 2}
	name := componentNameAtPosition(content, pos)
	if name != "" {
		t.Errorf("expected empty for HTML element, got %q", name)
	}
}

func TestComponentNameAtPosition_OutsideTag(t *testing.T) {
	content := `package views

func Page() {
	<div>hello</div>
}
`
	// Position on 'hello' text — not a tag.
	pos := protocol.Position{Line: 3, Character: 6}
	name := componentNameAtPosition(content, pos)
	if name != "" {
		t.Errorf("expected empty for text content, got %q", name)
	}
}

func TestComponentNameFromLine(t *testing.T) {
	content := `package views

func Page() {
	<div>
		<Header title="hello" />
	</div>
}
`
	// Test the fallback line-based extraction.
	pos := protocol.Position{Line: 4, Character: 3}
	name := componentNameFromLine(content, pos)
	if name != "Header" {
		t.Errorf("expected 'Header', got %q", name)
	}
}

func TestComponentNameFromLine_NotComponent(t *testing.T) {
	content := "  <div>hello</div>"
	pos := protocol.Position{Line: 0, Character: 3}
	name := componentNameFromLine(content, pos)
	if name != "" {
		t.Errorf("expected empty for lowercase tag, got %q", name)
	}
}
