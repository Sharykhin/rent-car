package postgres

import (
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"
)

// Connect connects to postgres sql
func Connect(postgresURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %v", err)
	}

	return db, nil
}
