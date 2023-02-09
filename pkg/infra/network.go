package infra

import (
	"github.com/mikeletux/stakefish-api/pkg/models"
	"net"
)

type AccessInfra interface {
	LookupIPv4Addr(domain string) ([]models.Address, error)
}

type ImpInfra struct{}

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
