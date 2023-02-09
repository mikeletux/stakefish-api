package controller

import (
	"github.com/mikeletux/stakefish-api/pkg/models"
	"net"
	"regexp"
	"time"
)

const ipRegExp string = `^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$`

func checkIfStringIsIP(ip string) bool {
	regexp := regexp.MustCompile(ipRegExp)
	return regexp.MatchString(ip)
}

func lookupIPv4Addr(domain string) ([]models.Address, error) {
	var ipv4 []models.Address
	ips, err := net.LookupIP(domain)

	if err != nil {
		return nil, err
	}

	for _, ip := range ips {
		if ip.To4() != nil { // This means ip struct is an IPv4
			ipv4 = append(ipv4, models.Address{Ip: ip.String()})
		}
	}

	return ipv4, nil
}

func retrieveDomainLookup(addresses []models.Address, domain string) models.Query {
	return models.Query{
		Addresses: addresses,
		ClientIp:  "192.168.1.2", // Change this :)
		CreatedAt: time.Now().Unix(),
		Domain:    domain,
	}
}
