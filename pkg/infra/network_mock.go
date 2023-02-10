package infra

import (
	"fmt"
	"github.com/mikeletux/stakefish-api/pkg/models"
)

// MockInfra is a mock struct used for testing purposes.
type MockInfra struct {
	LookupIPv4AddrError    bool
	LookupIPv4AddrErrorMsg string
	Addresses              []models.Address
}

func (m MockInfra) LookupIPv4Addr(domain string) ([]models.Address, error) {
	if m.LookupIPv4AddrError {
		return nil, fmt.Errorf(m.LookupIPv4AddrErrorMsg)
	}
	return m.Addresses, nil
}
