package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	srv *http.Server
}

func NewServer(port int, r *gin.Engine) *Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	return &Server{
		srv: srv,
	}
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown error: %w", err)
	}

	return nil
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}
