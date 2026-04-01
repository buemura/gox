package gox

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"
)

func TestRender(t *testing.T) {
	ctx := context.Background()

	t.Run("renders component to writer", func(t *testing.T) {
		c := ComponentFunc(func(ctx context.Context, w io.Writer) error {
			_, err := io.WriteString(w, "<h1>Hello</h1>")
			return err
		})

		var buf bytes.Buffer
		if err := Render(ctx, &buf, c); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := buf.String()
		want := "<h1>Hello</h1>"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("propagates write errors", func(t *testing.T) {
		wantErr := errors.New("write failed")
		c := ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return wantErr
		})

		var buf bytes.Buffer
		if err := Render(ctx, &buf, c); !errors.Is(err, wantErr) {
			t.Errorf("got error %v, want %v", err, wantErr)
		}
	})
}

func TestRenderToString(t *testing.T) {
	ctx := context.Background()

	t.Run("returns rendered HTML as string", func(t *testing.T) {
		c := ComponentFunc(func(ctx context.Context, w io.Writer) error {
			_, err := io.WriteString(w, "<p>Hello, World!</p>")
			return err
		})

		got, err := RenderToString(ctx, c)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := "<p>Hello, World!</p>"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns error on failure", func(t *testing.T) {
		wantErr := errors.New("render failed")
		c := ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return wantErr
		})

		got, err := RenderToString(ctx, c)
		if !errors.Is(err, wantErr) {
			t.Errorf("got error %v, want %v", err, wantErr)
		}
		if got != "" {
			t.Errorf("expected empty string on error, got %q", got)
		}
	})

	t.Run("renders empty component", func(t *testing.T) {
		c := ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return nil
		})

		got, err := RenderToString(ctx, c)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "" {
			t.Errorf("expected empty string, got %q", got)
		}
	})
}

func TestRaw(t *testing.T) {
	ctx := context.Background()

	t.Run("raw type preserves HTML content", func(t *testing.T) {
		raw := Raw("<script>alert('xss')</script>")

		c := ComponentFunc(func(ctx context.Context, w io.Writer) error {
			_, err := io.WriteString(w, string(raw))
			return err
		})

		got, err := RenderToString(ctx, c)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := "<script>alert('xss')</script>"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestComponentFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("implements Component interface", func(t *testing.T) {
		var c Component = ComponentFunc(func(ctx context.Context, w io.Writer) error {
			_, err := io.WriteString(w, "<div>test</div>")
			return err
		})

		got, err := RenderToString(ctx, c)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := "<div>test</div>"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestSanitizeURL(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"safe http URL", "https://example.com", "https://example.com"},
		{"safe relative URL", "/path/to/page", "/path/to/page"},
		{"blocks javascript scheme", "javascript:alert('xss')", "about:invalid#GoxBlockedURL"},
		{"blocks JavaScript mixed case", "JavaScript:alert(1)", "about:invalid#GoxBlockedURL"},
		{"blocks vbscript scheme", "vbscript:msgbox", "about:invalid#GoxBlockedURL"},
		{"blocks data scheme", "data:text/html,<script>alert(1)</script>", "about:invalid#GoxBlockedURL"},
		{"blocks with leading spaces", "  javascript:alert(1)", "about:invalid#GoxBlockedURL"},
		{"escapes HTML in URL", "https://example.com?a=1&b=2", "https://example.com?a=1&amp;b=2"},
		{"SafeURL bypasses sanitization", SafeURL("javascript:trusted()"), "javascript:trusted()"},
		{"SafeURL still HTML-escapes", SafeURL("https://example.com?a=1&b=2"), "https://example.com?a=1&amp;b=2"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SanitizeURL(tt.input)
			if got != tt.want {
				t.Errorf("SanitizeURL(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestSanitizeCSS(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"safe color value", "color: red", "color: red"},
		{"safe margin value", "margin: 10px", "margin: 10px"},
		{"blocks expression()", "width: expression(alert(1))", ""},
		{"blocks url()", "background: url(javascript:alert(1))", ""},
		{"blocks behavior", "behavior: url(xss.htc)", ""},
		{"blocks -moz-binding", "-moz-binding: url(xss.xml#xss)", ""},
		{"blocks javascript in CSS", "background: javascript:alert(1)", ""},
		{"blocks mixed case", "width: Expression(alert(1))", ""},
		{"escapes HTML entities", "content: '<script>'", "content: &#39;&lt;script&gt;&#39;"},
		{"SafeCSS bypasses sanitization", SafeCSS("background: url(safe.png)"), "background: url(safe.png)"},
		{"SafeCSS still HTML-escapes", SafeCSS("content: '<b>'"), "content: &#39;&lt;b&gt;&#39;"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SanitizeCSS(tt.input)
			if got != tt.want {
				t.Errorf("SanitizeCSS(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestAttrs(t *testing.T) {
	t.Run("can store and retrieve attributes", func(t *testing.T) {
		attrs := Attrs{
			"class":    "main",
			"id":       "app",
			"disabled": true,
		}

		if attrs["class"] != "main" {
			t.Errorf("got %v, want %q", attrs["class"], "main")
		}
		if attrs["id"] != "app" {
			t.Errorf("got %v, want %q", attrs["id"], "app")
		}
		if attrs["disabled"] != true {
			t.Errorf("got %v, want true", attrs["disabled"])
		}
	})
}
