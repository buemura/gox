package compiler

import (
	"fmt"
	"go/format"
	"strings"

	"github.com/buemura/gox/pkg/parser"
)

// Compiler walks an AST and emits Go source code.
type Compiler struct {
	buf    strings.Builder
	indent int
}

// Compile takes a parsed AST File and returns formatted Go source code.
func Compile(file *parser.File) (string, error) {
	c := &Compiler{}
	c.compileFile(file)

	src := c.buf.String()
	formatted, err := format.Source([]byte(src))
	if err != nil {
		return src, fmt.Errorf("format error: %w\n\nraw output:\n%s", err, src)
	}
	return string(formatted), nil
}

func (c *Compiler) write(s string) {
	c.buf.WriteString(s)
}

func (c *Compiler) writef(format string, args ...any) {
	fmt.Fprintf(&c.buf, format, args...)
}

func (c *Compiler) writeLine(s string) {
	c.writeIndent()
	c.write(s)
	c.write("\n")
}

func (c *Compiler) writeIndent() {
	for i := 0; i < c.indent; i++ {
		c.write("\t")
	}
}

func (c *Compiler) compileFile(file *parser.File) {
	// Package declaration
	c.writef("package %s\n\n", file.Package.Name)

	// Collect required imports
	requiredImports := map[string]bool{
		"io":                          true,
		"fmt":                         true,
		"html":                        true,
		"github.com/buemura/gox": true,
	}

	// gox is always needed since every component returns gox.Component
	// and uses gox.ComponentFunc
	needsSort := false
	for _, comp := range file.Components {
		if walkNeedsSort(comp.Body) {
			needsSort = true
		}
	}
	if needsSort {
		requiredImports["sort"] = true
	}

	// Merge user imports (strip quotes since parser preserves them from source)
	for _, imp := range file.Imports {
		path := strings.Trim(imp.Path, `"`)
		requiredImports[path] = true
	}

	// Emit imports
	c.writeLine("import (")
	c.indent++
	for path := range requiredImports {
		c.writef("\t%q\n", path)
	}
	c.indent--
	c.writeLine(")")
	c.write("\n")

	// Emit components
	for i, comp := range file.Components {
		c.compileComponent(comp)
		if i < len(file.Components)-1 {
			c.write("\n")
		}
	}
}


func walkNeedsSort(nodes []parser.Node) bool {
	for _, n := range nodes {
		switch v := n.(type) {
		case *parser.HTMLElement:
			for _, attr := range v.Attributes {
				if attr.Spread {
					return true
				}
			}
			if walkNeedsSort(v.Children) {
				return true
			}
		case *parser.IfNode:
			if walkNeedsSort(v.Then) || walkNeedsSort(v.Else) {
				return true
			}
		case *parser.ForNode:
			if walkNeedsSort(v.Body) {
				return true
			}
		case *parser.SwitchNode:
			for _, cc := range v.Cases {
				if walkNeedsSort(cc.Body) {
					return true
				}
			}
		}
	}
	return false
}

func (c *Compiler) compileComponent(comp *parser.ComponentDecl) {
	// Determine if component uses children (needs children parameter)
	hasChildren := walkHasChildren(comp.Body)

	params := comp.Params
	if hasChildren {
		if params != "" {
			params += ", "
		}
		params += "children gox.Component"
	}

	c.writef("func %s(%s) gox.Component {\n", comp.Name, params)
	c.indent++
	c.writeLine("return gox.ComponentFunc(func(w io.Writer) error {")
	c.indent++
	c.writeLine("var err error")
	c.writeLine("_ = err")

	c.compileNodes(comp.Body)

	c.writeLine("return nil")
	c.indent--
	c.writeLine("})")
	c.indent--
	c.writeLine("}")
}

func walkHasChildren(nodes []parser.Node) bool {
	for _, n := range nodes {
		switch v := n.(type) {
		case *parser.ChildrenNode:
			return true
		case *parser.HTMLElement:
			if walkHasChildren(v.Children) {
				return true
			}
		case *parser.IfNode:
			if walkHasChildren(v.Then) || walkHasChildren(v.Else) {
				return true
			}
		case *parser.ForNode:
			if walkHasChildren(v.Body) {
				return true
			}
		case *parser.SwitchNode:
			for _, cc := range v.Cases {
				if walkHasChildren(cc.Body) {
					return true
				}
			}
		}
	}
	return false
}

