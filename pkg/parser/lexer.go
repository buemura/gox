package parser

import "strings"

type lexerState int

const (
	stateTopLevel   lexerState = iota // package, import, func declarations
	stateTemplate                     // inside func body: HTML text, tags, expressions
	stateTag                          // inside an open tag: attributes, >, />
	stateExpression                   // inside {{ }}: reading Go code
	stateAttrValue                    // after AttrName with =: expects "string" or {{ expr }}
)

// Lexer tokenizes .gox file content into a stream of tokens.
type Lexer struct {
	input       string
	pos         int
	line        int
	col         int
	state       lexerState
	returnState lexerState // state to return to after expression closes
	braceDepth  int        // for tracking gox body nesting
}

// NewLexer creates a new Lexer for the given input.
func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
		pos:   0,
		line:  1,
		col:   1,
		state: stateTopLevel,
	}
}

// NextToken returns the next token from the input.
func (l *Lexer) NextToken() Token {
	if l.isAtEnd() {
		return Token{Type: TokenEOF, Line: l.line, Col: l.col}
	}
	switch l.state {
	case stateTopLevel:
		return l.lexTopLevel()
	case stateTemplate:
		return l.lexTemplate()
	case stateTag:
		return l.lexTag()
	case stateExpression:
		return l.lexExpression()
	case stateAttrValue:
		return l.lexAttrValue()
	default:
		return Token{Type: TokenEOF, Line: l.line, Col: l.col}
	}
}

// --- Character helpers ---

func (l *Lexer) isAtEnd() bool {
	return l.pos >= len(l.input)
}

func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}
	return l.input[l.pos]
}

func (l *Lexer) peekAt(offset int) byte {
	idx := l.pos + offset
	if idx >= len(l.input) {
		return 0
	}
	return l.input[idx]
}

func (l *Lexer) advance() byte {
	if l.isAtEnd() {
		return 0
	}
	ch := l.input[l.pos]
	l.pos++
	if ch == '\n' {
		l.line++
		l.col = 1
	} else {
		l.col++
	}
	return ch
}

func (l *Lexer) skipWhitespace() {
	for !l.isAtEnd() && isWhitespace(l.peek()) {
		l.advance()
	}
}

func (l *Lexer) startsWith(s string) bool {
	return l.pos+len(s) <= len(l.input) && l.input[l.pos:l.pos+len(s)] == s
}

func (l *Lexer) readWhile(pred func(byte) bool) string {
	start := l.pos
	for !l.isAtEnd() && pred(l.peek()) {
		l.advance()
	}
	return l.input[start:l.pos]
}

func (l *Lexer) readUntilEOL() string {
	start := l.pos
	for !l.isAtEnd() && l.peek() != '\n' {
		l.advance()
	}
	return l.input[start:l.pos]
}

// --- Top-level scanning ---

func (l *Lexer) lexTopLevel() Token {
	l.skipWhitespace()
	if l.isAtEnd() {
		return Token{Type: TokenEOF, Line: l.line, Col: l.col}
	}

	line, col := l.line, l.col

	if l.startsWith("package ") {
		value := l.readUntilEOL()
		return Token{Type: TokenPackageDecl, Value: value, Line: line, Col: col}
	}

	if l.startsWith("import ") {
		value := l.readUntilEOL()
		return Token{Type: TokenImportDecl, Value: value, Line: line, Col: col}
	}

	if l.startsWith("func ") {
		return l.lexGoxDecl(line, col)
	}

	// Skip unknown characters at top level
	l.advance()
	return l.lexTopLevel()
}

func (l *Lexer) lexGoxDecl(line, col int) Token {
	start := l.pos
	parenDepth := 0

	for !l.isAtEnd() {
		ch := l.peek()
		if ch == '(' {
			parenDepth++
			l.advance()
		} else if ch == ')' {
			parenDepth--
			l.advance()
		} else if ch == '{' && parenDepth == 0 {
			l.advance() // consume the '{'
			value := l.input[start:l.pos]
			l.state = stateTemplate
			l.braceDepth = 1
			return Token{Type: TokenGoxDecl, Value: strings.TrimRight(value, " \t"), Line: line, Col: col}
		} else {
			l.advance()
		}
	}

	// Unterminated gox declaration
	return Token{Type: TokenGoxDecl, Value: l.input[start:l.pos], Line: line, Col: col}
}

// --- Template mode ---

