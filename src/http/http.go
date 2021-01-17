package http

import (
	"context"
	"net"
	"net/http"

	"github.com/knwoop/lets-gnomock/src/service"
)

type Server struct {
	userService *service.UserService

	mux    *http.ServeMux
	server *http.Server
}

func New(userService *service.UserService) (*Server, error) {
	server := &Server{
		userService: userService,
		mux:         http.NewServeMux(),
	}
	return server, nil
}

func (s *Server) Serve(ln net.Listener) error {
	server := &http.Server{
		Handler: s.mux,
	}
	s.server = server

	if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) GracefulStop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) registerHandlers() {
	s.mux.Handle("/services/users", &createHandler{})
}
