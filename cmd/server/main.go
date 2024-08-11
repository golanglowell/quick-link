package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golanglowell/quick-link/internal/application"
	"github.com/golanglowell/quick-link/internal/infrastructure/repository/memory"
	"github.com/golanglowell/quick-link/internal/presentation/resthttp"
	"github.com/golanglowell/quick-link/pkg/config"
	"github.com/golanglowell/quick-link/pkg/logger"
)

func run(
	ctx context.Context,
	getenv func(string) string,
	stderr io.Writer,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	config := &config.Config{
		Host: getenv("HOST"),
		Port: getenv("PORT"),
	}
	if config.Port == "" {
		config.Port = "8081"
	}

	logger := logger.NewLogger(os.Stdout)
	linkRepo := memory.NewURLRepository()
	defer linkRepo.Close()
	shortenLinkUC := application.NewShortenURL(linkRepo)
	getLinkUC := application.NewGetLink(linkRepo)

	srv := resthttp.NewServer(
		logger,
		config,
		shortenLinkUC,
		getLinkUC,
	)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}

	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(stderr, "error listening and serving: %s\n", err)
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		fmt.Fprintf(stderr, "error shutting down http server: %s\n", err)
	}

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Getenv, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
