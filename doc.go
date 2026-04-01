// Package gox is a template engine for Go that compiles .gox template files
// into type-safe, pure Go code. It provides a component-based model similar
// to templ, where each component is a Go function returning a [Component].
//
// # Overview
//
// Gox compiles .gox files through a pipeline:
//
//	.gox file -> Lexer -> Tokens -> Parser -> AST -> Compiler -> .go file
//
// The .gox syntax combines Go package/import declarations, component
// definitions via func, standard HTML markup, and {{ }} expression blocks
// for embedding Go expressions and control flow.
//
// Install the CLI:
//
//	go install github.com/buemura/gox/cmd/gox@latest
//
// # Defining Components
//
// Create a .gox file with component definitions:
//
//	package views
//
//	func Hello(name string) {
//	  <div>
//	    <h1>Hello, {{ name }}!</h1>
//	  </div>
//	}
//
// Run gox generate to compile it into Go code, then use it in your application:
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//	    views.Hello("World").Render(r.Context(), w)
//	}
//
// # Rendering
//
// Components implement the [Component] interface with a Render method that
// accepts a [context.Context] and an [io.Writer]:
//
//	// Render directly
//	err := component.Render(ctx, w)
//
//	// Using the convenience function
//	err := gox.Render(ctx, w, component)
//
//	// Render to a string
//	html, err := gox.RenderToString(ctx, component)
//
// # Template Syntax
//
// Expressions use {{ }} blocks and are HTML-escaped by default:
//
//	<span>{{ user.Name }}</span>
//	<div class={{ activeClass }}>content</div>
//
// Raw (unescaped) HTML uses {{! }}:
//
//	{{! `<script>console.log("hello")</script>` }}
//
// Control flow uses standard Go syntax inside {{ }} blocks:
//
//	{{ if len(items) > 0 }}
//	  {{ for _, item := range items }}
//	    <li>{{ item.Name }}</li>
//	  {{ end }}
//	{{ end }}
//
// Boolean attributes use the ?= syntax:
//
//	<button disabled?={{ isDisabled }}>Submit</button>
//
// Spread attributes use [Attrs]:
//
//	<div {{ attrs }}>content</div>
//
// Components can accept children using {{ children }}, which auto-adds
// a children [Component] parameter:
//
//	func Card(title string) {
//	  <div class="card">
//	    <h2>{{ title }}</h2>
//	    {{ children }}
//	  </div>
//	}
//
// Uppercase tags are component calls (compiled to function calls), while
// lowercase tags are HTML elements (compiled to literal HTML output):
//
//	<Card title="Welcome">       // component call
//	  <p>Hello</p>               // HTML element
//	</Card>
//
// # Security
//
// The compiler automatically sanitizes dynamic values in security-sensitive
// attributes:
//
//   - URL attributes (href, src, action, formaction) are sanitized via
//     [SanitizeURL], blocking javascript:, vbscript:, and data: schemes.
//     Use [SafeURL] to bypass when the source is trusted.
//   - Style attributes are sanitized via [SanitizeCSS], blocking dangerous
//     patterns like expression(), url(), and behavior:.
//     Use [SafeCSS] to bypass when the source is trusted.
//   - All other dynamic expressions are HTML-escaped by default.
//
// # CLI Commands
//
// The gox CLI provides the following commands:
//
//	gox generate    Compile .gox files into Go source code
//	gox watch       Watch and recompile on changes, with optional hot-reload proxy
//	gox fmt         Format .gox files
//	gox lsp         Start the Language Server Protocol server
package gox
