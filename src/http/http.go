package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/knwoop/lets-gnomock/src/service"
)

type Server struct {
	userService *service.UserService

	addr   string
	mux    *http.ServeMux
	server *http.Server
}

func New(port int, userService *service.UserService) (*Server, error) {
	server := &Server{
		userService: userService,
		addr:        fmt.Sprintf(":%d", port),
		mux:         http.NewServeMux(),
	}
	server.registerHandlers()

	return server, nil
}

func (s *Server) Start() error {
	server := &http.Server{
		Addr:    s.addr,
		Handler: s.mux,
	}
	s.server = server

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) GracefulStop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) registerHandlers() {
	s.mux.Handle("/create", &createHandler{userService: s.userService})
	s.mux.Handle("/get", &getHandler{userService: s.userService})
}
