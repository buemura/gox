package parser

import "fmt"

// TokenType represents the type of a lexical token.
type TokenType int

const (
	TokenEOF          TokenType = iota
	TokenPackageDecl            // "package views"
	TokenImportDecl             // `import "fmt"`
	TokenGoxDecl                // "gox Hello(name string) {"
	TokenHTMLText               // raw text content
	TokenExprOpen               // "{{"
	TokenExprClose              // "}}"
	TokenRawExprOpen            // "{{!"
	TokenGoCode                 // Go expression inside {{ }}
	TokenOpenTag                // "<div", "<Header"
	TokenCloseTag               // "</div>", "</Header>"
	TokenSelfCloseTag           // "/>"
	TokenTagEnd                 // ">"
	TokenAttrName               // "class=" or "disabled" (boolean)
	TokenAttrValue              // `"main"` (quoted string value)
	TokenSpreadOp               // "attrs..."
	TokenBraceOpen              // "{"
	TokenBraceClose             // "}"
)

var tokenNames = map[TokenType]string{
	TokenEOF:          "EOF",
	TokenPackageDecl:  "PackageDecl",
	TokenImportDecl:   "ImportDecl",
	TokenGoxDecl:      "GoxDecl",
	TokenHTMLText:     "HTMLText",
	TokenExprOpen:     "ExprOpen",
	TokenExprClose:    "ExprClose",
	TokenRawExprOpen:  "RawExprOpen",
	TokenGoCode:       "GoCode",
	TokenOpenTag:      "OpenTag",
	TokenCloseTag:     "CloseTag",
	TokenSelfCloseTag: "SelfCloseTag",
	TokenTagEnd:       "TagEnd",
	TokenAttrName:     "AttrName",
	TokenAttrValue:    "AttrValue",
	TokenSpreadOp:     "SpreadOp",
	TokenBraceOpen:    "BraceOpen",
	TokenBraceClose:   "BraceClose",
}

func (t TokenType) String() string {
	if name, ok := tokenNames[t]; ok {
		return name
	}
	return fmt.Sprintf("TokenType(%d)", int(t))
}

// Token represents a single lexical token.
type Token struct {
	Type  TokenType
	Value string
	Line  int
	Col   int
}

func (t Token) String() string {
	return fmt.Sprintf("%s(%q) at %d:%d", t.Type, t.Value, t.Line, t.Col)
}
