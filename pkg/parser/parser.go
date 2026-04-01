package parser

import (
	"fmt"
	"strings"
	"unicode"
)

// ParseError represents a parsing error with source location.
type ParseError struct {
	Message string
	Line    int
	Col     int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error at %d:%d: %s", e.Line, e.Col, e.Message)
}

// Parser consumes a token stream from the Lexer and builds an AST.
type Parser struct {
	lexer   *Lexer
	current Token
	peeked  []Token // lookahead buffer
	errors  []ParseError
}

// NewParser creates a new parser for the given input.
func NewParser(input string) *Parser {
	p := &Parser{
		lexer: NewLexer(input),
	}
	p.advance()
	return p
}

// Parse parses the entire input and returns a File AST node.
func (p *Parser) Parse() (*File, error) {
	file := &File{}

	if p.current.Type == TokenPackageDecl {
		file.Package = p.parsePackageDecl()
	} else {
		return nil, p.errorf("expected package declaration")
	}

	for p.current.Type == TokenImportDecl {
		file.Imports = append(file.Imports, p.parseImportDecl())
	}

	for p.current.Type == TokenGoxDecl {
		comp, err := p.parseComponentDecl()
		if err != nil {
			return nil, err
		}
		file.Components = append(file.Components, comp)
	}

	if p.current.Type != TokenEOF {
		return nil, p.errorf("unexpected token %s", p.current.Type)
	}

	return file, nil
}

// advance moves to the next token.
func (p *Parser) advance() Token {
	prev := p.current
	if len(p.peeked) > 0 {
		p.current = p.peeked[0]
		p.peeked = p.peeked[1:]
	} else {
		p.current = p.lexer.NextToken()
	}
	return prev
}

// peek returns the next token without consuming it.
func (p *Parser) peek() Token {
	if len(p.peeked) == 0 {
		p.peeked = append(p.peeked, p.lexer.NextToken())
	}
	return p.peeked[0]
}

func (p *Parser) errorf(format string, args ...any) error {
	return &ParseError{
		Message: fmt.Sprintf(format, args...),
		Line:    p.current.Line,
		Col:     p.current.Col,
	}
}

// isAtKeyword checks if the parser is currently at {{ keyword }} where keyword
// is one of the control flow terminators (end, else, case, default).
func (p *Parser) isAtKeyword(keywords ...string) bool {
	if p.current.Type != TokenExprOpen {
		return false
	}
	next := p.peek()
	if next.Type != TokenGoCode {
		return false
	}
	code := strings.TrimSpace(next.Value)
	for _, kw := range keywords {
		if code == kw || strings.HasPrefix(code, kw+" ") {
			return true
		}
	}
	return false
}

// consumeEnd consumes a {{ end }} sequence.
func (p *Parser) consumeEnd() {
	if p.current.Type == TokenExprOpen {
		next := p.peek()
		if next.Type == TokenGoCode && strings.TrimSpace(next.Value) == "end" {
			p.advance() // ExprOpen
			p.advance() // GoCode("end")
			if p.current.Type == TokenExprClose {
				p.advance()
			}
		}
	}
}

// --- Top-level parsers ---

func (p *Parser) parsePackageDecl() *PackageDecl {
	tok := p.advance()
	name := strings.TrimPrefix(tok.Value, "package ")
	return &PackageDecl{Name: strings.TrimSpace(name), Line: tok.Line, Col: tok.Col}
}

func (p *Parser) parseImportDecl() *ImportDecl {
	tok := p.advance()
	path := strings.TrimPrefix(tok.Value, "import ")
	return &ImportDecl{Path: strings.TrimSpace(path), Line: tok.Line, Col: tok.Col}
}

func (p *Parser) parseComponentDecl() (*ComponentDecl, error) {
	tok := p.advance()
	name, params := parseGoxSignature(tok.Value)

	comp := &ComponentDecl{
		Name:   name,
		Params: params,
		Line:   tok.Line,
		Col:    tok.Col,
	}

	body, err := p.parseBody("end")
	if err != nil {
		return nil, err
	}
	comp.Body = body

	if p.current.Type != TokenBraceClose {
		return nil, p.errorf("unclosed component %q starting at %d:%d", name, tok.Line, tok.Col)
	}
	p.advance() // consume BraceClose

	return comp, nil
}

func parseGoxSignature(value string) (name, params string) {
	value = strings.TrimPrefix(value, "func ")
	value = strings.TrimSuffix(value, "{")
	value = strings.TrimSpace(value)

	if idx := strings.Index(value, "("); idx != -1 {
		name = value[:idx]
		rest := value[idx+1:]
		if ci := strings.LastIndex(rest, ")"); ci != -1 {
			params = strings.TrimSpace(rest[:ci])
		}
	} else {
		name = value
	}
	return
}

