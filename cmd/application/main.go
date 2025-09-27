package main

import (
	"context"
	"log"
	"os"

	"golang.org/x/sync/errgroup"

	_ "document_manager/api"
	"document_manager/config"
	"document_manager/internal/application"
	"document_manager/internal/application/infrastructure/server"
)

const (
	configPathEnv = "CONFIG_PATH"
)

func main() {
	ctx := context.Background()

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Got unhandled error: %v", r)
			os.Exit(1)
		}
	}()

	configPath := os.Getenv(configPathEnv)
	if configPath == "" {
		log.Printf("Failed to get key for env %s", configPathEnv)
		os.Exit(1)
	}

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Printf("Failed to get config. Reason: %v", err)
		os.Exit(1)
	}

	srv := server.New()

	if err := application.Init(ctx, srv, *cfg); err != nil {
		log.Printf("Failed to init application. Reason: %v", err)
		os.Exit(1)
	}

	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := srv.Start(); err != nil {
			log.Printf("Failed to listen and serve: %v", err)
			return err
		}
		return nil
	})

	eg.Go(func() error {
		<-egCtx.Done()
		if err := srv.Shutdown(); err != nil {
			log.Printf("Failed to shutdown server: %v", err)
			return err
		}
		return nil
	})

	// Ignore error, because errors are logged in the goroutines
	_ = eg.Wait()

	log.Println("Application Shutdown")

}
