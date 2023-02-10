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

// LookupDomain handles queries coming to "/v1/tools/lookup" endpoint.
// If domain is provided and exists, returns a models.Query struct with the IPs associated with the domain,
// the client ip who made the request, the unix timestamp when the query was made and the domain name queried.
// It also saves the query in the storage specified in Manager.dbConnector.
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
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.HTTPError{Message: fmt.Sprintf("Domain %s was not found", domain)})
		return
	}

	if len(ips) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.HTTPError{Message: fmt.Sprintf("No IPv4 are associated with %s", domain)})
		return
	}

	domainLookup := retrieveDomainLookup(ips, domain, getIpFromAddressPort(r.RemoteAddr))

	err = m.dbConnector.SaveQuery(domainLookup)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.HTTPError{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(domainLookup)
}

// ValidateIP handles queries coming to "/v1/tools/validate" endpoint.
// This handler checks if the IP address provided as a parameter is an actual IPv4.
// If it is an IPv4 it returns true, false otherwise.
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

// RetrieveHistory handles queries coming to "/v1/history" endpoint.
// It returns the last 20 queries that were done in a descendent manner (last query done is the first one displayed).
func (m *Manager) RetrieveHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queries, err := m.dbConnector.RetrieveLastTwentyQueries()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.HTTPError{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(queries)
}
