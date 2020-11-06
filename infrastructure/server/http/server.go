package http

import (
	"context"
	"log"
	"net/http"
	"time"
)

// Server represents a HTTP server instance
type Server struct {
	server *http.Server
}

// New create a new HTTP server
func New(port string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 55 * time.Second,
		},
	}
}

// ListenAndServe start the HTTP server with previously defined settings
func (s *Server) ListenAndServe() {
	go func() {
		log.Printf("beatport api running on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("error when starting the service: %q", err)
		}
	}()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown() {
	log.Printf("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Printf("unable to shut down the server in 60s: %q", err)
		return
	}
	log.Printf("server gracefully stopped")
}
