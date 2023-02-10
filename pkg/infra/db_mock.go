package infra

import (
	"fmt"
	"github.com/mikeletux/stakefish-api/pkg/models"
)

// MockConnector is a mocked DB connector used for testing purposes.
type MockConnector struct {
	SaveQueryError                    bool
	SaveQueryErrorMsg                 string
	Queries                           []models.Query
	RetrieveLastTwentyQueriesError    bool
	RetrieveLastTwentyQueriesErrorMsg string
}

func NewMockConnector() *MockConnector {
	return &MockConnector{
		Queries: make([]models.Query, 20),
	}
}

func (m *MockConnector) Ping() error {
	return nil
}

func (m *MockConnector) SaveQuery(query models.Query) error {
	if m.SaveQueryError {
		return fmt.Errorf(m.SaveQueryErrorMsg)
	}

	m.Queries = append(m.Queries, query)
	return nil
}

func (m *MockConnector) RetrieveLastTwentyQueries() ([]models.Query, error) {
	if m.RetrieveLastTwentyQueriesError {
		return nil, fmt.Errorf(m.RetrieveLastTwentyQueriesErrorMsg)
	}

	return m.Queries, nil
}

func (m *MockConnector) Close() error {
	return nil
}
