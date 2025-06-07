package server

import (
	"fmt"
	"net/http"

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

func (s *Server) Stop() {
	// ??
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}
