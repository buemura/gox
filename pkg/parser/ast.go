package parser

// Node is the interface implemented by all AST nodes.
type Node interface {
	nodeType() string
	Pos() (line, col int)
}

// File represents an entire .gox source file.
type File struct {
	Package    *PackageDecl
	Imports    []*ImportDecl
	Components []*ComponentDecl
}

func (f *File) nodeType() string    { return "File" }
func (f *File) Pos() (int, int)     { return 1, 1 }

// PackageDecl represents a package declaration (e.g., "package views").
type PackageDecl struct {
	Name string
	Line int
	Col  int
}

func (p *PackageDecl) nodeType() string    { return "PackageDecl" }
func (p *PackageDecl) Pos() (int, int)     { return p.Line, p.Col }

// ImportDecl represents a single import declaration (e.g., `import "fmt"`).
type ImportDecl struct {
	Path string
	Line int
	Col  int
}

func (i *ImportDecl) nodeType() string    { return "ImportDecl" }
func (i *ImportDecl) Pos() (int, int)     { return i.Line, i.Col }

// ComponentDecl represents a gox component declaration.
type ComponentDecl struct {
	Name   string
	Params string // raw parameter list, e.g. "name string, age int"
	Body   []Node
	Line   int
	Col    int
}

func (c *ComponentDecl) nodeType() string    { return "ComponentDecl" }
func (c *ComponentDecl) Pos() (int, int)     { return c.Line, c.Col }

// HTMLElement represents an HTML element with tag, attributes, and children.
type HTMLElement struct {
	Tag        string
	Attributes []Attribute
	Children   []Node
	SelfClose  bool
	Line       int
	Col        int
}

func (h *HTMLElement) nodeType() string    { return "HTMLElement" }
func (h *HTMLElement) Pos() (int, int)     { return h.Line, h.Col }

// Attribute represents an HTML attribute on an element.
type Attribute struct {
	Name    string // attribute name (empty for spread)
	Value   string // static value (empty if dynamic)
	Dynamic bool   // true if value is a Go expression
	Spread  bool   // true if this is a spread attribute
	Boolean bool   // true if no value (e.g., "disabled")
	Line    int
	Col     int
}

// TextNode represents static HTML text content.
type TextNode struct {
	Content string
	Line    int
	Col     int
}

func (t *TextNode) nodeType() string    { return "TextNode" }
func (t *TextNode) Pos() (int, int)     { return t.Line, t.Col }

// ExprNode represents an escaped Go expression ({{ expr }}).
type ExprNode struct {
	Expr string
	Line int
	Col  int
}

func (e *ExprNode) nodeType() string    { return "ExprNode" }
func (e *ExprNode) Pos() (int, int)     { return e.Line, e.Col }

// RawExprNode represents an unescaped Go expression ({{! expr }}).
type RawExprNode struct {
	Expr string
	Line int
	Col  int
}

func (r *RawExprNode) nodeType() string    { return "RawExprNode" }
func (r *RawExprNode) Pos() (int, int)     { return r.Line, r.Col }

// IfNode represents an if/else control flow block.
type IfNode struct {
	Condition string
	Then      []Node
	Else      []Node // nil if no else branch
	Line      int
	Col       int
}

func (i *IfNode) nodeType() string    { return "IfNode" }
func (i *IfNode) Pos() (int, int)     { return i.Line, i.Col }

// ForNode represents a for loop block.
type ForNode struct {
	Clause string // e.g. "_, item := range items"
	Body   []Node
	Line   int
	Col    int
}

func (f *ForNode) nodeType() string    { return "ForNode" }
func (f *ForNode) Pos() (int, int)     { return f.Line, f.Col }

// SwitchNode represents a switch/case/default block.
type SwitchNode struct {
	Expr  string
	Cases []*CaseClause
	Line  int
	Col   int
}

func (s *SwitchNode) nodeType() string    { return "SwitchNode" }
func (s *SwitchNode) Pos() (int, int)     { return s.Line, s.Col }

// CaseClause represents a single case or default clause in a switch.
type CaseClause struct {
	Value   string // empty for default
	Default bool
	Body    []Node
	Line    int
	Col     int
}

func (c *CaseClause) nodeType() string    { return "CaseClause" }
func (c *CaseClause) Pos() (int, int)     { return c.Line, c.Col }

// ComponentCall represents a call to another component (<Component /> or <Component>...</Component>).
type ComponentCall struct {
	Name       string
	Attributes []Attribute
	Children   []Node
	SelfClose  bool
	Line       int
	Col        int
}

func (c *ComponentCall) nodeType() string    { return "ComponentCall" }
func (c *ComponentCall) Pos() (int, int)     { return c.Line, c.Col }

// ChildrenNode represents a {{ children }} slot placeholder.
type ChildrenNode struct {
	Line int
	Col  int
}

func (c *ChildrenNode) nodeType() string    { return "ChildrenNode" }
func (c *ChildrenNode) Pos() (int, int)     { return c.Line, c.Col }
