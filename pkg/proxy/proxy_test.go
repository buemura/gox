package proxy

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestInjectReloadScript(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `<html><body><h1>Hello</h1></body></html>`)
	}))
	defer upstream.Close()

	srv, err := New(upstream.URL, ":0")
	if err != nil {
		t.Fatal(err)
	}

	// Use httptest to test the proxy handler directly.
	proxyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.proxy.ServeHTTP(w, r)
	}))
	defer proxyServer.Close()

	resp, err := http.Get(proxyServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var buf strings.Builder
	io.Copy(&buf, resp.Body)
	body := buf.String()

	if !strings.Contains(body, "__gox_reload") {
		t.Error("expected reload script to be injected into HTML response")
	}
	if !strings.Contains(body, "</body>") {
		t.Error("expected </body> tag to be preserved")
	}
	// Script should appear before </body>.
	scriptIdx := strings.Index(body, "__gox_reload")
	bodyIdx := strings.Index(body, "</body>")
	if scriptIdx > bodyIdx {
		t.Error("expected reload script to be injected before </body>")
	}
}

func TestNoInjectionForNonHTML(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":"ok"}`)
	}))
	defer upstream.Close()

	srv, err := New(upstream.URL, ":0")
	if err != nil {
		t.Fatal(err)
	}

	proxyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.proxy.ServeHTTP(w, r)
	}))
	defer proxyServer.Close()

	resp, err := http.Get(proxyServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var buf strings.Builder
	io.Copy(&buf, resp.Body)
	body := buf.String()

	if strings.Contains(body, "__gox_reload") {
		t.Error("should not inject reload script into non-HTML responses")
	}
	if body != `{"status":"ok"}` {
		t.Errorf("unexpected body: %s", body)
	}
}

func TestSSEReload(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	}))
	defer upstream.Close()

	srv, err := New(upstream.URL, ":0")
	if err != nil {
		t.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/__gox_reload", srv.handleSSE)
	mux.Handle("/", srv.proxy)
	proxyServer := httptest.NewServer(mux)
	defer proxyServer.Close()

	// Connect to SSE endpoint.
	resp, err := http.Get(proxyServer.URL + "/__gox_reload")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if ct := resp.Header.Get("Content-Type"); ct != "text/event-stream" {
		t.Fatalf("expected Content-Type text/event-stream, got %s", ct)
	}

	// Trigger a reload.
	go func() {
		time.Sleep(50 * time.Millisecond)
		srv.Reload()
	}()

	scanner := bufio.NewScanner(resp.Body)
	deadline := time.After(2 * time.Second)
	gotReload := false

	for {
		select {
		case <-deadline:
			if !gotReload {
				t.Fatal("timed out waiting for reload event")
			}
			return
		default:
		}
		if scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "event: reload") {
				gotReload = true
				return
			}
		} else {
			break
		}
	}

	if !gotReload {
		t.Fatal("did not receive reload event")
	}
}

func TestReloadMultipleListeners(t *testing.T) {
	srv, err := New("http://localhost:9999", ":0")
	if err != nil {
		t.Fatal(err)
	}

	// Simulate multiple listeners.
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)

	srv.mu.Lock()
	srv.listeners[ch1] = struct{}{}
	srv.listeners[ch2] = struct{}{}
	srv.mu.Unlock()

	srv.Reload()

	select {
	case <-ch1:
	case <-time.After(100 * time.Millisecond):
		t.Error("listener 1 did not receive reload")
	}

	select {
	case <-ch2:
	case <-time.After(100 * time.Millisecond):
		t.Error("listener 2 did not receive reload")
	}
}
