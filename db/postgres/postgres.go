package postgres

import (
	"database/sql"
	"fmt"

	"Luna.LMS/server/config"
	"Luna.LMS/server/db"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresWrapper struct {
	db              *sql.DB
	accountsWrapper *AccountsWrapper
	totpsWrapper    *TOTPsWrapper
	devicesWrapper  *DevicesWrapper
}

func (w *PostgresWrapper) Close() error {
	return w.db.Close()
}

func (w *PostgresWrapper) Accounts() db.Accounts {
	return w.accountsWrapper
}

func (w *PostgresWrapper) TOTPs() db.TOTPs {
	return w.totpsWrapper
}

func (w *PostgresWrapper) Devices() db.Devices {
	return w.devicesWrapper
}

// NewPostgresWrapper creates new PostgreSQL handler based on configuration
func NewPostgresWrapper(c *config.DatabasePostgresConfig) (*PostgresWrapper, error) {
	connURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DatabaseName,
	)
	db, err := sql.Open("pgx", connURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to postgresql database: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pining postgresql database: %w", err)
	}

	return &PostgresWrapper{
		db:              db,
		accountsWrapper: &AccountsWrapper{db: db},
		totpsWrapper:    &TOTPsWrapper{db: db},
		devicesWrapper:  &DevicesWrapper{db: db},
	}, nil
}
