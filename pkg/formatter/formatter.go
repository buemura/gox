package formatter

import (
	"fmt"
	"strings"

	"github.com/buemura/gox/pkg/parser"
)

// Format takes a parsed AST File and returns formatted .gox source.
func Format(file *parser.File) string {
	f := &formatter{indent: 0}
	f.emitFile(file)
	return f.buf.String()
}

// FormatFile reads a .gox file path, parses it, and returns formatted source.
func FormatFile(path string, src string) (string, error) {
	p := parser.NewParser(src)
	ast, err := p.Parse()
	if err != nil {
		return "", fmt.Errorf("%s: %w", path, err)
	}
	return Format(ast), nil
}

type formatter struct {
	buf    strings.Builder
	indent int
}

func (f *formatter) write(s string) {
	f.buf.WriteString(s)
}

func (f *formatter) writeln(s string) {
	f.buf.WriteString(s)
	f.buf.WriteByte('\n')
}

func (f *formatter) writeIndent() {
	for range f.indent {
		f.buf.WriteString("  ")
	}
}

func (f *formatter) emitFile(file *parser.File) {
	if file.Package != nil {
		f.writeln("package " + file.Package.Name)
	}

	if len(file.Imports) > 0 {
		f.write("\n")
		for _, imp := range file.Imports {
			f.writeln("import " + imp.Path)
		}
	}

	for _, comp := range file.Components {
		f.write("\n")
		f.emitComponent(comp)
	}
}

func (f *formatter) emitComponent(c *parser.ComponentDecl) {
	if c.Params != "" {
		f.writeln("func " + c.Name + "(" + c.Params + ") {")
	} else {
		f.writeln("func " + c.Name + "() {")
	}
	f.indent++
	f.emitNodes(c.Body)
	f.indent--
	f.writeln("}")
}

func (f *formatter) emitNodes(nodes []parser.Node) {
	for _, node := range nodes {
		f.emitNode(node)
	}
}

func (f *formatter) emitNode(node parser.Node) {
	switch n := node.(type) {
	case *parser.TextNode:
		f.emitText(n)
	case *parser.ExprNode:
		f.writeIndent()
		f.writeln("{{ " + strings.TrimSpace(n.Expr) + " }}")
	case *parser.RawExprNode:
		f.writeIndent()
		f.writeln("{{! " + strings.TrimSpace(n.Expr) + " }}")
	case *parser.HTMLElement:
		f.emitHTMLElement(n)
	case *parser.ComponentCall:
		f.emitComponentCall(n)
	case *parser.IfNode:
		f.emitIf(n)
	case *parser.ForNode:
		f.emitFor(n)
	case *parser.SwitchNode:
		f.emitSwitch(n)
	case *parser.ChildrenNode:
		f.writeIndent()
		f.writeln("{{ children }}")
	}
}

func (f *formatter) emitText(t *parser.TextNode) {
	content := strings.TrimSpace(t.Content)
	if content == "" {
		return
	}
	f.writeIndent()
	f.writeln(content)
}

// isInline returns true if all children are text or expression nodes (no block elements).
func isInline(children []parser.Node) bool {
	if len(children) == 0 {
		return true
	}
	for _, child := range children {
		switch child.(type) {
		case *parser.TextNode, *parser.ExprNode, *parser.RawExprNode:
			continue
		default:
			return false
		}
	}
	return true
}

// hasContent returns true if there's any non-whitespace text in the children.
func hasContent(children []parser.Node) bool {
	for _, child := range children {
		switch n := child.(type) {
		case *parser.TextNode:
			if strings.TrimSpace(n.Content) != "" {
				return true
			}
		case *parser.ExprNode, *parser.RawExprNode, *parser.ChildrenNode:
			return true
		default:
			return true
		}
	}
	return false
}