func (c *Compiler) compileNodes(nodes []parser.Node) {
	for _, node := range nodes {
		c.compileNode(node)
	}
}

func (c *Compiler) compileNode(node parser.Node) {
	switch n := node.(type) {
	case *parser.TextNode:
		c.compileTextNode(n)
	case *parser.ExprNode:
		c.compileExprNode(n)
	case *parser.RawExprNode:
		c.compileRawExprNode(n)
	case *parser.HTMLElement:
		c.compileHTMLElement(n)
	case *parser.IfNode:
		c.compileIfNode(n)
	case *parser.ForNode:
		c.compileForNode(n)
	case *parser.SwitchNode:
		c.compileSwitchNode(n)
	case *parser.ComponentCall:
		c.compileComponentCall(n)
	case *parser.ChildrenNode:
		c.compileChildrenNode()
	}
}

func (c *Compiler) emitWriteString(s string) {
	if s == "" {
		return
	}
	c.writef("\t_, err = io.WriteString(w, %q)\n", s)
	c.writeIndent()
	c.writeLine("if err != nil { return err }")
}

func (c *Compiler) emitWriteExpr(expr string) {
	c.writeIndent()
	c.writef("_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf(\"%%v\", %s)))\n", expr)
	c.writeIndent()
	c.writeLine("if err != nil { return err }")
}

func (c *Compiler) emitWriteRawExpr(expr string) {
	c.writeIndent()
	c.writef("_, err = io.WriteString(w, fmt.Sprintf(\"%%v\", %s))\n", expr)
	c.writeIndent()
	c.writeLine("if err != nil { return err }")
}

func (c *Compiler) compileTextNode(n *parser.TextNode) {
	if n.Content == "" {
		return
	}
	c.writeIndent()
	c.emitWriteString(n.Content)
}

func (c *Compiler) compileExprNode(n *parser.ExprNode) {
	c.emitWriteExpr(n.Expr)
}

func (c *Compiler) compileRawExprNode(n *parser.RawExprNode) {
	c.emitWriteRawExpr(n.Expr)
}

func (c *Compiler) compileHTMLElement(el *parser.HTMLElement) {
	// Open tag
	c.writeIndent()
	c.emitWriteString("<" + el.Tag)

	// Static and dynamic attributes
	for _, attr := range el.Attributes {
		if attr.Spread {
			c.compileSpreadAttr(attr)
		} else if attr.Boolean {
			c.compileBooleanAttr(attr)
		} else if attr.Dynamic {
			c.compileDynamicAttr(attr)
		} else {
			// Static attribute - emit inline
			c.writeIndent()
			c.emitWriteString(fmt.Sprintf(` %s="%s"`, attr.Name, attr.Value))
		}
	}

	if el.SelfClose {
		c.writeIndent()
		c.emitWriteString(" />")
		return
	}

	c.writeIndent()
	c.emitWriteString(">")

	// Children
	c.compileNodes(el.Children)

	// Close tag
	c.writeIndent()
	c.emitWriteString("</" + el.Tag + ">")
}

func (c *Compiler) compileDynamicAttr(attr parser.Attribute) {
	// Dynamic attribute: emit name="<escaped value>"
	c.writeIndent()
	c.emitWriteString(fmt.Sprintf(` %s="`, attr.Name))
	c.emitWriteExpr(attr.Value)
	c.writeIndent()
	c.emitWriteString(`"`)
}

func (c *Compiler) compileBooleanAttr(attr parser.Attribute) {
	if attr.Dynamic {
		// Conditional boolean: only emit if expression is truthy
		c.writeIndent()
		c.writef("if %s {\n", attr.Value)
		c.indent++
		c.writeIndent()
		c.emitWriteString(fmt.Sprintf(` %s`, attr.Name))
		c.indent--
		c.writeLine("}")
	} else {
		// Static boolean attribute
		c.writeIndent()
		c.emitWriteString(fmt.Sprintf(` %s`, attr.Name))
	}
}

