package zex

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Server is the default server for the application
type Server struct {
	app *App
}

// NewServer creates a new Server instance
func NewServer(app *App) *Server {
	return &Server{app}
}

// ServeHTTP is the main handler for the server
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// log the current request information in development mode
	defer func(start time.Time, method string, path string, dev bool) {
		if !dev {
			return
		}
		serverLogger(start, method, path)
	}(start, r.Method, r.URL.Path, s.app.conf.Development)

	if s.handlePublic(w, r) {
		return
	}

	notFound := s.app.conf.NotFoundHandler
	finalHandler := func(w http.ResponseWriter, r *http.Request) {
		for _, route := range s.app.exportRoutes() {
			method := route.Method()
			if method == r.Method || method == "*" {
				ok, params := route.comparePath(s.app.CompleteRouter, r.URL.Path)
				if !ok {
					continue
				}

				s.handleRoute(route, w, r, params)
				return
			}
		}
		notFound(w, r)
	}

	chainHandler := s.chainMiddlewares(finalHandler, s.app.middlewares...)
	chainHandler(w, r)
}

// handleRoute handles the route
func (s *Server) handleRoute(route Route, w http.ResponseWriter, r *http.Request, params map[string]string) {
	r = r.WithContext(context.WithValue(r.Context(), ContextParams, params))
	handler := s.chainMiddlewares(route.Handler(), route.Middlewares()...)
	handler(w, r)
}

// chainMiddlewares chains the middlewares
func (s *Server) chainMiddlewares(handler http.HandlerFunc, middlewares ...MiddlewareFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// handlePublic handles the public files
func (s *Server) handlePublic(w http.ResponseWriter, r *http.Request) bool {
	for prefix, path := range s.app.public {
		if strings.HasPrefix(r.URL.Path, prefix) {
			relativePath := strings.TrimPrefix(r.URL.Path, prefix)
			fullPath := filepath.Join(path, relativePath)

			if stat, err := os.Stat(fullPath); err == nil && !stat.IsDir() {
				http.ServeFile(w, r, fullPath)
				return true
			}
		}
	}
	return false
}
