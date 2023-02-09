package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewController() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", GetUnixTime).Methods("GET")
	//	router.HandleFunc("/metrics", TBDfunc).Methods("GET")
	//	router.HandleFunc("/health", TBDfunc).Methods("POST")
	//	router.HandleFunc("/v1/tools/lookup", TBDfunc).Methods("PUT")
	//	router.HandleFunc("/v1/tools/validate", TBDfunc).Methods("DELETE")
	//	router.HandleFunc("/v1/history", TBDfunc).Methods("DELETE")

	return router
}
