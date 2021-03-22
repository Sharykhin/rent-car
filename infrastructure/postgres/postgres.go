package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Connect(postgresURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %v", err)
	}

	return db, nil
}
