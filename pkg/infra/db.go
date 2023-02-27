package infra

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/mikeletux/stakefish-api/pkg/models"
)

// DBConnector is the interface that must be implemented to access the storage backend.
type DBConnector interface {
	Ping() error
	SaveQuery(query models.Query) error
	RetrieveLastTwentyQueries() ([]models.Query, error)
	Close() error
}

// PostgresConnector is a struct that implements DBConnector and gives access to a PostgreSQL storage backed.
type PostgresConnector struct {
	db *pg.DB
}

// NewPostgresConnector returns a new PostgresConnector given an address, user, password and database.
func NewPostgresConnector(addr, user, pass, database string) *PostgresConnector {
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: pass,
		Database: database,
	})

	return &PostgresConnector{db: db}
}

// Ping checks if the database connection is healthy. Return an error otherwise.
func (p *PostgresConnector) Ping() error {
	ctx := context.Background() // We should be handling here timeouts.
	return p.db.Ping(ctx)
}

// SaveQuery inserts in the database backend the relational representation of models.Query struct.
func (p *PostgresConnector) SaveQuery(query models.Query) error {
	_, err := p.db.Model(&query).Insert()
	if err != nil {
		return err
	}

	for _, ip := range query.Addresses {
		ip.QueryID = query.Id
		_, err := p.db.Model(&ip).Insert()
		if err != nil {
			return err
		}
	}

	return nil
}

// RetrieveLastTwentyQueries returns the last 20 queries stored in the database in a descendent manner.
func (p *PostgresConnector) RetrieveLastTwentyQueries() ([]models.Query, error) {
	queries := make([]models.Query, 0, 20)
	err := p.db.Model(&queries).
		Relation("Addresses").
		Order("created_at DESC").
		Limit(20).
		Select()
	if err != nil {
		return nil, err
	}
	return queries, nil
}

// Close closes the connection to the database.
func (p *PostgresConnector) Close() error {
	return p.db.Close()
}
