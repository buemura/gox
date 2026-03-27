package compiler

import (
	"strings"
	"testing"

	"github.com/buemura/gox/pkg/parser"
)

func TestCompileSimpleComponent(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Hello",
				Params: "name string",
				Body: []parser.Node{
					&parser.TextNode{Content: "<h1>Hello, "},
					&parser.ExprNode{Expr: "name"},
					&parser.TextNode{Content: "!</h1>"},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "package views")
	mustContain(t, src, "func Hello(name string) gox.Component")
	mustContain(t, src, `io.WriteString(w, "<h1>Hello, ")`)
	mustContain(t, src, `html.EscapeString(fmt.Sprintf("%v", name))`)
	mustContain(t, src, `io.WriteString(w, "!</h1>")`)
}

func TestCompileRawExpression(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Raw",
				Params: "content string",
				Body: []parser.Node{
					&parser.TextNode{Content: "<div>"},
					&parser.RawExprNode{Expr: "content"},
					&parser.TextNode{Content: "</div>"},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, `fmt.Sprintf("%v", content)`)
	mustNotContain(t, src, "html.EscapeString")
}

func TestCompileIfElse(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Greeting",
				Params: "isAdmin bool",
				Body: []parser.Node{
					&parser.IfNode{
						Condition: "isAdmin",
						Then: []parser.Node{
							&parser.TextNode{Content: "<span>Admin</span>"},
						},
						Else: []parser.Node{
							&parser.TextNode{Content: "<span>User</span>"},
						},
					},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "if isAdmin {")
	mustContain(t, src, "} else {")
	mustContain(t, src, `"<span>Admin</span>"`)
	mustContain(t, src, `"<span>User</span>"`)
}

func TestCompileForLoop(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "List",
				Params: "items []string",
				Body: []parser.Node{
					&parser.TextNode{Content: "<ul>"},
					&parser.ForNode{
						Clause: "_, item := range items",
						Body: []parser.Node{
							&parser.TextNode{Content: "<li>"},
							&parser.ExprNode{Expr: "item"},
							&parser.TextNode{Content: "</li>"},
						},
					},
					&parser.TextNode{Content: "</ul>"},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "for _, item := range items {")
	mustContain(t, src, `"<li>"`)
	mustContain(t, src, `html.EscapeString(fmt.Sprintf("%v", item))`)
}

func TestCompileSwitchCase(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Status",
				Params: "status string",
				Body: []parser.Node{
					&parser.SwitchNode{
						Expr: "status",
						Cases: []*parser.CaseClause{
							{
								Value: `"active"`,
								Body: []parser.Node{
									&parser.TextNode{Content: "<span>Active</span>"},
								},
							},
							{
								Default: true,
								Body: []parser.Node{
									&parser.TextNode{Content: "<span>Unknown</span>"},
								},
							},
						},
					},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "switch status {")
	mustContain(t, src, `case "active":`)
	mustContain(t, src, "default:")
}

func TestCompileHTMLElement(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name: "Page",
				Body: []parser.Node{
					&parser.HTMLElement{
						Tag: "div",
						Attributes: []parser.Attribute{
							{Name: "id", Value: "main"},
							{Name: "class", Value: "container", Dynamic: true},
						},
						Children: []parser.Node{
							&parser.TextNode{Content: "Hello"},
						},
					},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, `"<div"`)
	mustContain(t, src, `" id=\"main\""`)
	mustContain(t, src, `" class=\""`)
	mustContain(t, src, `html.EscapeString(fmt.Sprintf("%v", container))`)
	mustContain(t, src, `"</div>"`)
}

func TestCompileSelfClosingElement(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name: "Form",
				Body: []parser.Node{
					&parser.HTMLElement{
						Tag:       "input",
						SelfClose: true,
						Attributes: []parser.Attribute{
							{Name: "type", Value: "text"},
						},
					},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, `"<input"`)
	mustContain(t, src, `" />"`)
	mustNotContain(t, src, `"</input>"`)
}

func TestCompileBooleanAttribute(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Button",
				Params: "isDisabled bool",
				Body: []parser.Node{
					&parser.HTMLElement{
						Tag: "button",
						Attributes: []parser.Attribute{
							{Name: "disabled", Boolean: true, Dynamic: true, Value: "isDisabled"},
						},
						Children: []parser.Node{
							&parser.TextNode{Content: "Submit"},
						},
					},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "if isDisabled {")
	mustContain(t, src, `" disabled"`)
}