func (l *Lexer) lexTemplate() Token {
	if l.isAtEnd() {
		return Token{Type: TokenEOF, Line: l.line, Col: l.col}
	}

	line, col := l.line, l.col

	// Check for expression open: {{! or {{
	if l.startsWith("{{!") {
		l.pos += 3
		l.col += 3
		l.returnState = stateTemplate
		l.state = stateExpression
		return Token{Type: TokenRawExprOpen, Value: "{{!", Line: line, Col: col}
	}
	if l.startsWith("{{") {
		l.pos += 2
		l.col += 2
		l.returnState = stateTemplate
		l.state = stateExpression
		return Token{Type: TokenExprOpen, Value: "{{", Line: line, Col: col}
	}

	// Check for close tag: </
	if l.startsWith("</") {
		return l.lexCloseTag(line, col)
	}

	// Check for open tag: < followed by letter
	if l.peek() == '<' && l.peekAt(1) != 0 && isLetter(l.peekAt(1)) {
		return l.lexOpenTag(line, col)
	}

	// Check for gox body close
	if l.peek() == '}' {
		l.braceDepth--
		if l.braceDepth == 0 {
			l.advance()
			l.state = stateTopLevel
			return Token{Type: TokenBraceClose, Value: "}", Line: line, Col: col}
		}
	}

	// Read HTML text until we hit a delimiter
	return l.lexHTMLText(line, col)
}

func (l *Lexer) lexHTMLText(line, col int) Token {
	start := l.pos
	for !l.isAtEnd() {
		if l.startsWith("{{") || l.startsWith("</") {
			break
		}
		if l.peek() == '<' && l.peekAt(1) != 0 && isLetter(l.peekAt(1)) {
			break
		}
		if l.peek() == '}' && l.braceDepth == 1 {
			break
		}
		l.advance()
	}

	value := l.input[start:l.pos]
	if value == "" {
		// Don't emit empty text tokens; try next token
		return l.lexTemplate()
	}
	return Token{Type: TokenHTMLText, Value: value, Line: line, Col: col}
}

func (l *Lexer) lexOpenTag(line, col int) Token {
	l.advance() // consume '<'
	name := l.readWhile(isTagNameChar)
	l.state = stateTag
	return Token{Type: TokenOpenTag, Value: "<" + name, Line: line, Col: col}
}

func (l *Lexer) lexCloseTag(line, col int) Token {
	l.advance() // consume '<'
	l.advance() // consume '/'
	name := l.readWhile(isTagNameChar)
	// consume '>'
	if !l.isAtEnd() && l.peek() == '>' {
		l.advance()
	}
	return Token{Type: TokenCloseTag, Value: "</" + name + ">", Line: line, Col: col}
}

// --- Tag mode ---

func (l *Lexer) lexTag() Token {
	l.skipWhitespace()
	if l.isAtEnd() {
		return Token{Type: TokenEOF, Line: l.line, Col: l.col}
	}

	line, col := l.line, l.col

	// Self-closing tag
	if l.startsWith("/>") {
		l.pos += 2
		l.col += 2
		l.state = stateTemplate
		return Token{Type: TokenSelfCloseTag, Value: "/>", Line: line, Col: col}
	}

	// Tag end
	if l.peek() == '>' {
		l.advance()
		l.state = stateTemplate
		return Token{Type: TokenTagEnd, Value: ">", Line: line, Col: col}
	}

	// Expression open (for spread attributes)
	if l.startsWith("{{") {
		l.pos += 2
		l.col += 2
		l.returnState = stateTag
		l.state = stateExpression
		return Token{Type: TokenExprOpen, Value: "{{", Line: line, Col: col}
	}

	// Attribute name
	if isLetter(l.peek()) || l.peek() == '_' {
		return l.lexAttrName(line, col)
	}

	// Skip unexpected characters
	l.advance()
	return l.lexTag()
}

func (l *Lexer) lexAttrName(line, col int) Token {
	name := l.readWhile(isAttrNameChar)

	if !l.isAtEnd() && l.peek() == '=' {
		l.advance() // consume '='
		l.state = stateAttrValue
		return Token{Type: TokenAttrName, Value: name + "=", Line: line, Col: col}
	}

	// Boolean attribute (no =)
	return Token{Type: TokenAttrName, Value: name, Line: line, Col: col}
}

// --- Attribute value mode ---

