# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What is Gox

Gox is a template engine for Go (similar to [templ](https://templ.guide/)). It compiles `.gox` template files into pure Go code. The `.gox` syntax combines Go package/import declarations, a `func` keyword for component definitions, HTML markup, and `{{ }}` expression blocks for Go expressions and control flow.

## Build & Run

```bash
go build -o gox ./cmd/gox          # build the CLI binary
go run ./cmd/gox generate           # compile all .gox files recursively
go run ./cmd/gox watch              # watch and recompile on changes
go run ./cmd/gox fmt                # format .gox files
go run ./cmd/gox lsp                # start LSP server (stdio)
```

Makefile targets: `make build`, `make test`, `make vet`, `make fmt`, `make check` (runs vet + test).

## Testing

```bash
go test ./...                       # run all tests
go test ./pkg/parser/               # run parser tests only
go test ./pkg/compiler/             # run compiler tests only
go test -run TestName ./pkg/parser/ # run a single test
```

## Architecture

The processing pipeline is: `.gox file → Lexer → Tokens → Parser → AST → Compiler → .go file`

Generated files use the naming convention `*_gox.go` (e.g., `views.gox` → `views_gox.go`).

### Key packages

- **`pkg/parser/`** — Lexer (`lexer.go`) and Parser (`parser.go`) that produce an AST (`ast.go`). The lexer is state-machine-based with states: `stateTopLevel`, `stateTemplate`, `stateTag`, `stateExpression`, `stateAttrValue`. Token types are defined in `token.go`.
- **`pkg/compiler/`** — Walks the parser AST and emits Go source code that writes HTML to an `io.Writer`. Output is run through `go/format` before writing. The compiler auto-adds required imports (`io`, `fmt`, `html`, `gox`) and conditionally adds `sort` when spread attributes are used.
- **`pkg/formatter/`** — Formats `.gox` source files.
- **`pkg/watcher/`** — File watcher using `fsnotify` for the `gox watch` command with configurable debounce.
- **`pkg/lsp/`** — Language Server Protocol implementation (hover, go-to-definition, diagnostics). Communicates over stdio via `jsonrpc2`.
- **`pkg/proxy/`** — HTTP reverse proxy with auto-reload for development; used by the `watch` command to proxy to a backend app and inject live-reload.
- **`cmd/gox/`** — CLI entry point dispatching `generate`, `watch`, `fmt`, and `lsp` subcommands.
- **`gox.go`** (root) — Runtime library providing `Component` interface, `ComponentFunc` adapter, `Attrs`, `Raw`, `SafeURL`, `SafeCSS`, sanitization helpers, and render functions. This is the package users import in their Go applications.

### Component compilation pattern

Each `func` component compiles to a Go function returning `gox.Component`. The generated code uses `gox.ComponentFunc` to wrap a closure that writes HTML via `io.WriteString`. Components that use `{{ children }}` get an additional `children gox.Component` parameter auto-appended by the compiler.

### Component vs HTML element distinction

Tags starting with an uppercase letter are parsed as `ComponentCall` nodes (compiled to function calls); lowercase tags become `HTMLElement` nodes (compiled to literal HTML output).

### Security: auto-sanitization

The compiler automatically applies `gox.SanitizeURL` for dynamic values in URL-sensitive attributes (`href`, `src`, `action`, `formaction`) and `gox.SanitizeCSS` for dynamic `style` attribute values. `SafeURL` and `SafeCSS` types bypass sanitization when the source is trusted.
