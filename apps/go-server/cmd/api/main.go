package main

import (
	"axon/server/internal/config"
	"axon/server/internal/http"
	"axon/server/pkg/db"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func run(ctx context.Context, cfg config.Config) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	log.Println("Starting up...")

	env := cfg.Env
	addr := fmt.Sprintf("%s:%s", env.HOST, env.PORT)

	// Initialize database pool
	pool, err := db.New(ctx, cfg.Env.DATABASE_URL)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	defer db.Close(pool)
	log.Println("Database connected")

	server, err := http.New(addr, &cfg)
	if err != nil {
		return err
	}

	// start listening in goroutine
	listenErr := make(chan error, 1)
	go func() {
		listenErr <- server.Start()
	}()

	// wait for signal or listen error
	select {
	case <-ctx.Done():
		log.Println("Initiating shutdown process...")
	case err := <-listenErr:
		if err != nil && err.Error() != "server closed" {
			return err
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- server.Shutdown(shutdownCtx)
	}()

	select {
	case err := <-done:
		if err != nil {
			return err
		}
	case <-shutdownCtx.Done():
		return shutdownCtx.Err()
	}

	log.Println("Server gracefully shutdown")
	return nil
}

func main() {
	ctx := context.Background()

	dotenvFlag := flag.Bool("dotenv", true, "skip with --dotenv=false")
	flag.Parse()

	if *dotenvFlag {
		log.Println("Loading .env file")
		if err := godotenv.Load(); err != nil {
			log.Fatalf(".env not loaded: %v", err)
		}
	} else {
		log.Println("Skipping .env loading")
	}

	cfg := config.Init()

	if err := run(ctx, cfg); err != nil {
		log.Fatalln(err)
	}
}
