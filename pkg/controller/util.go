package controller

import (
	"fmt"
	"github.com/mikeletux/stakefish-api/pkg/models"
	"net/url"
	"os"
	"regexp"
	"time"
)

const ipRegExp string = `^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$`

func checkIfStringIsIP(ip string) bool {
	regexp := regexp.MustCompile(ipRegExp)
	return regexp.MatchString(ip)
}

func retrieveDomainLookup(addresses []models.Address, domain string, clientIP string) models.Query {
	return models.Query{
		Addresses: addresses,
		ClientIp:  clientIP,
		CreatedAt: time.Now().Unix(),
		Domain:    domain,
	}
}

func getIpFromAddressPort(addressPort string) string {
	u, err := url.Parse(fmt.Sprintf("http://%s", addressPort))
	if err != nil {
		return addressPort
	}

	return u.Hostname()
}

func isRunningInK8s() bool {
	k8sServiceHost := os.Getenv("KUBERNETES_SERVICE_HOST")
	if len(k8sServiceHost) > 0 {
		return true
	}

	return false
}