func (f *formatter) emitAttrs(attrs []parser.Attribute) {
	for _, attr := range attrs {
		if attr.Spread {
			f.write(" {{ " + attr.Value + "... }}")
		} else if attr.Boolean {
			f.write(" " + attr.Name)
		} else if attr.Dynamic {
			f.write(" " + attr.Name + "={{ " + strings.TrimSpace(attr.Value) + " }}")
		} else {
			f.write(" " + attr.Name + "=\"" + attr.Value + "\"")
		}
	}
}

func (f *formatter) emitHTMLElement(el *parser.HTMLElement) {
	f.writeIndent()
	f.write("<" + el.Tag)
	f.emitAttrs(el.Attributes)

	if el.SelfClose || !hasContent(el.Children) {
		f.writeln(" />")
		return
	}

	if isInline(el.Children) {
		f.write(">")
		f.emitInlineChildren(el.Children)
		f.writeln("</" + el.Tag + ">")
		return
	}

	f.writeln(">")
	f.indent++
	f.emitNodes(el.Children)
	f.indent--
	f.writeIndent()
	f.writeln("</" + el.Tag + ">")
}

func (f *formatter) emitComponentCall(c *parser.ComponentCall) {
	f.writeIndent()
	f.write("<" + c.Name)
	f.emitAttrs(c.Attributes)

	if c.SelfClose || !hasContent(c.Children) {
		f.writeln(" />")
		return
	}

	if isInline(c.Children) {
		f.write(">")
		f.emitInlineChildren(c.Children)
		f.writeln("</" + c.Name + ">")
		return
	}

	f.writeln(">")
	f.indent++
	f.emitNodes(c.Children)
	f.indent--
	f.writeIndent()
	f.writeln("</" + c.Name + ">")
}

func (f *formatter) emitInlineChildren(children []parser.Node) {
	for i, child := range children {
		switch n := child.(type) {
		case *parser.TextNode:
			content := collapseWhitespace(n.Content)
			if i == 0 {
				content = strings.TrimLeft(content, " \t\n\r")
			}
			if i == len(children)-1 {
				content = strings.TrimRight(content, " \t\n\r")
			}
			if content != "" {
				f.write(content)
			}
		case *parser.ExprNode:
			f.write("{{ " + strings.TrimSpace(n.Expr) + " }}")
		case *parser.RawExprNode:
			f.write("{{! " + strings.TrimSpace(n.Expr) + " }}")
		}
	}
}

// collapseWhitespace replaces runs of whitespace with a single space.
func collapseWhitespace(s string) string {
	var buf strings.Builder
	inSpace := false
	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
			if !inSpace {
				buf.WriteByte(' ')
				inSpace = true
			}
		} else {
			buf.WriteRune(r)
			inSpace = false
		}
	}
	return buf.String()
}

func (f *formatter) emitIf(n *parser.IfNode) {
	f.writeIndent()
	f.writeln("{{ if " + strings.TrimSpace(n.Condition) + " }}")
	f.indent++
	f.emitNodes(n.Then)
	f.indent--
	if len(n.Else) > 0 {
		f.writeIndent()
		f.writeln("{{ else }}")
		f.indent++
		f.emitNodes(n.Else)
		f.indent--
	}
	f.writeIndent()
	f.writeln("{{ end }}")
}

func (f *formatter) emitFor(n *parser.ForNode) {
	f.writeIndent()
	f.writeln("{{ for " + strings.TrimSpace(n.Clause) + " }}")
	f.indent++
	f.emitNodes(n.Body)
	f.indent--
	f.writeIndent()
	f.writeln("{{ end }}")
}

func (f *formatter) emitSwitch(n *parser.SwitchNode) {
	f.writeIndent()
	f.writeln("{{ switch " + strings.TrimSpace(n.Expr) + " }}")
	f.indent++
	for _, c := range n.Cases {
		if c.Default {
			f.writeIndent()
			f.writeln("{{ default }}")
		} else {
			f.writeIndent()
			f.writeln("{{ case " + strings.TrimSpace(c.Value) + " }}")
		}
		f.indent++
		f.emitNodes(c.Body)
		f.indent--
	}
	f.indent--
	f.writeIndent()
	f.writeln("{{ end }}")
}
