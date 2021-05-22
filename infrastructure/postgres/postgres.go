package postgres

import (
	"errors"
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"
)

type (
	// Connection handles postgres connection
	Connection struct {
		dns string
		DB  *sql.DB
	}
)

// NewConnection creates a new Connection instance
func NewConnection(dns string) (*Connection, error) {
	if dns == "" {
		return nil, errors.New("[postgres] dns string is empty")
	}
	conn := Connection{
		dns: dns,
		DB:  nil,
	}

	return &conn, nil
}

// Connect connects to postgres
func (conn *Connection) Connect() error {
	db, err := sql.Open("postgres", conn.dns)
	if err != nil {
		return fmt.Errorf("[postgres][Connection] failed to connect to postgres: %v", err)
	}

	conn.DB = db

	return nil
}

// Close closes the postgres connection
func (conn *Connection) Close() error {
	err := conn.DB.Close()
	if err != nil {
		return fmt.Errorf("[postgres][Connection] failed to close postgres: %v", err)
	}

	return nil
}
