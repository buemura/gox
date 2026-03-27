package formatter

import (
	"testing"

	"github.com/buemura/gox/pkg/parser"
)

func TestFormatPackageAndImports(t *testing.T) {
	src := `package views
import "fmt"
import "strings"

func Hello() {
  <h1>Hello</h1>
}
`
	p := parser.NewParser(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	got := Format(ast)
	expected := `package views

import "fmt"
import "strings"

func Hello() {
  <h1>Hello</h1>
}
`
	if got != expected {
		t.Errorf("mismatch:\ngot:\n%s\nexpected:\n%s", got, expected)
	}
}

func TestFormatComponentWithParams(t *testing.T) {
	src := `package views

func Greet(name string, age int) {
  <p>Hello {{ name }}, age {{ age }}</p>
}
`
	p := parser.NewParser(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	got := Format(ast)
	expected := `package views

func Greet(name string, age int) {
  <p>Hello {{ name }}, age {{ age }}</p>
}
`
	if got != expected {
		t.Errorf("mismatch:\ngot:\n%s\nexpected:\n%s", got, expected)
	}
}

func TestFormatNestedElements(t *testing.T) {
	src := `package views

func Page() {
  <div>
    <h1>Title</h1>
    <p>Content</p>
  </div>
}
`
	p := parser.NewParser(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	got := Format(ast)
	expected := `package views

func Page() {
  <div>
    <h1>Title</h1>
    <p>Content</p>
  </div>
}
`
	if got != expected {
		t.Errorf("mismatch:\ngot:\n%s\nexpected:\n%s", got, expected)
	}
}

func TestFormatSelfClosing(t *testing.T) {
	src := `package views

func Form() {
  <input type="text" />
  <br />
}
`
	p := parser.NewParser(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	got := Format(ast)
	expected := `package views

func Form() {
  <input type="text" />
  <br />
}
`
	if got != expected {
		t.Errorf("mismatch:\ngot:\n%s\nexpected:\n%s", got, expected)
	}
}

func TestFormatIfElse(t *testing.T) {
	src := `package views

func Status(ok bool) {
  {{ if ok }}
    <p>OK</p>
  {{ else }}
    <p>Error</p>
  {{ end }}
}
`
	p := parser.NewParser(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	got := Format(ast)
	expected := `package views

func Status(ok bool) {
  {{ if ok }}
    <p>OK</p>
  {{ else }}
    <p>Error</p>
  {{ end }}
}
`
	if got != expected {
		t.Errorf("mismatch:\ngot:\n%s\nexpected:\n%s", got, expected)
	}
}

func TestFormatForLoop(t *testing.T) {
	src := `package views

func List(items []string) {
  <ul>
    {{ for _, item := range items }}
      <li>{{ item }}</li>
    {{ end }}
  </ul>
}
`
	p := parser.NewParser(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	got := Format(ast)
	expected := `package views

func List(items []string) {
  <ul>
    {{ for _, item := range items }}
      <li>{{ item }}</li>
    {{ end }}
  </ul>
}
`
	if got != expected {
		t.Errorf("mismatch:\ngot:\n%s\nexpected:\n%s", got, expected)
	}
}

func TestFormatSwitch(t *testing.T) {
	src := `package views

func Badge(level string) {
  {{ switch level }}
    {{ case "admin" }}
      <span>Admin</span>
    {{ case "user" }}
      <span>User</span>
    {{ default }}
      <span>Guest</span>
  {{ end }}
}
`
	p := parser.NewParser(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	got := Format(ast)
	expected := `package views

func Badge(level string) {
  {{ switch level }}
    {{ case "admin" }}
      <span>Admin</span>
    {{ case "user" }}
      <span>User</span>
    {{ default }}
      <span>Guest</span>
  {{ end }}
}
`
	if got != expected {
		t.Errorf("mismatch:\ngot:\n%s\nexpected:\n%s", got, expected)
	}
}

func TestFormatAttributes(t *testing.T) {
	src := `package views

func Button(cls string, disabled bool) {
  <button class={{ cls }} id="main" disabled>Click</button>
}
`
	p := parser.NewParser(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	got := Format(ast)
	expected := `package views

func Button(cls string, disabled bool) {
  <button class={{ cls }} id="main" disabled>Click</button>
}
`
	if got != expected {
		t.Errorf("mismatch:\ngot:\n%s\nexpected:\n%s", got, expected)
	}
}

func TestFormatRawExpr(t *testing.T) {
	src := `package views

func Content(html string) {
  <div>
    {{! html }}
  </div>
}
`
	p := parser.NewParser(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	got := Format(ast)
	// RawExprNode is inline, so div with only a raw expr stays on one line.
	expected := `package views

func Content(html string) {
  <div>{{! html }}</div>
}
`
	if got != expected {
		t.Errorf("mismatch:\ngot:\n%s\nexpected:\n%s", got, expected)
	}
}

func TestFormatChildren(t *testing.T) {
	src := `package views

func Layout() {
  <main>
    {{ children }}
  </main>
}
`
	p := parser.NewParser(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	got := Format(ast)
	expected := `package views

func Layout() {
  <main>
    {{ children }}
  </main>
}
`
	if got != expected {
		t.Errorf("mismatch:\ngot:\n%s\nexpected:\n%s", got, expected)
	}
}

func TestFormatComponentCall(t *testing.T) {
	src := `package views

func Page() {
  <Layout>
    <h1>Hello</h1>
  </Layout>
}
`
	p := parser.NewParser(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	got := Format(ast)
	expected := `package views

func Page() {
  <Layout>
    <h1>Hello</h1>
  </Layout>
}
`
	if got != expected {
		t.Errorf("mismatch:\ngot:\n%s\nexpected:\n%s", got, expected)
	}
}

func TestFormatMultipleComponents(t *testing.T) {
	src := `package views

func Hello() {
  <p>Hello</p>
}

func World() {
  <p>World</p>
}
`
	p := parser.NewParser(src)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	got := Format(ast)
	expected := `package views

func Hello() {
  <p>Hello</p>
}

func World() {
  <p>World</p>
}
`
	if got != expected {
		t.Errorf("mismatch:\ngot:\n%s\nexpected:\n%s", got, expected)
	}
}