func TestCompileComponentCall(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name: "Page",
				Body: []parser.Node{
					&parser.ComponentCall{
						Name:      "Header",
						SelfClose: true,
					},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "Header()")
	mustContain(t, src, ".Render(w)")
}

func TestCompileComponentCallWithChildren(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name: "Page",
				Body: []parser.Node{
					&parser.ComponentCall{
						Name: "Card",
						Attributes: []parser.Attribute{
							{Name: "title", Value: "Welcome", Dynamic: true},
						},
						Children: []parser.Node{
							&parser.TextNode{Content: "<p>Hello!</p>"},
						},
					},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "Card(Welcome,")
	mustContain(t, src, "gox.ComponentFunc")
	mustContain(t, src, `"<p>Hello!</p>"`)
}

func TestCompileChildrenSlot(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Card",
				Params: "title string",
				Body: []parser.Node{
					&parser.TextNode{Content: "<div>"},
					&parser.ChildrenNode{},
					&parser.TextNode{Content: "</div>"},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "children gox.Component")
	mustContain(t, src, "if children != nil {")
	mustContain(t, src, "children.Render(w)")
}

func TestCompileImports(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Imports: []*parser.ImportDecl{
			{Path: "strings"},
		},
		Components: []*parser.ComponentDecl{
			{
				Name: "Hello",
				Body: []parser.Node{
					&parser.TextNode{Content: "<div>hello</div>"},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, `"strings"`)
	mustContain(t, src, `"io"`)
	mustContain(t, src, `"html"`)
	mustContain(t, src, `"fmt"`)
}

func TestCompileSpreadAttributes(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Input",
				Params: "attrs gox.Attrs",
				Body: []parser.Node{
					&parser.HTMLElement{
						Tag:       "input",
						SelfClose: true,
						Attributes: []parser.Attribute{
							{Spread: true, Value: "attrs"},
						},
					},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "sort.Strings(__keys)")
	mustContain(t, src, "for _, __k := range __keys")
}

func TestCompileEndToEnd(t *testing.T) {
	// Parse a real .gox-like input through AST and compile
	input := `package views

import "fmt"

func Hello(name string) {
  <h1>Hello, {{ name }}!</h1>
}
`
	p := parser.NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "package views")
	mustContain(t, src, "func Hello(name string) gox.Component")
	mustContain(t, src, `html.EscapeString`)
}

func TestCompileEndToEndIfElse(t *testing.T) {
	input := `package views

func Greeting(isAdmin bool) {
  {{ if isAdmin }}
    <span>Admin</span>
  {{ else }}
    <span>User</span>
  {{ end }}
}
`
	p := parser.NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "if isAdmin {")
	mustContain(t, src, "} else {")
}

func TestCompileEndToEndForLoop(t *testing.T) {
	input := `package views

func List(items []string) {
  <ul>
    {{ for _, item := range items }}
      <li>{{ item }}</li>
    {{ end }}
  </ul>
}
`
	p := parser.NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "for _, item := range items {")
	mustContain(t, src, `io.WriteString(w, "<li"`)
	mustContain(t, src, `"</li>"`)
}

func TestCompileEndToEndSwitch(t *testing.T) {
	input := `package views

func Status(status string) {
  {{ switch status }}
    {{ case "active" }}
      <span>Active</span>
    {{ case "inactive" }}
      <span>Inactive</span>
    {{ default }}
      <span>Unknown</span>
  {{ end }}
}
`
	p := parser.NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "switch status {")
	mustContain(t, src, `case "active":`)
	mustContain(t, src, "default:")
}

func TestCompileEndToEndComponentCall(t *testing.T) {
	input := `package views

func Page() {
  <div>
    <Header />
  </div>
}
`
	p := parser.NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	mustContain(t, src, "Header()")
	mustContain(t, src, ".Render(w)")
}

func mustContain(t *testing.T, src, substr string) {
	t.Helper()
	if !strings.Contains(src, substr) {
		t.Errorf("expected output to contain %q\n\ngot:\n%s", substr, src)
	}
}

func mustNotContain(t *testing.T, src, substr string) {
	t.Helper()
	if strings.Contains(src, substr) {
		t.Errorf("expected output NOT to contain %q\n\ngot:\n%s", substr, src)
	}
}