// --- Body parser ---

// parseBody parses template nodes until a terminator is encountered.
// It stops (without consuming) when it sees: BraceClose, CloseTag, EOF,
// or {{ keyword }} where keyword is "end", "else", "case", or "default".
func (p *Parser) parseBody(stopAt ...string) ([]Node, error) {
	stopKeywords := []string{"end", "else", "case", "default"}
	_ = stopAt

	var nodes []Node
	for {
		// Check terminators
		if p.current.Type == TokenEOF || p.current.Type == TokenBraceClose || p.current.Type == TokenCloseTag {
			return nodes, nil
		}

		// Check for control flow keyword terminators
		if p.isAtKeyword(stopKeywords...) {
			return nodes, nil
		}

		switch p.current.Type {
		case TokenHTMLText:
			nodes = append(nodes, p.parseTextNode())
		case TokenOpenTag:
			node, err := p.parseElement()
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, node)
		case TokenExprOpen:
			node, err := p.parseExprBlock()
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, node)
		case TokenRawExprOpen:
			node, err := p.parseRawExprBlock()
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, node)
		default:
			return nil, p.errorf("unexpected token %s in body", p.current.Type)
		}
	}
}

func (p *Parser) parseTextNode() *TextNode {
	tok := p.advance()
	return &TextNode{Content: tok.Value, Line: tok.Line, Col: tok.Col}
}

// --- Element parser ---

func extractTagName(value string) string {
	name := strings.TrimPrefix(value, "</")
	name = strings.TrimPrefix(name, "<")
	name = strings.TrimSuffix(name, ">")
	return name
}

func isComponentName(name string) bool {
	return len(name) > 0 && unicode.IsUpper(rune(name[0]))
}

var voidElements = map[string]bool{
	"area": true, "base": true, "br": true, "col": true,
	"embed": true, "hr": true, "img": true, "input": true,
	"link": true, "meta": true, "param": true, "source": true,
	"track": true, "wbr": true,
}

func (p *Parser) parseElement() (Node, error) {
	tok := p.advance() // consume OpenTag
	tagName := extractTagName(tok.Value)
	line, col := tok.Line, tok.Col

	attrs, err := p.parseAttributes()
	if err != nil {
		return nil, err
	}

	selfClose := false
	if p.current.Type == TokenSelfCloseTag {
		selfClose = true
		p.advance()
	} else if p.current.Type == TokenTagEnd {
		p.advance()
	}

	if isComponentName(tagName) {
		return p.buildComponentCall(tagName, attrs, selfClose, line, col)
	}
	return p.buildHTMLElement(tagName, attrs, selfClose, line, col)
}

func (p *Parser) buildHTMLElement(tag string, attrs []Attribute, selfClose bool, line, col int) (*HTMLElement, error) {
	elem := &HTMLElement{
		Tag:        tag,
		Attributes: attrs,
		SelfClose:  selfClose,
		Line:       line,
		Col:        col,
	}

	if selfClose || voidElements[tag] {
		return elem, nil
	}

	children, err := p.parseBody()
	if err != nil {
		return nil, err
	}
	elem.Children = children

	if p.current.Type != TokenCloseTag {
		return nil, &ParseError{
			Message: fmt.Sprintf("expected closing tag </%s>", tag),
			Line:    p.current.Line,
			Col:     p.current.Col,
		}
	}
	closeTok := p.advance()
	closeTag := extractTagName(closeTok.Value)
	if closeTag != tag {
		return nil, &ParseError{
			Message: fmt.Sprintf("mismatched closing tag: expected </%s>, got </%s>", tag, closeTag),
			Line:    closeTok.Line,
			Col:     closeTok.Col,
		}
	}

	return elem, nil
}

func (p *Parser) buildComponentCall(name string, attrs []Attribute, selfClose bool, line, col int) (*ComponentCall, error) {
	comp := &ComponentCall{
		Name:       name,
		Attributes: attrs,
		SelfClose:  selfClose,
		Line:       line,
		Col:        col,
	}

	if selfClose {
		return comp, nil
	}

	children, err := p.parseBody()
	if err != nil {
		return nil, err
	}
	comp.Children = children

	if p.current.Type != TokenCloseTag {
		return nil, &ParseError{
			Message: fmt.Sprintf("expected closing tag </%s>", name),
			Line:    p.current.Line,
			Col:     p.current.Col,
		}
	}
	closeTok := p.advance()
	closeTag := extractTagName(closeTok.Value)
	if closeTag != name {
		return nil, &ParseError{
			Message: fmt.Sprintf("mismatched closing tag: expected </%s>, got </%s>", name, closeTag),
			Line:    closeTok.Line,
			Col:     closeTok.Col,
		}
	}

	return comp, nil
}

