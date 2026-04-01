package lsp

import (
	"testing"

	"go.lsp.dev/protocol"
)

func TestComputeDiagnostics_ValidFile(t *testing.T) {
	content := `package views

func Hello(name string) {
	<h1>Hello {{ name }}</h1>
}
`
	diags := computeDiagnostics(content)
	if len(diags) != 0 {
		t.Errorf("expected 0 diagnostics for valid file, got %d: %v", len(diags), diags)
	}
}

func TestComputeDiagnostics_MissingPackage(t *testing.T) {
	content := `func Hello() {
	<h1>Hello</h1>
}
`
	diags := computeDiagnostics(content)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if diags[0].Severity != protocol.DiagnosticSeverityError {
		t.Errorf("expected error severity, got %v", diags[0].Severity)
	}
	if diags[0].Source != "gox" {
		t.Errorf("expected source 'gox', got %q", diags[0].Source)
	}
}

func TestComputeDiagnostics_UnclosedTag(t *testing.T) {
	content := `package views

func Hello() {
	<div>
		<span>unclosed
	</div>
}
`
	diags := computeDiagnostics(content)
	if len(diags) == 0 {
		t.Fatal("expected at least 1 diagnostic for unclosed tag")
	}
	if diags[0].Severity != protocol.DiagnosticSeverityError {
		t.Errorf("expected error severity, got %v", diags[0].Severity)
	}
}

func TestPosRange(t *testing.T) {
	// Parser uses 1-based, LSP uses 0-based
	r := posRange(5, 10)
	if r.Start.Line != 4 {
		t.Errorf("expected line 4, got %d", r.Start.Line)
	}
	if r.Start.Character != 9 {
		t.Errorf("expected character 9, got %d", r.Start.Character)
	}
	if r.End.Character != 10 {
		t.Errorf("expected end character 10, got %d", r.End.Character)
	}
}

func TestPosRange_ZeroInput(t *testing.T) {
	r := posRange(0, 0)
	if r.Start.Line != 0 || r.Start.Character != 0 {
		t.Errorf("expected (0,0), got (%d,%d)", r.Start.Line, r.Start.Character)
	}
}
