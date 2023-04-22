package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/hytkgami/log-group-sample-go/controllers"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("failed to run: %+v", err)
	}
}

func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := http.NewServeMux()
	healthCheckController := controllers.NewHealthCheckController()
	router.HandleFunc("/ping", healthCheckController.Ping)
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		log.Println("Listening on port", port)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	return eg.Wait()
}
