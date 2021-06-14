package postgres

import (
	"errors"
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"

	"Sharykhin/rent-car/domain"
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
		return nil, domain.NewInternalError(
			errors.New("dns string is empty"),
			"[infrastructure][postgres][NewConnection]",
		)
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
		return domain.NewInternalError(
			fmt.Errorf("failed to connect to postgres: %v", err),
			"[infrastructure][postgres][Connection][Connect]",
		)

	}

	conn.DB = db

	return nil
}

// Close closes the postgres connection
func (conn *Connection) Close() error {
	err := conn.DB.Close()
	if err != nil {
		return domain.NewInternalError(
			fmt.Errorf("failed to close postgres connection: %v", err),
			"[infrastructure][postgres][Connection][Close]",
		)
	}

	return nil
}
