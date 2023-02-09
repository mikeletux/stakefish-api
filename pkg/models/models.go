package models

type UnixTime struct {
	Version   string `json:"version"`
	TimeStamp int64  `json:"date"`
	Isk8s     bool   `json:"kubernetes"`
}

type Address struct { // Address and ValidateIPRequest can be shared as both are the same in terms of data.
	Ip string `json:"ip"`
}

type Query struct {
	Addresses []Address `json:"addresses"`
	ClientIp  string    `json:"client_ip"`
	CreatedAt int64     `json:"created_at"`
	Domain    string    `json:"domain"`
}

type ValidateIPResponse struct {
	Status bool `json:"status"`
}

type HTTPError struct {
	Message string `json:"message"`
}
