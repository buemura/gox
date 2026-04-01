package parser

import "testing"

// tokenize collects all tokens from the lexer.
func tokenize(input string) []Token {
	l := NewLexer(input)
	var tokens []Token
	for {
		tok := l.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == TokenEOF {
			break
		}
	}
	return tokens
}

func assertTokens(t *testing.T, input string, expected []Token) {
	t.Helper()
	tokens := tokenize(input)
	if len(tokens) != len(expected) {
		t.Fatalf("token count: got %d, want %d\ngot:  %v", len(tokens), len(expected), tokens)
	}
	for i, want := range expected {
		got := tokens[i]
		if got.Type != want.Type {
			t.Errorf("token[%d] type: got %s, want %s\n  got:  %s", i, got.Type, want.Type, got)
		}
		if want.Value != "" && got.Value != want.Value {
			t.Errorf("token[%d] value: got %q, want %q", i, got.Value, want.Value)
		}
	}
}

func TestPackageDecl(t *testing.T) {
	assertTokens(t, "package views\n", []Token{
		{Type: TokenPackageDecl, Value: "package views"},
		{Type: TokenEOF},
	})
}

func TestImportDecl(t *testing.T) {
	assertTokens(t, "import \"fmt\"\n", []Token{
		{Type: TokenImportDecl, Value: `import "fmt"`},
		{Type: TokenEOF},
	})
}

func TestMultipleImports(t *testing.T) {
	input := "package views\n\nimport \"fmt\"\nimport \"strings\"\n"
	assertTokens(t, input, []Token{
		{Type: TokenPackageDecl, Value: "package views"},
		{Type: TokenImportDecl, Value: `import "fmt"`},
		{Type: TokenImportDecl, Value: `import "strings"`},
		{Type: TokenEOF},
	})
}

func TestGoxDecl(t *testing.T) {
	input := "func Hello(name string) {\n}"
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl, Value: "func Hello(name string) {"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose, Value: "}"},
		{Type: TokenEOF},
	})
}

func TestGoxDeclMultiLineParams(t *testing.T) {
	input := "func Hello(\n  name string,\n  age int,\n) {\n}"
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl, Value: "func Hello(\n  name string,\n  age int,\n) {"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose, Value: "}"},
		{Type: TokenEOF},
	})
}

func TestSimpleHTMLText(t *testing.T) {
	input := "func Hello() {\n  Hello, world!\n}"
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl, Value: "func Hello() {"},
		{Type: TokenHTMLText, Value: "\n  Hello, world!\n"},
		{Type: TokenBraceClose, Value: "}"},
		{Type: TokenEOF},
	})
}

