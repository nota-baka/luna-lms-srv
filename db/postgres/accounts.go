package postgres

import (
	"context"
	"database/sql"

	"Luna.LMS/server/entities"
)

type AccountsWrapper struct {
	db *sql.DB
}

func (w *AccountsWrapper) Add(c context.Context, acc *entities.Account) error {
	// TODO: implement me
	panic("implement me")
}

func (w *AccountsWrapper) GetById(c context.Context, id string) (*entities.Account, error) {
	// TODO: implement me
	panic("implement me")
}

func (w *AccountsWrapper) GetByIds(c context.Context, ids []string) ([]entities.Account, error) {
	// TODO: implement me
	panic("implement me")
}

func (w *AccountsWrapper) GetByDeviceId(c context.Context, id string) (*entities.Account, error) {
	// TODO: implement me
	panic("implement me")
}

func (w *AccountsWrapper) GetByEmail(c context.Context, email string) (*entities.Account, error) {
	// TODO: implement me

}

func (w *AccountsWrapper) Modify(c context.Context, acc *entities.Account) error {
	// TODO: implement me
	panic("implement me")
}

func (w *AccountsWrapper) Remove(c context.Context, acc *entities.Account) error {
	// TODO: implement me
	panic("implement me")
}
