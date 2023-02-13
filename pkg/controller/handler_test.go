package controller

import (
	"bytes"
	"encoding/json"
	"github.com/mikeletux/stakefish-api/pkg/debug"
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

	manager := NewController(db, mockInfra, debug.NewBuiltinStdoutLogger())

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

func TestValidateIP(t *testing.T) {
	testTable := []struct {
		ipAddr string
		isIp   bool
	}{
		{
			ipAddr: "192.168.10.1",
			isIp:   true,
		},
		{
			ipAddr: "256.500.1.3",
			isIp:   false,
		},
		{
			ipAddr: "15.76.45.1",
			isIp:   true,
		},
		{
			ipAddr: "0.300.1.2",
			isIp:   false,
		},
	}

	db := infra.NewMockConnector()
	mockInfra := infra.MockInfra{}

	manager := NewController(db, mockInfra, debug.NewBuiltinStdoutLogger())

	for _, test := range testTable {
		requestBody, err := json.Marshal(models.Address{Ip: test.ipAddr})
		if err != nil {
			t.Errorf("error when marshalling struct models.Address to byte")
		}

		r, _ := http.NewRequest("POST", "/v1/tools/validate", bytes.NewReader(requestBody))
		w := httptest.NewRecorder()

		manager.ValidateIP(w, r)

		var validateIp models.ValidateIPResponse
		err = json.NewDecoder(w.Body).Decode(&validateIp)
		if err != nil {
			t.Errorf("error when decoding JSON from response")
		}

		if validateIp.Status != test.isIp {
			t.Errorf("ip check for %s is %t when it was supposed to be %t", test.ipAddr, validateIp.Status, test.isIp)
		}
	}

}

func TestRetrieveHistory(t *testing.T) {}
