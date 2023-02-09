package infra

type DBConnector interface {
	Connect() error
	MakeQuery(query string) error
	Close() error
}

type PostgresConnector struct {
	// TO-DO
}

func NewPostgresConnector() *PostgresConnector {
	return &PostgresConnector{}
}

func (p *PostgresConnector) Connect() error {
	// TO-DO
	return nil
}

func (p *PostgresConnector) MakeQuery(query string) error {
	// TO-DO
	return nil
}

func (p *PostgresConnector) Close() error {
	// TO-DO
	return nil
}