// --- Attribute parsers ---

func (p *Parser) parseAttributes() ([]Attribute, error) {
	var attrs []Attribute

	for p.current.Type != TokenTagEnd && p.current.Type != TokenSelfCloseTag && p.current.Type != TokenEOF {
		switch p.current.Type {
		case TokenAttrName:
			attr, err := p.parseAttribute()
			if err != nil {
				return nil, err
			}
			attrs = append(attrs, attr)
		case TokenExprOpen:
			// Spread attributes: {{ attrs... }}
			p.advance() // consume ExprOpen
			if p.current.Type == TokenGoCode || p.current.Type == TokenSpreadOp {
				tok := p.advance()
				val := strings.TrimSpace(tok.Value)
				attrs = append(attrs, Attribute{
					Name:   val,
					Spread: true,
					Line:   tok.Line,
					Col:    tok.Col,
				})
				if p.current.Type == TokenExprClose {
					p.advance()
				}
			}
		default:
			return nil, p.errorf("unexpected token %s in tag attributes", p.current.Type)
		}
	}

	return attrs, nil
}

func (p *Parser) parseAttribute() (Attribute, error) {
	nameTok := p.advance()
	attrName := nameTok.Value
	line, col := nameTok.Line, nameTok.Col

	// Conditional boolean: name?={{ expr }}
	isBooleanCond := strings.HasSuffix(attrName, "?=")
	if isBooleanCond {
		attrName = strings.TrimSuffix(attrName, "?=")
		if p.current.Type == TokenExprOpen {
			p.advance() // consume ExprOpen
			if p.current.Type == TokenGoCode {
				codeTok := p.advance()
				if p.current.Type == TokenExprClose {
					p.advance()
				}
				return Attribute{
					Name:    attrName,
					Value:   strings.TrimSpace(codeTok.Value),
					Dynamic: true,
					Boolean: true,
					Line:    line,
					Col:     col,
				}, nil
			}
			return Attribute{}, p.errorf("expected Go expression in conditional boolean attribute")
		}
		return Attribute{}, p.errorf("conditional boolean attribute %q requires a dynamic expression (e.g., %s?={{ expr }})", attrName, attrName)
	}

	hasValue := strings.HasSuffix(attrName, "=")
	if hasValue {
		attrName = strings.TrimSuffix(attrName, "=")
	} else {
		return Attribute{Name: attrName, Boolean: true, Line: line, Col: col}, nil
	}

	switch p.current.Type {
	case TokenAttrValue:
		valTok := p.advance()
		val := valTok.Value
		if len(val) >= 2 && val[0] == '"' && val[len(val)-1] == '"' {
			val = val[1 : len(val)-1]
		}
		return Attribute{Name: attrName, Value: val, Line: line, Col: col}, nil
	case TokenExprOpen:
		p.advance() // consume ExprOpen
		if p.current.Type == TokenGoCode {
			codeTok := p.advance()
			if p.current.Type == TokenExprClose {
				p.advance()
			}
			return Attribute{
				Name:    attrName,
				Value:   strings.TrimSpace(codeTok.Value),
				Dynamic: true,
				Line:    line,
				Col:     col,
			}, nil
		}
		return Attribute{}, p.errorf("expected Go expression in dynamic attribute")
	default:
		return Attribute{}, p.errorf("expected attribute value for %q", attrName)
	}
}

// --- Expression parsers ---

// parseExprBlock parses {{ ... }}. By the time this is called, we know it's NOT
// a stop keyword (the body parser already checked via isAtKeyword).
func (p *Parser) parseExprBlock() (Node, error) {
	openTok := p.advance() // consume ExprOpen
	line, col := openTok.Line, openTok.Col

	if p.current.Type != TokenGoCode {
		if p.current.Type == TokenExprClose {
			p.advance()
			return &TextNode{Content: "", Line: line, Col: col}, nil
		}
		return nil, &ParseError{Message: "expected expression inside {{ }}", Line: line, Col: col}
	}

	code := strings.TrimSpace(p.current.Value)

	switch {
	case code == "children":
		p.advance()
		if p.current.Type == TokenExprClose {
			p.advance()
		}
		return &ChildrenNode{Line: line, Col: col}, nil

	case strings.HasPrefix(code, "if "):
		return p.parseIfNode(line, col)

	case strings.HasPrefix(code, "for "):
		return p.parseForNode(line, col)

	case strings.HasPrefix(code, "switch"):
		return p.parseSwitchNode(line, col)

	default:
		exprTok := p.advance()
		if p.current.Type == TokenExprClose {
			p.advance()
		}
		return &ExprNode{Expr: strings.TrimSpace(exprTok.Value), Line: line, Col: col}, nil
	}
}

