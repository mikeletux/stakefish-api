package infra

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/mikeletux/stakefish-api/pkg/models"
)

type DBConnector interface {
	Ping() error
	SaveQuery(query models.Query) error
	RetrieveLastTwentyQueries() ([]models.Query, error)
	Close() error
}

type PostgresConnector struct {
	db *pg.DB
}

func NewPostgresConnector() *PostgresConnector {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "postgres",
		Database: "stakefish",
	})

	return &PostgresConnector{db: db}
}

func (p *PostgresConnector) Ping() error {
	ctx := context.Background() // We should be handling here timeouts.
	return p.db.Ping(ctx)
}

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

func (p *PostgresConnector) RetrieveLastTwentyQueries() ([]models.Query, error) {
	var queries []models.Query
	err := p.db.Model(&queries).
		Relation("Addresses").
		Select()
	if err != nil {
		return nil, err
	}
	return queries, nil
}

func (p *PostgresConnector) Close() error {
	return p.db.Close()
}
