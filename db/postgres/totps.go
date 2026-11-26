package postgres

import (
	"context"
	"database/sql"

	"Luna.LMS/server/entities"
)

type TOTPsWrapper struct {
	db *sql.DB
}

func (w *TOTPsWrapper) Add(c *context.Context, totp *entities.TOTP) error {
	// TODO: implement me
	panic("implement me")
}

func (w *TOTPsWrapper) GetByAccountId(c *context.Context, id string) (*entities.TOTP, error) {
	// TODO: implement me
	panic("implement me")
}

func (w *TOTPsWrapper) Modify(c *context.Context, totp *entities.TOTP) error {
	// TODO: implement me
	panic("implement me")
}

func (w *TOTPsWrapper) Remove(c *context.Context, totp *entities.TOTP) error {
	// TODO: implement me
	panic("implement me")
}
