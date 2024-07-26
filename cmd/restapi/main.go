package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/bcdxn/go-todo/pkg/config"
	"github.com/bcdxn/go-todo/pkg/rest"
	"github.com/bcdxn/go-todo/pkg/todo"
	"github.com/hashicorp/go-hclog"
)

// main is the entry point into our program that serves a REST API with  To-Do management
// capabilities.
func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig(config.OptionsNewConfig{
		FilePath: os.Args[1],
	})
	if err != nil {
		log.Fatalf("%s\n", err)
		os.Exit(1)
	}

	if err := run(ctx, os.Stdout, os.Stderr, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	fmt.Println("exited successfully")
}

// run initializes dependences and starts the server. It is meant to be called by the application's
// entrypoint, `main`.
func run(
	ctx context.Context,
	stdout, stderr io.Writer,
	cfg config.Config,
) error {
	l := hclog.New(&hclog.LoggerOptions{
		Name:       "gotodo",
		Level:      hclog.LevelFromString(cfg.Logger.Level),
		Output:     stdout,
		JSONFormat: cfg.Logger.Format == "json",
	})

	restServer := rest.NewServer(
		cfg,
		l,
		todo.NewStaticService(l),
	)

	go func() {
		log.Printf("listening on %s\n", restServer.Addr)
		if err := restServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(stderr, "error listening and serving: %s\n", err)
		}
	}()

	waitForInterrupt(ctx, restServer)

	return nil
}

// waitForInterrupt will keep the main Go process running until an OS interrupt signal (e.g. ctrl+c)
// is received. It will then gracefully shutdown the HTTP server at the core of the REST API by
// allowing connections to be drained. A timeout will eventually kill the server if connections are
// not drained fast enough.
func waitForInterrupt(ctx context.Context, httpServer *http.Server) {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
}