func (l *Lexer) lexAttrValue() Token {
	l.skipWhitespace()
	if l.isAtEnd() {
		l.state = stateTag
		return Token{Type: TokenEOF, Line: l.line, Col: l.col}
	}

	line, col := l.line, l.col

	// Quoted string value
	if l.peek() == '"' {
		value := l.readQuotedString()
		l.state = stateTag
		return Token{Type: TokenAttrValue, Value: value, Line: line, Col: col}
	}

	// Dynamic attribute value: {{ expr }}
	if l.startsWith("{{") {
		l.pos += 2
		l.col += 2
		l.returnState = stateTag
		l.state = stateExpression
		return Token{Type: TokenExprOpen, Value: "{{", Line: line, Col: col}
	}

	// Unquoted attribute value (read until whitespace or >)
	value := l.readWhile(func(b byte) bool {
		return !isWhitespace(b) && b != '>' && b != '/'
	})
	l.state = stateTag
	return Token{Type: TokenAttrValue, Value: value, Line: line, Col: col}
}

func (l *Lexer) readQuotedString() string {
	var sb strings.Builder
	sb.WriteByte(l.advance()) // consume opening quote

	for !l.isAtEnd() {
		ch := l.peek()
		if ch == '\\' && l.peekAt(1) == '"' {
			sb.WriteByte(l.advance()) // consume '\'
			sb.WriteByte(l.advance()) // consume '"'
			continue
		}
		if ch == '"' {
			sb.WriteByte(l.advance()) // consume closing quote
			break
		}
		sb.WriteByte(l.advance())
	}

	return sb.String()
}

// --- Expression mode ---

func (l *Lexer) lexExpression() Token {
	l.skipWhitespace()
	if l.isAtEnd() {
		return Token{Type: TokenEOF, Line: l.line, Col: l.col}
	}

	line, col := l.line, l.col

	// Check for immediate close
	if l.startsWith("}}") {
		l.pos += 2
		l.col += 2
		l.state = l.returnState
		return Token{Type: TokenExprClose, Value: "}}", Line: line, Col: col}
	}

	// Read Go code until }}
	code := l.readGoCode()
	trimmed := strings.TrimSpace(code)

	if strings.HasSuffix(trimmed, "...") {
		return Token{Type: TokenSpreadOp, Value: trimmed, Line: line, Col: col}
	}

	return Token{Type: TokenGoCode, Value: trimmed, Line: line, Col: col}
}

func (l *Lexer) readGoCode() string {
	start := l.pos
	braceDepth := 0

	for !l.isAtEnd() {
		// Check for }} at brace depth 0
		if l.startsWith("}}") && braceDepth == 0 {
			break
		}

		ch := l.peek()

		// Track braces inside Go code
		if ch == '{' {
			braceDepth++
			l.advance()
			continue
		}
		if ch == '}' {
			braceDepth--
			l.advance()
			continue
		}

		// Skip over string literals
		if ch == '"' {
			l.skipDoubleQuotedString()
			continue
		}
		if ch == '\'' {
			l.skipSingleQuotedString()
			continue
		}
		if ch == '`' {
			l.skipRawString()
			continue
		}

		l.advance()
	}

	return l.input[start:l.pos]
}

func (l *Lexer) skipDoubleQuotedString() {
	l.advance() // consume opening "
	for !l.isAtEnd() {
		ch := l.peek()
		if ch == '\\' {
			l.advance() // consume '\'
			l.advance() // consume escaped char
			continue
		}
		if ch == '"' {
			l.advance() // consume closing "
			return
		}
		l.advance()
	}
}

func (l *Lexer) skipSingleQuotedString() {
	l.advance() // consume opening '
	for !l.isAtEnd() {
		ch := l.peek()
		if ch == '\\' {
			l.advance()
			l.advance()
			continue
		}
		if ch == '\'' {
			l.advance()
			return
		}
		l.advance()
	}
}

func (l *Lexer) skipRawString() {
	l.advance() // consume opening `
	for !l.isAtEnd() {
		if l.peek() == '`' {
			l.advance()
			return
		}
		l.advance()
	}
}

// --- Character classification ---

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isTagNameChar(ch byte) bool {
	return isLetter(ch) || (ch >= '0' && ch <= '9') || ch == '-'
}

func isAttrNameChar(ch byte) bool {
	return isLetter(ch) || (ch >= '0' && ch <= '9') || ch == '-' || ch == ':'
}
