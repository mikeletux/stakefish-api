package infra

import (
	"github.com/mikeletux/stakefish-api/pkg/models"
	"net"
)

// AccessInfra is the interface that must be implemented to access the internet.
type AccessInfra interface {
	LookupIPv4Addr(domain string) ([]models.Address, error)
}

// ImpInfra implements AccessInfra interface and gives access to the actual internet.
type ImpInfra struct{}

// LookupIPv4Addr retrieves all IPv4 address from a domain name accessing the actual internet.
func (i ImpInfra) LookupIPv4Addr(domain string) ([]models.Address, error) {
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
