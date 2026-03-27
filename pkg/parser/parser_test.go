package parser

import (
	"testing"
)

func TestParsePackageDecl(t *testing.T) {
	input := `package views

func Hello() {
  <h1>Hello</h1>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if file.Package == nil {
		t.Fatal("expected package declaration")
	}
	if file.Package.Name != "views" {
		t.Errorf("expected package name 'views', got %q", file.Package.Name)
	}
}

func TestParseImportDecls(t *testing.T) {
	input := `package views

import "fmt"
import "strings"

func Hello() {
  <p>Hello</p>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(file.Imports) != 2 {
		t.Fatalf("expected 2 imports, got %d", len(file.Imports))
	}
	if file.Imports[0].Path != `"fmt"` {
		t.Errorf("expected import path '\"fmt\"', got %q", file.Imports[0].Path)
	}
	if file.Imports[1].Path != `"strings"` {
		t.Errorf("expected import path '\"strings\"', got %q", file.Imports[1].Path)
	}
}

func TestParseComponentDecl(t *testing.T) {
	input := `package views

func Hello(name string) {
  <h1>Hello</h1>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(file.Components) != 1 {
		t.Fatalf("expected 1 component, got %d", len(file.Components))
	}
	comp := file.Components[0]
	if comp.Name != "Hello" {
		t.Errorf("expected component name 'Hello', got %q", comp.Name)
	}
	if comp.Params != "name string" {
		t.Errorf("expected params 'name string', got %q", comp.Params)
	}
}

func TestParseMultipleComponents(t *testing.T) {
	input := `package views

func Header() {
  <header>Header</header>
}

func Footer() {
  <footer>Footer</footer>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(file.Components) != 2 {
		t.Fatalf("expected 2 components, got %d", len(file.Components))
	}
	if file.Components[0].Name != "Header" {
		t.Errorf("expected 'Header', got %q", file.Components[0].Name)
	}
	if file.Components[1].Name != "Footer" {
		t.Errorf("expected 'Footer', got %q", file.Components[1].Name)
	}
}

func TestParseHTMLElement(t *testing.T) {
	input := `package views

func Hello() {
  <div><p>Hello</p></div>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	comp := file.Components[0]
	// Find the div element
	div := findElement(t, comp.Body, "div")
	if div == nil {
		t.Fatal("expected <div> element")
	}
	pElem := findElement(t, div.Children, "p")
	if pElem == nil {
		t.Fatal("expected <p> element inside <div>")
	}
	text := findText(pElem.Children)
	if text != "Hello" {
		t.Errorf("expected text 'Hello', got %q", text)
	}
}

func TestParseSelfClosingElement(t *testing.T) {
	input := `package views

func Hello() {
  <br />
  <input />
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	comp := file.Components[0]
	var selfClosing int
	for _, node := range comp.Body {
		if elem, ok := node.(*HTMLElement); ok && elem.SelfClose {
			selfClosing++
		}
	}
	if selfClosing != 2 {
		t.Errorf("expected 2 self-closing elements, got %d", selfClosing)
	}
}

func TestParseVoidElements(t *testing.T) {
	input := `package views

func Hello() {
  <img>
  <hr>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	comp := file.Components[0]
	var voids int
	for _, node := range comp.Body {
		if elem, ok := node.(*HTMLElement); ok {
			if voidElements[elem.Tag] {
				voids++
			}
		}
	}
	if voids != 2 {
		t.Errorf("expected 2 void elements, got %d", voids)
	}
}

func TestParseStaticAttributes(t *testing.T) {
	input := `package views

func Hello() {
  <div id="main" class="container"></div>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	div := findElement(t, file.Components[0].Body, "div")
	if div == nil {
		t.Fatal("expected <div> element")
	}
	if len(div.Attributes) != 2 {
		t.Fatalf("expected 2 attributes, got %d", len(div.Attributes))
	}
	assertAttr(t, div.Attributes[0], "id", "main", false, false, false)
	assertAttr(t, div.Attributes[1], "class", "container", false, false, false)
}

func TestParseDynamicAttributes(t *testing.T) {
	input := `package views

func Hello(cls string) {
  <div class={{ cls }}></div>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	div := findElement(t, file.Components[0].Body, "div")
	if div == nil {
		t.Fatal("expected <div>")
	}
	if len(div.Attributes) != 1 {
		t.Fatalf("expected 1 attribute, got %d", len(div.Attributes))
	}
	attr := div.Attributes[0]
	if attr.Name != "class" {
		t.Errorf("expected attr name 'class', got %q", attr.Name)
	}
	if !attr.Dynamic {
		t.Error("expected dynamic attribute")
	}
	if attr.Value != "cls" {
		t.Errorf("expected attr value 'cls', got %q", attr.Value)
	}
}

func TestParseBooleanAttributes(t *testing.T) {
	input := `package views

func Hello() {
  <button disabled>Submit</button>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	btn := findElement(t, file.Components[0].Body, "button")
	if btn == nil {
		t.Fatal("expected <button>")
	}
	if len(btn.Attributes) != 1 {
		t.Fatalf("expected 1 attribute, got %d", len(btn.Attributes))
	}
	if !btn.Attributes[0].Boolean {
		t.Error("expected boolean attribute")
	}
	if btn.Attributes[0].Name != "disabled" {
		t.Errorf("expected attr name 'disabled', got %q", btn.Attributes[0].Name)
	}
}

func TestParseSpreadAttributes(t *testing.T) {
	input := `package views

func Hello() {
  <input {{ attrs... }} />
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var inputElem *HTMLElement
	for _, node := range file.Components[0].Body {
		if elem, ok := node.(*HTMLElement); ok && elem.Tag == "input" {
			inputElem = elem
		}
	}
	if inputElem == nil {
		t.Fatal("expected <input> element")
	}
	if len(inputElem.Attributes) != 1 {
		t.Fatalf("expected 1 attribute, got %d", len(inputElem.Attributes))
	}
	if !inputElem.Attributes[0].Spread {
		t.Error("expected spread attribute")
	}
}

func TestParseExprNode(t *testing.T) {
	input := `package views

func Hello(name string) {
  <p>{{ name }}</p>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	pElem := findElement(t, file.Components[0].Body, "p")
	if pElem == nil {
		t.Fatal("expected <p>")
	}
	var expr *ExprNode
	for _, child := range pElem.Children {
		if e, ok := child.(*ExprNode); ok {
			expr = e
		}
	}
	if expr == nil {
		t.Fatal("expected ExprNode")
	}
	if expr.Expr != "name" {
		t.Errorf("expected expr 'name', got %q", expr.Expr)
	}
}

func TestParseRawExprNode(t *testing.T) {
	input := `package views

func Hello(rawHTML string) {
  <div>{{! rawHTML }}</div>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	div := findElement(t, file.Components[0].Body, "div")
	if div == nil {
		t.Fatal("expected <div>")
	}
	var rawExpr *RawExprNode
	for _, child := range div.Children {
		if r, ok := child.(*RawExprNode); ok {
			rawExpr = r
		}
	}
	if rawExpr == nil {
		t.Fatal("expected RawExprNode")
	}
	if rawExpr.Expr != "rawHTML" {
		t.Errorf("expected expr 'rawHTML', got %q", rawExpr.Expr)
	}
}

func TestParseIfElse(t *testing.T) {
	input := `package views

func Hello(isAdmin bool) {
  {{ if isAdmin }}
    <span>Admin</span>
  {{ else }}
    <span>User</span>
  {{ end }}
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	comp := file.Components[0]
	var ifNode *IfNode
	for _, node := range comp.Body {
		if n, ok := node.(*IfNode); ok {
			ifNode = n
		}
	}
	if ifNode == nil {
		t.Fatal("expected IfNode")
	}
	if ifNode.Condition != "isAdmin" {
		t.Errorf("expected condition 'isAdmin', got %q", ifNode.Condition)
	}
	if len(ifNode.Then) == 0 {
		t.Error("expected non-empty then branch")
	}
	if len(ifNode.Else) == 0 {
		t.Error("expected non-empty else branch")
	}
	// Check then branch has <span>Admin</span>
	thenSpan := findElement(t, ifNode.Then, "span")
	if thenSpan == nil {
		t.Error("expected <span> in then branch")
	}
	// Check else branch has <span>User</span>
	elseSpan := findElement(t, ifNode.Else, "span")
	if elseSpan == nil {
		t.Error("expected <span> in else branch")
	}
}

func TestParseIfWithoutElse(t *testing.T) {
	input := `package views

func Hello(show bool) {
  {{ if show }}
    <p>Visible</p>
  {{ end }}
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var ifNode *IfNode
	for _, node := range file.Components[0].Body {
		if n, ok := node.(*IfNode); ok {
			ifNode = n
		}
	}
	if ifNode == nil {
		t.Fatal("expected IfNode")
	}
	if len(ifNode.Then) == 0 {
		t.Error("expected non-empty then branch")
	}
	if len(ifNode.Else) != 0 {
		t.Error("expected empty else branch")
	}
}

func TestParseForLoop(t *testing.T) {
	input := `package views

func List(items []string) {
  <ul>
    {{ for _, item := range items }}
      <li>{{ item }}</li>
    {{ end }}
  </ul>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ul := findElement(t, file.Components[0].Body, "ul")
	if ul == nil {
		t.Fatal("expected <ul>")
	}
	var forNode *ForNode
	for _, child := range ul.Children {
		if n, ok := child.(*ForNode); ok {
			forNode = n
		}
	}
	if forNode == nil {
		t.Fatal("expected ForNode")
	}
	if forNode.Clause != "_, item := range items" {
		t.Errorf("expected clause '_, item := range items', got %q", forNode.Clause)
	}
	// Check body has <li>
	li := findElement(t, forNode.Body, "li")
	if li == nil {
		t.Error("expected <li> inside for body")
	}
}

func TestParseSwitch(t *testing.T) {
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
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var switchNode *SwitchNode
	for _, node := range file.Components[0].Body {
		if n, ok := node.(*SwitchNode); ok {
			switchNode = n
		}
	}
	if switchNode == nil {
		t.Fatal("expected SwitchNode")
	}
	if switchNode.Expr != "status" {
		t.Errorf("expected expr 'status', got %q", switchNode.Expr)
	}
	if len(switchNode.Cases) != 3 {
		t.Fatalf("expected 3 cases, got %d", len(switchNode.Cases))
	}
	if switchNode.Cases[0].Value != `"active"` {
		t.Errorf("expected case value '\"active\"', got %q", switchNode.Cases[0].Value)
	}
	if switchNode.Cases[1].Value != `"inactive"` {
		t.Errorf("expected case value '\"inactive\"', got %q", switchNode.Cases[1].Value)
	}
	if !switchNode.Cases[2].Default {
		t.Error("expected default case")
	}
}

func TestParseComponentCall(t *testing.T) {
	input := `package views

func Page() {
  <Header />
  <main>Content</main>
  <Footer />
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	comp := file.Components[0]
	var components []*ComponentCall
	for _, node := range comp.Body {
		if c, ok := node.(*ComponentCall); ok {
			components = append(components, c)
		}
	}
	if len(components) != 2 {
		t.Fatalf("expected 2 component calls, got %d", len(components))
	}
	if components[0].Name != "Header" {
		t.Errorf("expected 'Header', got %q", components[0].Name)
	}
	if !components[0].SelfClose {
		t.Error("expected Header to be self-closing")
	}
	if components[1].Name != "Footer" {
		t.Errorf("expected 'Footer', got %q", components[1].Name)
	}
}

func TestParseComponentCallWithChildren(t *testing.T) {
	input := `package views

func Page() {
  <Card title="Welcome">
    <p>Hello, world!</p>
  </Card>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var card *ComponentCall
	for _, node := range file.Components[0].Body {
		if c, ok := node.(*ComponentCall); ok && c.Name == "Card" {
			card = c
		}
	}
	if card == nil {
		t.Fatal("expected Card component call")
	}
	if len(card.Attributes) != 1 {
		t.Fatalf("expected 1 attribute, got %d", len(card.Attributes))
	}
	if card.Attributes[0].Name != "title" {
		t.Errorf("expected attr name 'title', got %q", card.Attributes[0].Name)
	}
	if card.Attributes[0].Value != "Welcome" {
		t.Errorf("expected attr value 'Welcome', got %q", card.Attributes[0].Value)
	}
	// Check children contain a <p>
	pElem := findElement(t, card.Children, "p")
	if pElem == nil {
		t.Error("expected <p> inside Card children")
	}
}

func TestParseChildrenSlot(t *testing.T) {
	input := `package views

func Card(title string) {
  <div class="card">
    <h2>{{ title }}</h2>
    {{ children }}
  </div>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	div := findElement(t, file.Components[0].Body, "div")
	if div == nil {
		t.Fatal("expected <div>")
	}
	var childrenNode *ChildrenNode
	for _, child := range div.Children {
		if c, ok := child.(*ChildrenNode); ok {
			childrenNode = c
		}
	}
	if childrenNode == nil {
		t.Fatal("expected ChildrenNode")
	}
}

func TestParseNestedStructure(t *testing.T) {
	input := `package views

import "fmt"

func Page(title string, items []string) {
  <html>
    <head><title>{{ title }}</title></head>
    <body>
      <Header />
      <main>
        {{ if len(items) > 0 }}
          <ul>
            {{ for _, item := range items }}
              <li>{{ item }}</li>
            {{ end }}
          </ul>
        {{ else }}
          <p>No items</p>
        {{ end }}
      </main>
      <Footer />
    </body>
  </html>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if file.Package.Name != "views" {
		t.Errorf("expected package 'views', got %q", file.Package.Name)
	}
	if len(file.Imports) != 1 {
		t.Errorf("expected 1 import, got %d", len(file.Imports))
	}
	if len(file.Components) != 1 {
		t.Fatalf("expected 1 component, got %d", len(file.Components))
	}
	comp := file.Components[0]
	if comp.Name != "Page" {
		t.Errorf("expected 'Page', got %q", comp.Name)
	}

	// Verify nested structure: html > body > main contains an IfNode
	html := findElement(t, comp.Body, "html")
	if html == nil {
		t.Fatal("expected <html>")
	}
	body := findElement(t, html.Children, "body")
	if body == nil {
		t.Fatal("expected <body>")
	}
	main := findElement(t, body.Children, "main")
	if main == nil {
		t.Fatal("expected <main>")
	}
	var ifNode *IfNode
	for _, child := range main.Children {
		if n, ok := child.(*IfNode); ok {
			ifNode = n
		}
	}
	if ifNode == nil {
		t.Fatal("expected IfNode inside <main>")
	}
	if ifNode.Condition != "len(items) > 0" {
		t.Errorf("expected condition 'len(items) > 0', got %q", ifNode.Condition)
	}
}

func TestParseComponentWithNoParams(t *testing.T) {
	input := `package views

func Hello() {
  <p>Hello</p>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if file.Components[0].Params != "" {
		t.Errorf("expected empty params, got %q", file.Components[0].Params)
	}
}

func TestParseErrorMismatchedTags(t *testing.T) {
	input := `package views

func Hello() {
  <div></span>
}`
	p := NewParser(input)
	_, err := p.Parse()
	if err == nil {
		t.Fatal("expected error for mismatched tags")
	}
}

func TestParseErrorMissingPackage(t *testing.T) {
	input := `func Hello() {
  <p>Hello</p>
}`
	p := NewParser(input)
	_, err := p.Parse()
	if err == nil {
		t.Fatal("expected error for missing package declaration")
	}
}

func TestParseTextNode(t *testing.T) {
	input := `package views

func Hello() {
  <p>Hello World</p>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	pElem := findElement(t, file.Components[0].Body, "p")
	if pElem == nil {
		t.Fatal("expected <p>")
	}
	text := findText(pElem.Children)
	if text != "Hello World" {
		t.Errorf("expected 'Hello World', got %q", text)
	}
}

func TestParseMixedTextAndExpressions(t *testing.T) {
	input := `package views

func Hello(name string) {
  <h1>Hello, {{ name }}!</h1>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	h1 := findElement(t, file.Components[0].Body, "h1")
	if h1 == nil {
		t.Fatal("expected <h1>")
	}
	// Should have: TextNode("Hello, "), ExprNode("name"), TextNode("!")
	if len(h1.Children) < 3 {
		t.Fatalf("expected at least 3 children in h1, got %d", len(h1.Children))
	}

	text0, ok := h1.Children[0].(*TextNode)
	if !ok {
		t.Fatalf("expected TextNode first, got %T", h1.Children[0])
	}
	if text0.Content != "Hello, " {
		t.Errorf("expected 'Hello, ', got %q", text0.Content)
	}

	expr, ok := h1.Children[1].(*ExprNode)
	if !ok {
		t.Fatalf("expected ExprNode second, got %T", h1.Children[1])
	}
	if expr.Expr != "name" {
		t.Errorf("expected 'name', got %q", expr.Expr)
	}
}

func TestParseLineColumnTracking(t *testing.T) {
	input := `package views

func Hello() {
  <p>Hello</p>
}`
	p := NewParser(input)
	file, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if file.Package.Line != 1 {
		t.Errorf("expected package at line 1, got %d", file.Package.Line)
	}
	comp := file.Components[0]
	if comp.Line != 3 {
		t.Errorf("expected component at line 3, got %d", comp.Line)
	}
}

// -- helpers --

func findElement(t *testing.T, nodes []Node, tag string) *HTMLElement {
	t.Helper()
	for _, node := range nodes {
		if elem, ok := node.(*HTMLElement); ok && elem.Tag == tag {
			return elem
		}
	}
	return nil
}

func findText(nodes []Node) string {
	for _, node := range nodes {
		if text, ok := node.(*TextNode); ok {
			return text.Content
		}
	}
	return ""
}

func assertAttr(t *testing.T, attr Attribute, name, value string, dynamic, boolean, spread bool) {
	t.Helper()
	if attr.Name != name {
		t.Errorf("expected attr name %q, got %q", name, attr.Name)
	}
	if attr.Value != value {
		t.Errorf("expected attr value %q, got %q", value, attr.Value)
	}
	if attr.Dynamic != dynamic {
		t.Errorf("expected attr dynamic=%v, got %v", dynamic, attr.Dynamic)
	}
	if attr.Boolean != boolean {
		t.Errorf("expected attr boolean=%v, got %v", boolean, attr.Boolean)
	}
	if attr.Spread != spread {
		t.Errorf("expected attr spread=%v, got %v", spread, attr.Spread)
	}
}
