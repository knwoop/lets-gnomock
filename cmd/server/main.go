package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/knwoop/lets-gnomock/src/config"
	"github.com/knwoop/lets-gnomock/src/db"

	"github.com/knwoop/lets-gnomock/src/http"
	"github.com/knwoop/lets-gnomock/src/service"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] %s", err.Error())
		os.Exit(1)
	}
}

func run() error {
	env, err := config.ReadFromEnv()
	if err != nil {
		return fmt.Errorf("failed to read env vars: %w", err)
	}
	conn, err := db.NewClient(env)
	if err != nil {
		return fmt.Errorf("failed New Database Client: %w", err)
	}
	s, err := service.NewUserService(conn)
	if err != nil {
		return fmt.Errorf("failed to new user service: %w", err)
	}

	httpserver, err := http.New(8080, s)
	if err != nil {
		return fmt.Errorf("failed to new http server: %w", err)
	}
	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGTERM, syscall.SIGINT)
	errCh := make(chan error, 1)

	go func() {
		errCh <- httpserver.Start()
	}()

	select {
	case <-termCh:
		return fmt.Errorf("failed to terminate server")
	case <-errCh:
		return fmt.Errorf("failed to serve http server")
	}

	ctx := context.Background()
	if err := httpserver.GracefulStop(ctx); err != nil {
		return err
	}
	return nil
}