func (c *Compiler) compileSpreadAttr(attr parser.Attribute) {
	// Spread attributes: iterate map keys in sorted order
	c.writeIndent()
	c.writef("{\n")
	c.indent++
	c.writef("\t__keys := make([]string, 0, len(%s))\n", attr.Value)
	c.writeIndent()
	c.writef("for __k := range %s {\n", attr.Value)
	c.indent++
	c.writeLine("__keys = append(__keys, __k)")
	c.indent--
	c.writeLine("}")
	c.writeLine("sort.Strings(__keys)")
	c.writeLine("for _, __k := range __keys {")
	c.indent++
	c.writef("\t__v := %s[__k]\n", attr.Value)
	c.writeIndent()
	c.writeLine(`if __vBool, __ok := __v.(bool); __ok {`)
	c.indent++
	c.writeLine("if __vBool {")
	c.indent++
	c.writeIndent()
	c.write(`_, err = io.WriteString(w, " " + __k)`)
	c.write("\n")
	c.writeIndent()
	c.writeLine("if err != nil { return err }")
	c.indent--
	c.writeLine("}")
	c.indent--
	c.writeLine("} else {")
	c.indent++
	c.writeIndent()
	c.write(`_, err = io.WriteString(w, fmt.Sprintf(" %s=\"%v\"", __k, __v))`)
	c.write("\n")
	c.writeIndent()
	c.writeLine("if err != nil { return err }")
	c.indent--
	c.writeLine("}")
	c.indent--
	c.writeLine("}")
	c.indent--
	c.writeLine("}")
}

func (c *Compiler) compileIfNode(n *parser.IfNode) {
	c.writeIndent()
	c.writef("if %s {\n", n.Condition)
	c.indent++
	c.compileNodes(n.Then)
	c.indent--
	if len(n.Else) > 0 {
		c.writeLine("} else {")
		c.indent++
		c.compileNodes(n.Else)
		c.indent--
	}
	c.writeLine("}")
}

func (c *Compiler) compileForNode(n *parser.ForNode) {
	c.writeIndent()
	c.writef("for %s {\n", n.Clause)
	c.indent++
	c.compileNodes(n.Body)
	c.indent--
	c.writeLine("}")
}

func (c *Compiler) compileSwitchNode(n *parser.SwitchNode) {
	c.writeIndent()
	c.writef("switch %s {\n", n.Expr)
	for _, cc := range n.Cases {
		if cc.Default {
			c.writeLine("default:")
		} else {
			c.writef("case %s:\n", cc.Value)
		}
		c.indent++
		c.compileNodes(cc.Body)
		c.indent--
	}
	c.writeLine("}")
}

func (c *Compiler) compileComponentCall(call *parser.ComponentCall) {
	// Build argument list from attributes
	var args []string
	for _, attr := range call.Attributes {
		if attr.Spread {
			continue // spreads handled differently for components if needed
		}
		if attr.Dynamic {
			args = append(args, attr.Value)
		} else {
			args = append(args, fmt.Sprintf("%q", attr.Value))
		}
	}

	// If component has children, render them as a Component argument
	if len(call.Children) > 0 {
		args = append(args, c.buildChildrenComponent(call.Children))
	}

	c.writeIndent()
	c.writef("err = %s(%s).Render(w)\n", call.Name, strings.Join(args, ", "))
	c.writeIndent()
	c.writeLine("if err != nil { return err }")
}

func (c *Compiler) buildChildrenComponent(nodes []parser.Node) string {
	// Build an inline gox.ComponentFunc for children
	inner := &Compiler{indent: 2}
	inner.compileNodes(nodes)

	var b strings.Builder
	b.WriteString("gox.ComponentFunc(func(w io.Writer) error {\n")
	b.WriteString("\t\tvar err error\n")
	b.WriteString("\t\t_ = err\n")
	b.WriteString(inner.buf.String())
	b.WriteString("\t\treturn nil\n")
	b.WriteString("\t})")
	return b.String()
}

func (c *Compiler) compileChildrenNode() {
	c.writeLine("if children != nil {")
	c.indent++
	c.writeLine("err = children.Render(w)")
	c.writeLine("if err != nil { return err }")
	c.indent--
	c.writeLine("}")
}
