package controller

import (
	"encoding/json"
	"fmt"
	"github.com/mikeletux/stakefish-api/pkg/models"
	"net/http"
	"strings"
	"time"
)

// GetUnixTime handles queries coming to "/" endpoint.
// Returns a JSON object with the API version, UNIX timestamp and if the API is running in k8s.
// TO-DO get app version better and figure out if the app is running in k8s using env vars.
func (m *Manager) GetUnixTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var unixTime models.UnixTime

	unixTime.Version = "0.1.0"
	unixTime.TimeStamp = time.Now().Unix()
	unixTime.Isk8s = false

	json.NewEncoder(w).Encode(unixTime)
}

func (m *Manager) LookupDomain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	domain := r.URL.Query().Get("domain")
	if len(strings.TrimSpace(domain)) == 0 { // If no domain is provided, inform client.
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.HTTPError{Message: "No domain provided"})
		return
	}

	ips, err := m.networkInfra.LookupIPv4Addr(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.HTTPError{Message: "No domain provided"})
		return
	}

	if len(ips) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.HTTPError{Message: fmt.Sprintf("No IPv4 are associated with %s", domain)})
		return
	}

	// DONT FORGET TO ADD IT TO THE DATABASE!

	domainLookup := retrieveDomainLookup(ips, domain)
	json.NewEncoder(w).Encode(domainLookup)
}

func (m *Manager) ValidateIP(w http.ResponseWriter, r *http.Request) {
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

func (m *Manager) RetrieveHistory(w http.ResponseWriter, r *http.Request) {
	// TO-DO
}
