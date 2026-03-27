# Gox

Gox is a template engine for Go that compiles `.gox` template files into pure Go code. It is similar to [templ](https://templ.guide/).

The `.gox` syntax combines Go package/import declarations, component definitions via `func`, HTML markup, and `{{ }}` expression blocks for Go expressions and control flow.

## Installation

```bash
go install github.com/buemura/gox/cmd/gox@latest
```

## Quick Start

Create a `.gox` file:

```gox
package views

func Hello(name string) {
  <div>
    <h1>Hello, {{ name }}!</h1>
  </div>
}
```

Generate the Go code:

```bash
gox generate
```

This compiles all `.gox` files recursively into `.go` files that you can use in your application.

## Usage

### CLI Commands

```bash
gox generate    # compile all .gox files recursively
gox watch       # watch and recompile on changes
gox fmt         # format .gox files
```

### Runtime API

Import the `gox` package in your Go application to render components:

```go
import "github.com/buemura/gox"

// Render to an io.Writer (e.g. http.ResponseWriter)
gox.Render(w, views.Hello("World"))

// Render to a string
html, err := gox.RenderToString(views.Hello("World"))
```

## Template Syntax

### Components

Define components with `func`. Each component compiles to a Go function returning `gox.Component`:

```gox
package views

func Button(label string) {
  <button class="btn">{{ label }}</button>
}
```

### Expressions

Use `{{ }}` blocks to embed Go expressions:

```gox
<span>{{ user.Name }}</span>
<div class={{ activeClass }}>content</div>
```

### Control Flow

Standard Go control flow works inside `{{ }}` blocks:

```gox
{{ if loggedIn }}
  <p>Welcome back!</p>
{{ else }}
  <p>Please log in.</p>
{{ end }}

{{ for _, item := range items }}
  <li>{{ item.Name }}</li>
{{ end }}
```

### Children (Slots)

Components can accept children using `{{ children }}`:

```gox
func Layout(title string) {
  <html>
    <head><title>{{ title }}</title></head>
    <body>{{ children }}</body>
  </html>
}
```

Use it by nesting content inside the component tag:

```gox
<Layout title="Home">
  <h1>Welcome</h1>
</Layout>
```

### Component Calls vs HTML Elements

- **Uppercase** tags (e.g. `<Button />`) are component calls — compiled to Go function calls.
- **Lowercase** tags (e.g. `<button>`) are HTML elements — compiled to literal HTML output.

### Spread Attributes

Use `gox.Attrs` to pass dynamic attributes:

```gox
<div {{ attrs }}>content</div>
```

### Raw HTML

Use `gox.Raw` to output unescaped HTML:

```gox
<div>{{ gox.Raw(htmlString) }}</div>
```

## Architecture

The processing pipeline:

```
.gox file → Lexer → Tokens → Parser → AST → Compiler → .go file
```

| Package | Description |
|---|---|
| `cmd/gox/` | CLI entry point (`generate`, `watch`, `fmt`) |
| `pkg/parser/` | Lexer and parser that produce an AST |
| `pkg/compiler/` | Walks the AST and emits Go source code |
| `pkg/formatter/` | Formats `.gox` source files |
| `pkg/watcher/` | File watcher for the `watch` command |
| `gox.go` | Runtime library (`Component`, `Render`, `Attrs`, `Raw`) |

## Examples

See the [examples/](examples/) directory for complete applications:

- **[todo](examples/todo/)** — A todo app with Tailwind CSS
- **[social_media](examples/social_media/)** — A social media app with authentication, feed, and profiles

## License

MIT
