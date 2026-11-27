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
	_, err := w.db.ExecContext(c, "INSERT INTO accounts (id, role, name, email_address, password_hash) VALUES ($1, $2, $3, $4, $5)",
		acc.Id, acc.Role, acc.Name, acc.Email, acc.PasswordHash)
	if err != nil {
		return fmt.Errorf("failed to add account: %w", err)
	}
	return nil
}

func (w *AccountsWrapper) GetById(c context.Context, id string) (*entities.Account, error) {
	var acc entities.Account
	err := w.db.QueryRowContext(c, "SELECT id, role, name, email_address, password_hash FROM accounts WHERE id = $1", id).Scan(&acc.Id, &acc.Email, &acc.Name, &acc.Role, &acc.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by id: %w", err)
	}
	return &acc, nil
}

func (w *AccountsWrapper) GetByDeviceId(c context.Context, id string) (*entities.Account, error) {
	var acc entities.Account
	err := w.db.QueryRowContext(c, "SELECT a.id, a.role, a.name, a.email_address, a.password_hash FROM accounts AS a JOIN devices AS b ON a.id = b.account_id WHERE b.id = $1", id).Scan(&acc.Id, &acc.Email, &acc.Name, &acc.Role, &acc.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by device id: %w", err)
	}
	return &acc, nil
}

func (w *AccountsWrapper) GetByEmail(c context.Context, email string) (*entities.Account, error) {
	var acc entities.Account
	err := w.db.QueryRowContext(c, "SELECT id, role, name, email_address, password_hash FROM accounts WHERE email_address = $1", email).Scan(&acc.Id, &acc.Email, &acc.Name, &acc.Role, &acc.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by email: %w", err)
	}
	return &acc, nil
}

func (w *AccountsWrapper) Modify(c context.Context, acc *entities.Account) error {
	_, err := w.db.ExecContext(c, "UPDATE accounts SET role = $2, name = $3, email_address = $4, password_hash = $5 WHERE id = $1",
		acc.Id, acc.Role, acc.Email, acc.PasswordHash)
	if err != nil {
		return fmt.Errorf("failed to modify account: %w", err)
	}
	return nil
}

func (w *AccountsWrapper) RemoveById(c context.Context, id string) error {
	_, err := w.db.ExecContext(c, "DELETE FROM accounts WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to remove account by id: %w", err)
	}
	return nil
}
