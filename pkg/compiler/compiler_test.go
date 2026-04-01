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

	// Static attrs coalesced with tag open up to dynamic attr boundary
	mustContain(t, src, `"<div id=\"main\" class=\""`)
	mustContain(t, src, `html.EscapeString(fmt.Sprintf("%v", container))`)
	// Closing quote, >, text, and close tag coalesced
	mustContain(t, src, `"\">Hello</div>"`)
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

	// Fully static self-closing element coalesced into one string
	mustContain(t, src, `"<input type=\"text\" />"`)
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
	mustContain(t, src, ".Render(ctx, w)")
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
	mustContain(t, src, "children.Render(ctx, w)")
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
	// <li> coalesced with preceding whitespace before dynamic expr
	mustContain(t, src, `<li>`)
	mustContain(t, src, `html.EscapeString(fmt.Sprintf("%v", item))`)
	mustContain(t, src, `</li>`)
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
	mustContain(t, src, ".Render(ctx, w)")
}

// countWriteStrings counts the number of io.WriteString calls in the source.
func countWriteStrings(src string) int {
	return strings.Count(src, "io.WriteString(w,")
}

func TestCoalescing_PureStatic(t *testing.T) {
	// Pure static template should produce a single io.WriteString call
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name: "Static",
				Body: []parser.Node{
					&parser.HTMLElement{
						Tag: "div",
						Attributes: []parser.Attribute{
							{Name: "class", Value: "container"},
						},
						Children: []parser.Node{
							&parser.HTMLElement{
								Tag: "h1",
								Children: []parser.Node{
									&parser.TextNode{Content: "Hello World"},
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

	count := countWriteStrings(src)
	if count != 1 {
		t.Errorf("expected 1 io.WriteString call for pure static template, got %d\n\n%s", count, src)
	}
	mustContain(t, src, `"<div class=\"container\"><h1>Hello World</h1></div>"`)
}

func TestCoalescing_OneExpression(t *testing.T) {
	// Template with one expression should produce exactly 2 static io.WriteString calls
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Greeting",
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

	// 2 static writes (before and after expr) + 1 dynamic write = 3 total
	count := countWriteStrings(src)
	if count != 3 {
		t.Errorf("expected 3 io.WriteString calls (2 static + 1 dynamic), got %d\n\n%s", count, src)
	}
}

func TestCoalescing_DynamicAttrsFlush(t *testing.T) {
	// Adjacent dynamic attrs should flush between each
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Styled",
				Params: "cls string, id string",
				Body: []parser.Node{
					&parser.HTMLElement{
						Tag: "div",
						Attributes: []parser.Attribute{
							{Name: "class", Value: "cls", Dynamic: true},
							{Name: "id", Value: "id", Dynamic: true},
						},
						Children: []parser.Node{
							&parser.TextNode{Content: "content"},
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

	// Should have: `<div class="` | dynamic | `" id="` | dynamic | `">content</div>`
	mustContain(t, src, `"<div class=\""`)
	mustContain(t, src, `"\" id=\""`)
	mustContain(t, src, `"\">content</div>"`)
}

func TestCoalescing_ControlFlowBreaks(t *testing.T) {
	// Control flow should break coalescing correctly
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Page",
				Params: "show bool",
				Body: []parser.Node{
					&parser.TextNode{Content: "<div>"},
					&parser.IfNode{
						Condition: "show",
						Then: []parser.Node{
							&parser.TextNode{Content: "<span>visible</span>"},
						},
					},
					&parser.TextNode{Content: "</div>"},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	// <div> flushed before if, <span>visible</span> inside if body, </div> after
	mustContain(t, src, `"<div>"`)
	mustContain(t, src, `"<span>visible</span>"`)
	mustContain(t, src, `"</div>"`)
}

func TestCoalescing_ComponentCallFlushes(t *testing.T) {
	// Component calls should flush accumulated static content
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name: "Page",
				Body: []parser.Node{
					&parser.TextNode{Content: "<div>"},
					&parser.ComponentCall{
						Name:      "Header",
						SelfClose: true,
					},
					&parser.TextNode{Content: "</div>"},
				},
			},
		},
	}

	src, err := Compile(file)
	if err != nil {
		t.Fatalf("Compile error: %v", err)
	}

	// <div> flushed before component call, </div> flushed at end
	mustContain(t, src, `"<div>"`)
	mustContain(t, src, "Header().Render(ctx, w)")
	mustContain(t, src, `"</div>"`)
}

func TestCompileURLSanitization(t *testing.T) {
	// Dynamic href should use gox.SanitizeURL
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Link",
				Params: "url string",
				Body: []parser.Node{
					&parser.HTMLElement{
						Tag: "a",
						Attributes: []parser.Attribute{
							{Name: "href", Value: "url", Dynamic: true},
						},
						Children: []parser.Node{
							&parser.TextNode{Content: "Click"},
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

	mustContain(t, src, "gox.SanitizeURL(url)")
	mustNotContain(t, src, "html.EscapeString")
}

func TestCompileURLSanitizationSrcAttr(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Image",
				Params: "imgSrc string",
				Body: []parser.Node{
					&parser.HTMLElement{
						Tag:       "img",
						SelfClose: true,
						Attributes: []parser.Attribute{
							{Name: "src", Value: "imgSrc", Dynamic: true},
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

	mustContain(t, src, "gox.SanitizeURL(imgSrc)")
}

func TestCompileURLSanitizationActionAttr(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Form",
				Params: "target string",
				Body: []parser.Node{
					&parser.HTMLElement{
						Tag: "form",
						Attributes: []parser.Attribute{
							{Name: "action", Value: "target", Dynamic: true},
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

	mustContain(t, src, "gox.SanitizeURL(target)")
}

func TestCompileCSSSanitization(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Styled",
				Params: "css string",
				Body: []parser.Node{
					&parser.HTMLElement{
						Tag: "div",
						Attributes: []parser.Attribute{
							{Name: "style", Value: "css", Dynamic: true},
						},
						Children: []parser.Node{
							&parser.TextNode{Content: "content"},
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

	mustContain(t, src, "gox.SanitizeCSS(css)")
	mustNotContain(t, src, "html.EscapeString")
}

func TestCompileNonURLAttrStillEscapes(t *testing.T) {
	// Non-URL attributes should still use html.EscapeString
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Div",
				Params: "cls string",
				Body: []parser.Node{
					&parser.HTMLElement{
						Tag: "div",
						Attributes: []parser.Attribute{
							{Name: "class", Value: "cls", Dynamic: true},
						},
						Children: []parser.Node{
							&parser.TextNode{Content: "text"},
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

	mustContain(t, src, "html.EscapeString")
	mustNotContain(t, src, "gox.SanitizeURL")
	mustNotContain(t, src, "gox.SanitizeCSS")
}

func TestCompileEndToEndURLSanitization(t *testing.T) {
	input := `package views

func Link(url string) {
  <a href={{ url }}>Click</a>
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

	mustContain(t, src, "gox.SanitizeURL(url)")
}

func TestCompileEndToEndCSSSanitization(t *testing.T) {
	input := `package views

func Box(css string) {
  <div style={{ css }}>content</div>
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

	mustContain(t, src, "gox.SanitizeCSS(css)")
}

func TestCompileSpreadWithURLSanitization(t *testing.T) {
	file := &parser.File{
		Package: &parser.PackageDecl{Name: "views"},
		Components: []*parser.ComponentDecl{
			{
				Name:   "Link",
				Params: "attrs gox.Attrs",
				Body: []parser.Node{
					&parser.HTMLElement{
						Tag: "a",
						Attributes: []parser.Attribute{
							{Spread: true, Value: "attrs"},
						},
						Children: []parser.Node{
							&parser.TextNode{Content: "Click"},
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

	mustContain(t, src, "gox.SanitizeURL(__v)")
	mustContain(t, src, "gox.SanitizeCSS(__v)")
	mustContain(t, src, "strings.ToLower(__k)")
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

func TestCompileConditionalBooleanAttribute(t *testing.T) {
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
	mustNotContain(t, src, `disabled="`)
}

func TestCompileEndToEndConditionalBoolean(t *testing.T) {
	input := `package views

func Button(isDisabled bool) {
  <button disabled?={{ isDisabled }}>Submit</button>
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

	mustContain(t, src, "func Button(isDisabled bool) gox.Component")
	mustContain(t, src, "if isDisabled {")
	mustContain(t, src, `" disabled"`)
}

func TestCompileEndToEndMultipleConditionalBooleans(t *testing.T) {
	input := `package views

func Input(req bool, ro bool) {
  <input type="text" required?={{ req }} readonly?={{ ro }} />
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

	mustContain(t, src, "if req {")
	mustContain(t, src, `" required"`)
	mustContain(t, src, "if ro {")
	mustContain(t, src, `" readonly"`)
}