func TestOpenAndCloseTag(t *testing.T) {
	input := "func Hello() {\n<h1>Hello</h1>\n}"
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<h1"},
		{Type: TokenTagEnd, Value: ">"},
		{Type: TokenHTMLText, Value: "Hello"},
		{Type: TokenCloseTag, Value: "</h1>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestSelfClosingTag(t *testing.T) {
	input := "func Hello() {\n<br />\n}"
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<br"},
		{Type: TokenSelfCloseTag, Value: "/>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestExpressionEscaped(t *testing.T) {
	input := "func Hello(name string) {\n<p>{{ name }}</p>\n}"
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<p"},
		{Type: TokenTagEnd},
		{Type: TokenExprOpen, Value: "{{"},
		{Type: TokenGoCode, Value: "name"},
		{Type: TokenExprClose, Value: "}}"},
		{Type: TokenCloseTag, Value: "</p>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestRawExpression(t *testing.T) {
	input := "func Hello() {\n<div>{{! rawHTML }}</div>\n}"
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<div"},
		{Type: TokenTagEnd},
		{Type: TokenRawExprOpen, Value: "{{!"},
		{Type: TokenGoCode, Value: "rawHTML"},
		{Type: TokenExprClose, Value: "}}"},
		{Type: TokenCloseTag, Value: "</div>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestStaticAttribute(t *testing.T) {
	input := `func Hello() {
<div class="main">text</div>
}`
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<div"},
		{Type: TokenAttrName, Value: "class="},
		{Type: TokenAttrValue, Value: `"main"`},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "text"},
		{Type: TokenCloseTag, Value: "</div>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestDynamicAttribute(t *testing.T) {
	input := "func Hello(cls string) {\n<div class={{ cls }}>text</div>\n}"
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<div"},
		{Type: TokenAttrName, Value: "class="},
		{Type: TokenExprOpen, Value: "{{"},
		{Type: TokenGoCode, Value: "cls"},
		{Type: TokenExprClose, Value: "}}"},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "text"},
		{Type: TokenCloseTag, Value: "</div>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestBooleanAttribute(t *testing.T) {
	input := "func Hello() {\n<button disabled>Submit</button>\n}"
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<button"},
		{Type: TokenAttrName, Value: "disabled"},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "Submit"},
		{Type: TokenCloseTag, Value: "</button>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestSpreadAttribute(t *testing.T) {
	input := "func Hello() {\n<input {{ attrs... }} />\n}"
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<input"},
		{Type: TokenExprOpen, Value: "{{"},
		{Type: TokenSpreadOp, Value: "attrs..."},
		{Type: TokenExprClose, Value: "}}"},
		{Type: TokenSelfCloseTag, Value: "/>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestControlFlowIf(t *testing.T) {
	input := `func Hello(isAdmin bool) {
{{ if isAdmin }}
  <span>Admin</span>
{{ else }}
  <span>User</span>
{{ end }}
}`
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: "if isAdmin"},
		{Type: TokenExprClose},
		{Type: TokenHTMLText, Value: "\n  "},
		{Type: TokenOpenTag, Value: "<span"},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "Admin"},
		{Type: TokenCloseTag, Value: "</span>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: "else"},
		{Type: TokenExprClose},
		{Type: TokenHTMLText, Value: "\n  "},
		{Type: TokenOpenTag, Value: "<span"},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "User"},
		{Type: TokenCloseTag, Value: "</span>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: "end"},
		{Type: TokenExprClose},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestControlFlowFor(t *testing.T) {
	input := `func List(items []string) {
<ul>
  {{ for _, item := range items }}
    <li>{{ item }}</li>
  {{ end }}
</ul>
}`
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<ul"},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "\n  "},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: "for _, item := range items"},
		{Type: TokenExprClose},
		{Type: TokenHTMLText, Value: "\n    "},
		{Type: TokenOpenTag, Value: "<li"},
		{Type: TokenTagEnd},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: "item"},
		{Type: TokenExprClose},
		{Type: TokenCloseTag, Value: "</li>"},
		{Type: TokenHTMLText, Value: "\n  "},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: "end"},
		{Type: TokenExprClose},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenCloseTag, Value: "</ul>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestControlFlowSwitch(t *testing.T) {
	input := `func Status(status string) {
{{ switch status }}
  {{ case "active" }}
    <span>Active</span>
  {{ default }}
    <span>Unknown</span>
{{ end }}
}`
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: "switch status"},
		{Type: TokenExprClose},
		{Type: TokenHTMLText, Value: "\n  "},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: `case "active"`},
		{Type: TokenExprClose},
		{Type: TokenHTMLText, Value: "\n    "},
		{Type: TokenOpenTag, Value: "<span"},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "Active"},
		{Type: TokenCloseTag, Value: "</span>"},
		{Type: TokenHTMLText, Value: "\n  "},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: "default"},
		{Type: TokenExprClose},
		{Type: TokenHTMLText, Value: "\n    "},
		{Type: TokenOpenTag, Value: "<span"},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "Unknown"},
		{Type: TokenCloseTag, Value: "</span>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: "end"},
		{Type: TokenExprClose},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestChildren(t *testing.T) {
	input := `func Card(title string) {
<div>
  {{ children }}
</div>
}`
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<div"},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "\n  "},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: "children"},
		{Type: TokenExprClose},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenCloseTag, Value: "</div>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestComponentCall(t *testing.T) {
	input := `func Page() {
<Header />
<Footer />
}`
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<Header"},
		{Type: TokenSelfCloseTag, Value: "/>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<Footer"},
		{Type: TokenSelfCloseTag, Value: "/>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestMultipleAttributes(t *testing.T) {
	input := `func Hello() {
<div id="app" class="main" data-x="1">text</div>
}`
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<div"},
		{Type: TokenAttrName, Value: "id="},
		{Type: TokenAttrValue, Value: `"app"`},
		{Type: TokenAttrName, Value: "class="},
		{Type: TokenAttrValue, Value: `"main"`},
		{Type: TokenAttrName, Value: "data-x="},
		{Type: TokenAttrValue, Value: `"1"`},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "text"},
		{Type: TokenCloseTag, Value: "</div>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestExpressionWithBraces(t *testing.T) {
	input := `func Hello() {
{{ map[string]int{"a": 1} }}
}`
	tokens := tokenize(input)
	// Find the GoCode token
	for _, tok := range tokens {
		if tok.Type == TokenGoCode {
			if tok.Value != `map[string]int{"a": 1}` {
				t.Errorf("GoCode value: got %q, want %q", tok.Value, `map[string]int{"a": 1}`)
			}
			return
		}
	}
	t.Error("no GoCode token found")
}

func TestExpressionWithStringContainingBraces(t *testing.T) {
	input := "func Hello() {\n{{ fmt.Sprintf(\"}}\") }}\n}"
	tokens := tokenize(input)
	for _, tok := range tokens {
		if tok.Type == TokenGoCode {
			if tok.Value != `fmt.Sprintf("}}")` {
				t.Errorf("GoCode value: got %q, want %q", tok.Value, `fmt.Sprintf("}}")`)
			}
			return
		}
	}
	t.Error("no GoCode token found")
}

func TestMultipleComponents(t *testing.T) {
	input := `package views

func Hello() {
<h1>Hello</h1>
}

func World() {
<h2>World</h2>
}`
	assertTokens(t, input, []Token{
		{Type: TokenPackageDecl, Value: "package views"},
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<h1"},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "Hello"},
		{Type: TokenCloseTag, Value: "</h1>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<h2"},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "World"},
		{Type: TokenCloseTag, Value: "</h2>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestLineColumnTracking(t *testing.T) {
	input := "package views\n\nfunc Hello() {\n<h1>Hi</h1>\n}"
	tokens := tokenize(input)

	// package views is at line 1, col 1
	if tokens[0].Line != 1 || tokens[0].Col != 1 {
		t.Errorf("PackageDecl position: got %d:%d, want 1:1", tokens[0].Line, tokens[0].Col)
	}

	// func Hello() { is at line 3
	if tokens[1].Line != 3 {
		t.Errorf("GoxDecl line: got %d, want 3", tokens[1].Line)
	}
}

func TestFullComponent(t *testing.T) {
	input := `package views

import "fmt"

func Hello(name string) {
  <h1 class="title">Hello, {{ name }}!</h1>
}`
	assertTokens(t, input, []Token{
		{Type: TokenPackageDecl, Value: "package views"},
		{Type: TokenImportDecl, Value: `import "fmt"`},
		{Type: TokenGoxDecl, Value: "func Hello(name string) {"},
		{Type: TokenHTMLText, Value: "\n  "},
		{Type: TokenOpenTag, Value: "<h1"},
		{Type: TokenAttrName, Value: "class="},
		{Type: TokenAttrValue, Value: `"title"`},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "Hello, "},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: "name"},
		{Type: TokenExprClose},
		{Type: TokenHTMLText, Value: "!"},
		{Type: TokenCloseTag, Value: "</h1>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestDynamicBooleanAttribute(t *testing.T) {
	input := "func Hello() {\n<button disabled={{ isDisabled }}>Submit</button>\n}"
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<button"},
		{Type: TokenAttrName, Value: "disabled="},
		{Type: TokenExprOpen, Value: "{{"},
		{Type: TokenGoCode, Value: "isDisabled"},
		{Type: TokenExprClose, Value: "}}"},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "Submit"},
		{Type: TokenCloseTag, Value: "</button>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestComponentWithChildren(t *testing.T) {
	input := `func Page() {
<Card title="Welcome">
  <p>Hello</p>
</Card>
}`
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<Card"},
		{Type: TokenAttrName, Value: "title="},
		{Type: TokenAttrValue, Value: `"Welcome"`},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "\n  "},
		{Type: TokenOpenTag, Value: "<p"},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "Hello"},
		{Type: TokenCloseTag, Value: "</p>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenCloseTag, Value: "</Card>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestEmptyComponent(t *testing.T) {
	input := "func Empty() {\n}"
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl, Value: "func Empty() {"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestExpressionWithComplexGoCode(t *testing.T) {
	input := `func Hello(user User) {
<a href={{ fmt.Sprintf("/users/%d", user.ID) }}>Profile</a>
}`
	assertTokens(t, input, []Token{
		{Type: TokenGoxDecl},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenOpenTag, Value: "<a"},
		{Type: TokenAttrName, Value: "href="},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: `fmt.Sprintf("/users/%d", user.ID)`},
		{Type: TokenExprClose},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "Profile"},
		{Type: TokenCloseTag, Value: "</a>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestConditionalBooleanAttribute(t *testing.T) {
	input := `package views

func Button(isDisabled bool) {
  <button disabled?={{ isDisabled }}>Submit</button>
}`
	assertTokens(t, input, []Token{
		{Type: TokenPackageDecl, Value: "package views"},
		{Type: TokenGoxDecl, Value: "func Button(isDisabled bool) {"},
		{Type: TokenHTMLText},
		{Type: TokenOpenTag, Value: "<button"},
		{Type: TokenAttrName, Value: "disabled?="},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: "isDisabled"},
		{Type: TokenExprClose},
		{Type: TokenTagEnd},
		{Type: TokenHTMLText, Value: "Submit"},
		{Type: TokenCloseTag, Value: "</button>"},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}

func TestConditionalBooleanSelfClosing(t *testing.T) {
	input := `package views

func Form(checked bool) {
  <input type="checkbox" checked?={{ checked }} />
}`
	assertTokens(t, input, []Token{
		{Type: TokenPackageDecl, Value: "package views"},
		{Type: TokenGoxDecl, Value: "func Form(checked bool) {"},
		{Type: TokenHTMLText},
		{Type: TokenOpenTag, Value: "<input"},
		{Type: TokenAttrName, Value: "type="},
		{Type: TokenAttrValue, Value: `"checkbox"`},
		{Type: TokenAttrName, Value: "checked?="},
		{Type: TokenExprOpen},
		{Type: TokenGoCode, Value: "checked"},
		{Type: TokenExprClose},
		{Type: TokenSelfCloseTag},
		{Type: TokenHTMLText, Value: "\n"},
		{Type: TokenBraceClose},
		{Type: TokenEOF},
	})
}