func (p *Parser) parseRawExprBlock() (Node, error) {
	openTok := p.advance() // consume RawExprOpen
	line, col := openTok.Line, openTok.Col

	if p.current.Type != TokenGoCode {
		return nil, &ParseError{Message: "expected expression inside {{! }}", Line: line, Col: col}
	}

	exprTok := p.advance()
	if p.current.Type == TokenExprClose {
		p.advance()
	}
	return &RawExprNode{Expr: strings.TrimSpace(exprTok.Value), Line: line, Col: col}, nil
}

// --- Control flow parsers ---

func (p *Parser) parseIfNode(line, col int) (*IfNode, error) {
	code := strings.TrimSpace(p.current.Value)
	condition := strings.TrimPrefix(code, "if ")
	p.advance() // consume GoCode
	if p.current.Type == TokenExprClose {
		p.advance()
	}

	node := &IfNode{
		Condition: strings.TrimSpace(condition),
		Line:      line,
		Col:       col,
	}

	then, err := p.parseBody()
	if err != nil {
		return nil, err
	}
	node.Then = then

	// Check for {{ else }}
	if p.isAtKeyword("else") {
		p.advance() // ExprOpen
		p.advance() // GoCode("else")
		if p.current.Type == TokenExprClose {
			p.advance()
		}

		elseBody, err := p.parseBody()
		if err != nil {
			return nil, err
		}
		node.Else = elseBody
	}

	// Consume {{ end }}
	p.consumeEnd()

	return node, nil
}

func (p *Parser) parseForNode(line, col int) (*ForNode, error) {
	code := strings.TrimSpace(p.current.Value)
	clause := strings.TrimPrefix(code, "for ")
	p.advance() // consume GoCode
	if p.current.Type == TokenExprClose {
		p.advance()
	}

	node := &ForNode{
		Clause: strings.TrimSpace(clause),
		Line:   line,
		Col:    col,
	}

	body, err := p.parseBody()
	if err != nil {
		return nil, err
	}
	node.Body = body

	p.consumeEnd()

	return node, nil
}

func (p *Parser) parseSwitchNode(line, col int) (*SwitchNode, error) {
	code := strings.TrimSpace(p.current.Value)
	expr := strings.TrimPrefix(code, "switch")
	expr = strings.TrimSpace(expr)
	p.advance() // consume GoCode
	if p.current.Type == TokenExprClose {
		p.advance()
	}

	node := &SwitchNode{Expr: expr, Line: line, Col: col}

	// Parse cases until {{ end }}
	for {
		// Skip whitespace text
		if p.current.Type == TokenHTMLText && strings.TrimSpace(p.current.Value) == "" {
			p.advance()
			continue
		}

		if p.isAtKeyword("end") {
			p.consumeEnd()
			break
		}

		if p.isAtKeyword("case") {
			clause, err := p.parseCaseClause()
			if err != nil {
				return nil, err
			}
			node.Cases = append(node.Cases, clause)
			continue
		}

		if p.isAtKeyword("default") {
			clause, err := p.parseDefaultClause()
			if err != nil {
				return nil, err
			}
			node.Cases = append(node.Cases, clause)
			continue
		}

		if p.current.Type == TokenEOF || p.current.Type == TokenBraceClose {
			break
		}

		p.advance() // skip unexpected content
	}

	return node, nil
}

func (p *Parser) parseCaseClause() (*CaseClause, error) {
	p.advance() // consume ExprOpen
	codeTok := p.advance() // consume GoCode "case ..."
	line, col := codeTok.Line, codeTok.Col

	value := strings.TrimPrefix(strings.TrimSpace(codeTok.Value), "case ")
	value = strings.TrimSpace(value)

	if p.current.Type == TokenExprClose {
		p.advance()
	}

	body, err := p.parseBody()
	if err != nil {
		return nil, err
	}

	return &CaseClause{Value: value, Body: body, Line: line, Col: col}, nil
}

func (p *Parser) parseDefaultClause() (*CaseClause, error) {
	p.advance() // consume ExprOpen
	codeTok := p.advance() // consume GoCode "default"
	line, col := codeTok.Line, codeTok.Col

	if p.current.Type == TokenExprClose {
		p.advance()
	}

	body, err := p.parseBody()
	if err != nil {
		return nil, err
	}

	return &CaseClause{Default: true, Body: body, Line: line, Col: col}, nil
}
