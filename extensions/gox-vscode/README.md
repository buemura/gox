# Gox for Visual Studio Code

Syntax highlighting support for [Gox](https://github.com/user/gox) template files (`.gox`).

## Features

- Syntax highlighting for `.gox` template files
- Go code highlighting in expressions via embedded grammar delegation
- Distinct coloring for component calls (uppercase tags) vs HTML elements (lowercase tags)
- Control flow highlighting (`if`, `else`, `for`, `switch`, `case`, `end`)
- Attribute support: static, dynamic (`{{ expr }}`), boolean, and spread (`{{ attrs... }}`)
- Comment support: Go (`//`, `/* */`) and HTML (`<!-- -->`)
- Bracket matching and auto-closing pairs
- Code folding for components, HTML elements, and control flow blocks

## Supported Syntax

```gox
package views

import "fmt"

func Greeting(name string) {
  <div class="greeting">
    <h1>Hello, {{ name }}!</h1>
    {{ if name != "" }}
      <p>Welcome back!</p>
    {{ else }}
      <p>Welcome, stranger!</p>
    {{ end }}
  </div>
}
```

## Installation

### From Source

1. Copy the `gox-vscode` folder to your VSCode extensions directory:
   - **macOS/Linux**: `~/.vscode/extensions/`
   - **Windows**: `%USERPROFILE%\.vscode\extensions\`
2. Restart VSCode
3. Open any `.gox` file

## Requirements

- Visual Studio Code 1.75.0 or later
- The Go extension for VSCode (for embedded Go highlighting)
