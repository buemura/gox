# Gox

[![Go Reference](https://pkg.go.dev/badge/github.com/buemura/gox.svg)](https://pkg.go.dev/github.com/buemura/gox)
[![CI](https://github.com/buemura/gox/actions/workflows/ci.yml/badge.svg)](https://github.com/buemura/gox/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/buemura/gox)](https://goreportcard.com/report/github.com/buemura/gox)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Gox is a template engine for Go that compiles `.gox` template files into pure Go code. Similar to [templ](https://templ.guide/), it provides type-safe, composable components with the full power of Go.

The `.gox` syntax combines Go package/import declarations, component definitions via `func`, HTML markup, and `{{ }}` expression blocks for Go expressions and control flow.

## Features

- **Type-safe templates** -- Components are compiled to Go functions with full type checking
- **Composable components** -- Build UIs with reusable, nestable components and children slots
- **Auto-sanitization** -- XSS protection with automatic HTML escaping and URL/CSS sanitization
- **Hot reload** -- Watch mode with reverse proxy and automatic browser reload
- **LSP support** -- Editor integration with diagnostics, hover, and go-to-definition
- **Zero runtime dependencies** -- Generated code uses only the Go standard library and the lightweight `gox` runtime

## Installation

```bash
go install github.com/buemura/gox/cmd/gox@latest
```

## Quick Start

Create a `views/views.gox` file:

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

This compiles `views/views.gox` into `views/views_gox.go`. Use the generated component in your application:

```go
package main

import (
    "net/http"
    "yourapp/views"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        views.Hello("World").Render(r.Context(), w)
    })
    http.ListenAndServe(":8080", nil)
}
```

## CLI Commands

### `gox generate`

Compile all `.gox` files recursively into Go source code.

```bash
gox generate              # compile from current directory
gox generate -dir ./views # compile from specific directory
gox generate -out _gox.go # custom output suffix (default: _gox.go)
```

Each `.gox` file produces a corresponding `*_gox.go` file (e.g., `views.gox` becomes `views_gox.go`).

### `gox watch`

Watch for `.gox` file changes and recompile automatically. Supports running a command on recompilation and proxying with hot-reload.

```bash
gox watch                                          # watch and recompile
gox watch -dir ./views                             # watch specific directory
gox watch -cmd "go run ."                          # restart app on changes
gox watch -proxy http://localhost:8080 -proxyport 3000  # proxy with hot-reload
```

| Flag         | Default   | Description                                      |
| ------------ | --------- | ------------------------------------------------ |
| `-dir`       | `.`       | Root directory to watch                          |
| `-out`       | `_gox.go` | Output file suffix                               |
| `-cmd`       |           | Command to run/restart after recompilation       |
| `-proxy`     |           | Upstream URL to proxy (enables hot-reload proxy) |
| `-proxyport` | `8081`    | Port for the hot-reload proxy server             |

When `-cmd` is set, the process is killed and restarted on each recompilation. When `-proxy` is set, a reverse proxy with auto-reload is started on the specified port.

### `gox fmt`

Format `.gox` files.

```bash
gox fmt                # format in place
gox fmt -dir ./views   # format specific directory
gox fmt -stdout        # print to stdout instead of writing
```

### `gox lsp`

Start the Language Server Protocol server over stdio. Used by editor extensions for diagnostics, hover, and go-to-definition support.

```bash
gox lsp
```

## Template Syntax

### Components

Define components with `func`. Each component compiles to a Go function returning `gox.Component`:

```gox
package views

func Button(label string, disabled bool) {
  <button class="btn" disabled?={{ disabled }}>
    {{ label }}
  </button>
}
```

### Expressions

Use `{{ }}` blocks to embed Go expressions. All expressions are HTML-escaped by default:

```gox
<span>{{ user.Name }}</span>
<div class={{ activeClass }}>content</div>
<a href={{ "/user/" + user.Username }}>Profile</a>
```

### Raw HTML

Use `{{! }}` to output unescaped HTML. This is useful for inline scripts, styles, or pre-sanitized content:

```gox
{{! `<script>console.log("hello")</script>` }}
{{! `<style>body { margin: 0; }</style>` }}
```

You can also use the `gox.Raw` type for unescaped content within expressions:

```gox
<div>{{ gox.Raw(trustedHTML) }}</div>
```

### Attributes

**Static attributes** work like regular HTML:

```gox
<input type="text" class="input" placeholder="Enter name" />
```

**Dynamic attributes** use `{{ }}` for Go expression values:

```gox
<div class={{ computeClass() }}>content</div>
<button onclick={{ fmt.Sprintf("doAction(%d)", id) }}>Click</button>
```

**Boolean attributes** use `?=` to conditionally include the attribute:

```gox
<input type="checkbox" checked?={{ isChecked }} />
<button disabled?={{ isDisabled }}>Submit</button>
```

**Spread attributes** use `gox.Attrs` to pass dynamic attribute maps:

```gox
<div {{ attrs }}>content</div>
```

```go
attrs := gox.Attrs{"class": "card", "id": "main"}
```

### Control Flow

**If/else:**

```gox
{{ if len(items) == 0 }}
  <p>No items yet.</p>
{{ else }}
  <p>{{ len(items) }} items found.</p>
{{ end }}
```

**For loops:**

```gox
{{ for _, item := range items }}
  <li>{{ item.Name }}</li>
{{ end }}
```

**Switch/case:**

```gox
{{ switch status }}
  {{ case "active" }}
    <span class="badge-green">Active</span>
  {{ case "inactive" }}
    <span class="badge-gray">Inactive</span>
  {{ default }}
    <span class="badge-red">Unknown</span>
{{ end }}
```

### Children (Slots)

Components can accept children using `{{ children }}`. The compiler automatically adds a `children gox.Component` parameter:

```gox
func Card(title string) {
  <div class="card">
    <h2>{{ title }}</h2>
    <div class="card-body">
      {{ children }}
    </div>
  </div>
}
```

Use it by nesting content inside the component tag:

```gox
<Card title="Welcome">
  <p>This is the card body content.</p>
  <Button label="Click me" />
</Card>
```

### Component Calls vs HTML Elements

- **Uppercase** tags (e.g., `<Button />`, `<Card>`) are component calls -- compiled to Go function calls
- **Lowercase** tags (e.g., `<button>`, `<div>`) are HTML elements -- compiled to literal HTML output

Component attributes map to function parameters by position:

```gox
// Definition
func Avatar(src string, size string) { ... }

// Usage
<Avatar src={{ user.AvatarURL }} size="lg" />
```

## Runtime API

Import the `gox` package in your Go application:

```go
import "github.com/buemura/gox"
```

### Rendering

```go
// Render to an io.Writer (e.g., http.ResponseWriter)
err := views.MyComponent(args).Render(r.Context(), w)

// Or use the convenience function
err := gox.Render(ctx, w, views.MyComponent(args))

// Render to a string
html, err := gox.RenderToString(ctx, views.MyComponent(args))
```

### Types

| Type                | Description                                                     |
| ------------------- | --------------------------------------------------------------- |
| `gox.Component`     | Interface with `Render(ctx context.Context, w io.Writer) error` |
| `gox.ComponentFunc` | Adapter to use a function as a `Component`                      |
| `gox.Attrs`         | `map[string]any` for spread attributes                          |
| `gox.Raw`           | String type that bypasses HTML escaping                         |
| `gox.SafeURL`       | String type that bypasses URL sanitization                      |
| `gox.SafeCSS`       | String type that bypasses CSS sanitization                      |

### Auto-Sanitization

Gox automatically sanitizes dynamic values in security-sensitive attributes:

- **URL attributes** (`href`, `src`, `action`, `formaction`): Dynamic values are passed through `gox.SanitizeURL`, which blocks `javascript:`, `vbscript:`, and `data:` schemes. Use `gox.SafeURL` to bypass when the source is trusted.
- **Style attributes**: Dynamic `style` values are passed through `gox.SanitizeCSS`, which blocks dangerous patterns like `expression()`, `url()`, and `behavior:`. Use `gox.SafeCSS` to bypass when the source is trusted.
- **All other expressions**: HTML-escaped via `html.EscapeString` by default.

## Development Workflow

A typical development setup with `watch`, command restart, and hot-reload proxy:

```bash
gox watch -dir ./views -cmd "go run ." -proxy http://localhost:8080 -proxyport 3000
```

This will:

1. Compile all `.gox` files on startup
2. Watch for changes and recompile
3. Restart `go run .` after each recompilation
4. Proxy `localhost:8080` on port `3000` with auto-reload

Example Makefile:

```makefile
GOX = go run github.com/buemura/gox/cmd/gox

generate:
	$(GOX) generate -dir .

build: generate
	go build -o myapp .

dev: generate
	$(GOX) watch -dir . -cmd "go run ." &
	wait

clean:
	rm -f myapp views/*_gox.go
```

## Examples

See the [examples/](examples/) directory for complete applications:

- **[todo](examples/todo/)** -- A todo app with Tailwind CSS
- **[social_media](examples/social_media/)** -- A social media app with authentication, feed, and profiles

## Architecture

The processing pipeline:

```
.gox file → Lexer → Tokens → Parser → AST → Compiler → .go file
```

| Package          | Description                                                           |
| ---------------- | --------------------------------------------------------------------- |
| `cmd/gox/`       | CLI entry point (`generate`, `watch`, `fmt`, `lsp`)                   |
| `pkg/parser/`    | Lexer and parser that produce an AST                                  |
| `pkg/compiler/`  | Walks the AST and emits Go source code                                |
| `pkg/formatter/` | Formats `.gox` source files                                           |
| `pkg/watcher/`   | File watcher for the `watch` command                                  |
| `pkg/lsp/`       | Language Server Protocol server                                       |
| `pkg/proxy/`     | HTTP reverse proxy with hot-reload                                    |
| `gox.go`         | Runtime library (`Component`, `Render`, `Attrs`, `Raw`, sanitization) |

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License -- see the [LICENSE](LICENSE) file for details.
