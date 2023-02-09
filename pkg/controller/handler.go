package controller

import (
	"encoding/json"
	"github.com/mikeletux/stakefish-api/pkg/models"
	"net/http"
	"time"
)

// GetUnixTime handles queries coming to "/" endpoint.
// Returns a JSON object with the API version, UNIX timestamp and if the API is running in k8s.
// TO-DO get app version better and figure out if the app is running in k8s using env vars.
func GetUnixTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var unixTime models.UnixTime

	unixTime.Version = "0.1.0"
	unixTime.TimeStamp = time.Now().Unix()
	unixTime.Isk8s = false

	json.NewEncoder(w).Encode(unixTime)
	w.WriteHeader(http.StatusOK)
}

func LookupDomain(w http.ResponseWriter, r *http.Request) {
	// TO-DO
}

func ValidateIP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ipAddr models.Address
	err := json.NewDecoder(r.Body).Decode(&ipAddr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.HTTPError{Message: "Error decoding incoming body"})
		return
	}

	result := checkIfStringIsIP(ipAddr.Ip) // TO-DO
	json.NewEncoder(w).Encode(models.ValidateIPResponse{Status: result})
}

func RetrieveHistory(w http.ResponseWriter, r *http.Request) {
	// TO-DO
}
