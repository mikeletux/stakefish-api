package models

type UnixTime struct {
	Version   string `json:"version"`
	TimeStamp int64  `json:"date"`
	Isk8s     bool   `json:"kubernetes"`
}
