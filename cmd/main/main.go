package main

import (
	"github.com/mikeletux/stakefish-api/pkg/controller"
	"github.com/mikeletux/stakefish-api/pkg/database"
	"log"
	"net/http"
	"time"
)

func main() {
	dbController := database.NewPostgresConnector()

	r := controller.NewController(dbController)

	srv := &http.Server{
		Handler: r.Router,
		Addr:    "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Bind to a port and pass our router in
	log.Fatal(srv.ListenAndServe())
}
