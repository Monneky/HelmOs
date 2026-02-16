package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"helmos/api/internal/handler"
)

// Server wraps the HTTP server and router.
type Server struct {
	httpServer *http.Server
	mux        *http.ServeMux
}

// New creates a new Server listening on addr.
func New(addr string) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /db-check", handler.DBCheck)

	return &Server{
		mux: mux,
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      corsMiddleware(mux),
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

// corsMiddleware adds CORS headers to all responses.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Start starts the HTTP server (blocking).
func (s *Server) Start() error {
	log.Printf("@helmos/api listening on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
