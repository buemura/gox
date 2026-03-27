package gox

import (
	"bytes"
	"io"
)

// Component is the interface returned by compiled templates.
type Component interface {
	Render(w io.Writer) error
}

// ComponentFunc adapts a function into a Component.
type ComponentFunc func(w io.Writer) error

// Render implements the Component interface.
func (f ComponentFunc) Render(w io.Writer) error {
	return f(w)
}

// Attrs is a map for spread attributes.
type Attrs map[string]any

// Raw marks a string as safe HTML (will not be escaped).
type Raw string

// Render is a convenience function to render a component to an io.Writer.
func Render(w io.Writer, c Component) error {
	return c.Render(w)
}

// RenderToString renders a component and returns the result as a string.
func RenderToString(c Component) (string, error) {
	var buf bytes.Buffer
	if err := c.Render(&buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
