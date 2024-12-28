package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/bcdxn/go-todo/internal/rest"
)

func main() {
	ctx := context.Background()
	logger := slog.Default()

	srv := rest.NewServer(
		logger,
	)
	httpServer := &http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", "3000"),
		Handler: srv,
	}
	go func() {
		logger.Info("server listening", "address", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		} else {
			logger.Info("graceful shutdown complete")
		}
	}()
	wg.Wait()
}
