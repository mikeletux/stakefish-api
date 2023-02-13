package main

import (
	"context"
	"github.com/mikeletux/stakefish-api/pkg/controller"
	"github.com/mikeletux/stakefish-api/pkg/debug"
	"github.com/mikeletux/stakefish-api/pkg/infra"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const gracefulShutdownTime int = 15

func main() {
	log := debug.NewBuiltinStdoutLogger()

	config := getEnvVars(log)

	dbController := infra.NewPostgresConnector(config.dbHost, config.dbUser, config.dbPass, config.dbName)
	networkInfra := infra.ImpInfra{}

	r := controller.NewController(dbController, networkInfra, log)

	srv := &http.Server{
		Handler:      r.Router,
		Addr:         config.apiListeningAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	log.Debug("Server started")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Debug("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Debug("HTTP Server stopped")

	dbController.Close() // Shutting down gracefully the database connector.
	log.Debug("Database Connector closed")
}

func getEnvVars(log debug.Logger) config {
	var config config
	config.dbHost = os.Getenv("FISH_PG_HOST")
	if len(config.dbHost) == 0 {
		log.Fatal("env var `FISH_PG_HOST` needs to be set to `hostname:port` for postgres")
	}

	config.dbUser = os.Getenv("FISH_PG_USER")
	if len(config.dbUser) == 0 {
		log.Fatal("env var `FISH_PG_USER` needs to be set to database username")
	}

	config.dbPass = os.Getenv("FISH_PG_PASS")
	if len(config.dbPass) == 0 {
		log.Fatal("env var `FISH_PG_PASS` needs to be set to database password")
	}

	config.dbName = os.Getenv("FISH_PG_DATABASE")
	if len(config.dbName) == 0 {
		log.Fatal("env var `FISH_PG_DATABASE` needs to be set to database name")
	}

	config.apiListeningAddr = os.Getenv("FISH_PG_API_ADDR")
	if len(config.apiListeningAddr) == 0 {
		log.Fatal("env var `FISH_PG_API_ADDR` needs to be set to rest API `listening_addr:listening_port`")
	}

	return config
}
