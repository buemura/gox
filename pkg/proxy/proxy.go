// Package proxy provides a reverse proxy with hot-reload support for gox watch.
// It proxies requests to the upstream application server and injects a small
// script into HTML responses that listens for Server-Sent Events (SSE) to
// trigger automatic browser reloads after successful recompilation.
package proxy

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
)

// reloadScript is injected before </body> in HTML responses.
// It connects to the SSE endpoint and reloads the page on "reload" events.
const reloadScript = `<script>
(function(){var es=new EventSource("/__gox_reload");es.addEventListener("reload",function(){window.location.reload()});es.onerror=function(){es.close();setTimeout(function(){window.location.reload()},1000)};})();
</script>`

// Server is a reverse proxy that injects hot-reload support into HTML responses.
type Server struct {
	proxy    *httputil.ReverseProxy
	upstream *url.URL
	addr     string

	mu        sync.Mutex
	listeners map[chan struct{}]struct{}
}

// New creates a new proxy server.
// upstream is the URL of the application server (e.g. "http://localhost:8080").
// listenAddr is the address to listen on (e.g. ":8081").
func New(upstream string, listenAddr string) (*Server, error) {
	u, err := url.Parse(upstream)
	if err != nil {
		return nil, fmt.Errorf("parsing upstream URL: %w", err)
	}

	s := &Server{
		upstream:  u,
		addr:      listenAddr,
		listeners: make(map[chan struct{}]struct{}),
	}

	s.proxy = &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = u.Scheme
			req.URL.Host = u.Host
			req.Host = u.Host
		},
		ModifyResponse: s.injectReloadScript,
	}

	return s, nil
}

// Addr returns the listen address.
func (s *Server) Addr() string {
	return s.addr
}

// Reload notifies all connected browsers to reload.
func (s *Server) Reload() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for ch := range s.listeners {
		select {
		case ch <- struct{}{}:
		default:
		}
	}
}

// ListenAndServe starts the proxy server. It blocks until the server is closed.
func (s *Server) ListenAndServe() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	// Update addr with the actual address (useful when port is 0).
	s.addr = ln.Addr().String()
	mux := http.NewServeMux()
	mux.HandleFunc("/__gox_reload", s.handleSSE)
	mux.Handle("/", s.proxy)
	return http.Serve(ln, mux)
}

// handleSSE serves the Server-Sent Events endpoint for reload notifications.
func (s *Server) handleSSE(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	flusher.Flush()

	ch := make(chan struct{}, 1)
	s.mu.Lock()
	s.listeners[ch] = struct{}{}
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.listeners, ch)
		s.mu.Unlock()
	}()

	for {
		select {
		case <-r.Context().Done():
			return
		case <-ch:
			fmt.Fprintf(w, "event: reload\ndata: reload\n\n")
			flusher.Flush()
		}
	}
}

// injectReloadScript modifies HTML responses to include the reload script.
func (s *Server) injectReloadScript(resp *http.Response) error {
	ct := resp.Header.Get("Content-Type")
	if !strings.Contains(ct, "text/html") {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	html := string(body)

	// Inject before </body> if present, otherwise append.
	if idx := strings.LastIndex(strings.ToLower(html), "</body>"); idx != -1 {
		html = html[:idx] + reloadScript + html[idx:]
	} else {
		html += reloadScript
	}

	resp.Body = io.NopCloser(strings.NewReader(html))
	resp.ContentLength = int64(len(html))
	resp.Header.Set("Content-Length", fmt.Sprintf("%d", len(html)))
	return nil
}
