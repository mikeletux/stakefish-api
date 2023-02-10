package models

type UnixTime struct {
	Version   string `json:"version"`
	TimeStamp int64  `json:"date"`
	Isk8s     bool   `json:"kubernetes"`
}

type Address struct { // Address and ValidateIPRequest can be shared as both are the same in terms of data.
	Id int64  `json:"-" pg:",default,pk"`
	Ip string `json:"ip"`
}

type Query struct {
	tableName struct{} `json:"-" pg:"query"`
	Id        int64    `json:"-" pg:"id,default,pk"`
	// Addresses []Address `json:"addresses" pg:"rel:has-many"`
	ClientIp  string `json:"client_ip" pg:"client_ip"`
	CreatedAt int64  `json:"created_at" pg:"created_at"`
	Domain    string `json:"domain" pg:"domain"`
}

type ValidateIPResponse struct {
	Status bool `json:"status"`
}

type HTTPError struct {
	Message string `json:"message"`
}
