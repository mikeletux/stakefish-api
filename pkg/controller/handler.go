package controller

import (
	"encoding/json"
	"github.com/mikeletux/stakefish-api/pkg/models"
	"net/http"
	"time"
)

func GetUnixTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var unixTime models.UnixTime

	unixTime.Version = "0.1.0"
	unixTime.TimeStamp = time.Now().Unix()
	unixTime.Isk8s = false

	json.NewEncoder(w).Encode(unixTime)
}
