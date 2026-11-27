package postgres

import (
	"context"
	"database/sql"

	"Luna.LMS/server/entities"
)

type TOTPsAttemptsWrapper struct {
	db *sql.DB
}

func (w *TOTPsAttemptsWrapper) Add(c context.Context, attempt *entities.TOTPAttempt) error {
	// TODO: implement me
	panic("implement me")
}

func (w *TOTPsAttemptsWrapper) GetById(c context.Context, id string) (*entities.TOTPAttempt, error) {
	// TODO: implement me
	panic("implement me")
}

func (w *TOTPsAttemptsWrapper) Modify(c context.Context, attempt *entities.TOTPAttempt) error {
	// TODO: implement me
	panic("implement me")
}

func (w *TOTPsAttemptsWrapper) RemoveById(c context.Context, id string) error {
	// TODO: implement me
	panic("implement me")
}
