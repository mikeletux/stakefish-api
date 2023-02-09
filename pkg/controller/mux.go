package controller

import (
	"github.com/gorilla/mux"
	"github.com/mikeletux/stakefish-api/pkg/database"
	"net/http"
)

type Manager struct {
	Router http.Handler

	dbConnector database.Connector
}

func NewController(dbConnector database.Connector) *Manager {
	manager := &Manager{dbConnector: dbConnector}

	router := mux.NewRouter()

	router.HandleFunc("/", manager.GetUnixTime).Methods("GET")
	//	router.HandleFunc("/metrics", TBDfunc).Methods("GET")
	//	router.HandleFunc("/health", TBDfunc).Methods("POST")
	router.HandleFunc("/v1/tools/lookup", manager.LookupDomain).Methods("GET")
	router.HandleFunc("/v1/tools/validate", manager.ValidateIP).Methods("POST")
	// router.HandleFunc("/v1/history", RetrieveHistory).Methods("GET")

	manager.Router = router

	return manager
}
