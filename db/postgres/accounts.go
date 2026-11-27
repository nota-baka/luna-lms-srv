package postgres

import (
	"context"
	"database/sql"
	"fmt"

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
	var acc entities.Account
	err := w.db.QueryRowContext(c, "SELECT id, email_address, name, role, password_hash FROM Accounts WHERE email_address = $1", email).Scan(&acc.Id, &acc.Email, &acc.Name, &acc.Role, &acc.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get email by id: %w", err)
	}
	return &acc, nil
}

func (w *AccountsWrapper) Modify(c context.Context, acc *entities.Account) error {
	// TODO: implement me
	panic("implement me")
}

func (w *AccountsWrapper) RemoveById(c context.Context, id string) error {
	// TODO: implement me
	panic("implement me")
}
