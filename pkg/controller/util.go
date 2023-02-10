package controller

import (
	"github.com/mikeletux/stakefish-api/pkg/models"
	"regexp"
	"time"
)

const ipRegExp string = `^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$`

func checkIfStringIsIP(ip string) bool {
	regexp := regexp.MustCompile(ipRegExp)
	return regexp.MatchString(ip)
}

func retrieveDomainLookup(addresses []models.Address, domain string) models.Query {
	return models.Query{
		// Addresses: addresses,
		ClientIp:  "192.168.1.2", // Change this :)
		CreatedAt: time.Now().Unix(),
		Domain:    domain,
	}
}
