package gox

import (
	"bytes"
	"context"
	"fmt"
	"html"
	"io"
	"strings"
)

// Component is the interface returned by compiled templates.
type Component interface {
	Render(ctx context.Context, w io.Writer) error
}

// ComponentFunc adapts a function into a Component.
type ComponentFunc func(ctx context.Context, w io.Writer) error

// Render implements the Component interface.
func (f ComponentFunc) Render(ctx context.Context, w io.Writer) error {
	return f(ctx, w)
}

// Attrs is a map for spread attributes.
type Attrs map[string]any

// Raw marks a string as safe HTML (will not be escaped).
type Raw string

// Render is a convenience function to render a component to an io.Writer.
func Render(ctx context.Context, w io.Writer, c Component) error {
	return c.Render(ctx, w)
}

// RenderToString renders a component and returns the result as a string.
func RenderToString(ctx context.Context, c Component) (string, error) {
	var buf bytes.Buffer
	if err := c.Render(ctx, &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// SafeURL marks a URL string as trusted, bypassing URL sanitization.
// Use this only when the URL source is known to be safe.
type SafeURL string

// SafeCSS marks a CSS string as trusted, bypassing CSS sanitization.
// Use this only when the CSS source is known to be safe.
type SafeCSS string

// blockedURLSchemes are URL schemes that can execute code.
var blockedURLSchemes = []string{"javascript:", "vbscript:", "data:"}

// dangerousCSSPatterns are CSS patterns that can be used for injection attacks.
var dangerousCSSPatterns = []string{"expression(", "url(", "behavior:", "-moz-binding:", "javascript:"}

// SanitizeURL sanitizes a URL value for use in href, src, action, and formaction attributes.
// It blocks javascript:, vbscript:, and data: URL schemes by replacing them with a safe value.
// Values of type SafeURL bypass sanitization.
func SanitizeURL(v any) string {
	if safe, ok := v.(SafeURL); ok {
		return html.EscapeString(string(safe))
	}
	s := fmt.Sprintf("%v", v)
	normalized := strings.TrimSpace(strings.ToLower(s))
	for _, scheme := range blockedURLSchemes {
		if strings.HasPrefix(normalized, scheme) {
			return "about:invalid#GoxBlockedURL"
		}
	}
	return html.EscapeString(s)
}

// SanitizeCSS sanitizes a CSS value for use in style attributes.
// It blocks dangerous CSS properties and values such as expression(), url(), and behavior.
// Values of type SafeCSS bypass sanitization.
func SanitizeCSS(v any) string {
	if safe, ok := v.(SafeCSS); ok {
		return html.EscapeString(string(safe))
	}
	s := fmt.Sprintf("%v", v)
	normalized := strings.TrimSpace(strings.ToLower(s))
	for _, pattern := range dangerousCSSPatterns {
		if strings.Contains(normalized, pattern) {
			return ""
		}
	}
	return html.EscapeString(s)
}
