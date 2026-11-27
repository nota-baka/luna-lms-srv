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
	_, err := w.db.ExecContext(c, "INSERT INTO Accounts (id, role, name, email_address, password_hash) VALUES ($1, $2, $3, $4, $5)",
		acc.Id, acc.Role, acc.Name, acc.Email, acc.PasswordHash)
	if err != nil {
		return fmt.Errorf("failed to add account: %w", err)
	}
	return nil
}

func (w *AccountsWrapper) GetById(c context.Context, id string) (*entities.Account, error) {
	var acc entities.Account
	err := w.db.QueryRowContext(c, "SELECT id, role, name, email_address, password_hash FROM Accounts WHERE id = $1", id).Scan(&acc.Id, &acc.Email, &acc.Name, &acc.Role, &acc.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by id: %w", err)
	}
	return &acc, nil
}

func (w *AccountsWrapper) GetByDeviceId(c context.Context, id string) (*entities.Account, error) {
	var acc entities.Account
	err := w.db.QueryRowContext(c, "SELECT A.id, A.role, A.name, A.email_address, A.password_hash FROM Accounts AS A JOIN Devices AS B ON A.id = B.account_id WHERE B.id = $1", id).Scan(&acc.Id, &acc.Email, &acc.Name, &acc.Role, &acc.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by device id: %w", err)
	}
	return &acc, nil
}

func (w *AccountsWrapper) GetByEmail(c context.Context, email string) (*entities.Account, error) {
	var acc entities.Account
	err := w.db.QueryRowContext(c, "SELECT id, role, name, email_address, password_hash FROM Accounts WHERE email_address = $1", email).Scan(&acc.Id, &acc.Email, &acc.Name, &acc.Role, &acc.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by email: %w", err)
	}
	return &acc, nil
}

func (w *AccountsWrapper) Modify(c context.Context, acc *entities.Account) error {
	// TODO: implement me
	_, err := w.db.ExecContext(c, "UPDATE Accounts SET role = $2, name = $3, email_address = $4, password_hash = $5 WHERE id = $1",
		acc.Id, acc.Role, acc.Email, acc.PasswordHash)
	if err != nil {
		return fmt.Errorf("failed to modify account: %w", err)
	}
	return nil
}

func (w *AccountsWrapper) RemoveById(c context.Context, id string) error {
	_, err := w.db.ExecContext(c, "DELETE FROM Accounts WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to remove account by id: %w", err)
	}
	return nil
}
