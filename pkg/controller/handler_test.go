package controller

import (
	"encoding/json"
	"github.com/mikeletux/stakefish-api/pkg/infra"
	"github.com/mikeletux/stakefish-api/pkg/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const unixPastTime = 423705600 // June 6th, 1983 00:00:00

func TestGetUnixTime(t *testing.T) {
	db := infra.NewMockConnector()
	mockInfra := infra.MockInfra{}

	manager := NewController(db, mockInfra)

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	manager.GetUnixTime(w, r)

	result := reflect.DeepEqual(w.Code, http.StatusOK)
	if !result {
		t.Errorf("got: %d expected: %d for HTTP response code", w.Code, http.StatusOK)
	}

	var unixTime models.UnixTime
	err := json.NewDecoder(w.Body).Decode(&unixTime)
	if err != nil {
		t.Errorf("error decoding JSON response")
	}

	// Check version

	if unixPastTime > unixTime.TimeStamp {
		t.Errorf("got timestamp is not correct. Expected to be greater than %d", unixPastTime)
	}

	if unixTime.Isk8s {
		t.Errorf("unit tests shouldn't be running in k8s.")
	}
}

func TestLookupDomain(t *testing.T) {}

func TestValidateIP(t *testing.T) {}

func TestRetrieveHistory(t *testing.T) {}
