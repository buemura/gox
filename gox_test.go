package gox

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestRender(t *testing.T) {
	t.Run("renders component to writer", func(t *testing.T) {
		c := ComponentFunc(func(w io.Writer) error {
			_, err := io.WriteString(w, "<h1>Hello</h1>")
			return err
		})

		var buf bytes.Buffer
		if err := Render(&buf, c); err != nil {
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
		c := ComponentFunc(func(w io.Writer) error {
			return wantErr
		})

		var buf bytes.Buffer
		if err := Render(&buf, c); !errors.Is(err, wantErr) {
			t.Errorf("got error %v, want %v", err, wantErr)
		}
	})
}

func TestRenderToString(t *testing.T) {
	t.Run("returns rendered HTML as string", func(t *testing.T) {
		c := ComponentFunc(func(w io.Writer) error {
			_, err := io.WriteString(w, "<p>Hello, World!</p>")
			return err
		})

		got, err := RenderToString(c)
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
		c := ComponentFunc(func(w io.Writer) error {
			return wantErr
		})

		got, err := RenderToString(c)
		if !errors.Is(err, wantErr) {
			t.Errorf("got error %v, want %v", err, wantErr)
		}
		if got != "" {
			t.Errorf("expected empty string on error, got %q", got)
		}
	})

	t.Run("renders empty component", func(t *testing.T) {
		c := ComponentFunc(func(w io.Writer) error {
			return nil
		})

		got, err := RenderToString(c)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "" {
			t.Errorf("expected empty string, got %q", got)
		}
	})
}

func TestRaw(t *testing.T) {
	t.Run("raw type preserves HTML content", func(t *testing.T) {
		raw := Raw("<script>alert('xss')</script>")

		c := ComponentFunc(func(w io.Writer) error {
			_, err := io.WriteString(w, string(raw))
			return err
		})

		got, err := RenderToString(c)
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
	t.Run("implements Component interface", func(t *testing.T) {
		var c Component = ComponentFunc(func(w io.Writer) error {
			_, err := io.WriteString(w, "<div>test</div>")
			return err
		})

		got, err := RenderToString(c)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := "<div>test</div>"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
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
