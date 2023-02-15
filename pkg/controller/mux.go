package controller

import (
	"github.com/gorilla/mux"
	"github.com/mikeletux/stakefish-api/pkg/debug"
	"github.com/mikeletux/stakefish-api/pkg/infra"
	"net/http"
)

// Manager handles the mux router creation and handles dependencies to access the internet and the storage backend.
type Manager struct {
	Router http.Handler

	dbConnector  infra.DBConnector
	networkInfra infra.AccessInfra
	log          debug.Logger
}

func NewController(dbConnector infra.DBConnector, networkInfra infra.AccessInfra, log debug.Logger) *Manager {
	manager := &Manager{
		dbConnector:  dbConnector,
		networkInfra: networkInfra,
		log:          log,
	}

	router := mux.NewRouter()

	router.HandleFunc("/", manager.GetUnixTime).Methods("GET")
	//	router.HandleFunc("/metrics", TBDfunc).Methods("GET")
	router.HandleFunc("/health", manager.HealthCheck).Methods("GET")
	router.HandleFunc("/v1/tools/lookup", manager.LookupDomain).Methods("GET")
	router.HandleFunc("/v1/tools/validate", manager.ValidateIP).Methods("POST")
	router.HandleFunc("/v1/history", manager.RetrieveHistory).Methods("GET")

	manager.Router = router

	return manager
}
