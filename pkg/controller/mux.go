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
	router := mux.NewRouter()

	router.HandleFunc("/", GetUnixTime).Methods("GET")
	//	router.HandleFunc("/metrics", TBDfunc).Methods("GET")
	//	router.HandleFunc("/health", TBDfunc).Methods("POST")
	// router.HandleFunc("/v1/tools/lookup", LookupDomain).Methods("GET")
	router.HandleFunc("/v1/tools/validate", ValidateIP).Methods("POST")
	// router.HandleFunc("/v1/history", RetrieveHistory).Methods("GET")

	return &Manager{Router: router, dbConnector: dbConnector}
}
