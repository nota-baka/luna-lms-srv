package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"Luna.LMS/server/entities"
)

type TOTPsAttemptsWrapper struct {
	db *sql.DB
}

func (w *TOTPsAttemptsWrapper) Add(c context.Context, attempt *entities.TOTPAttempt) error {
	_, err := w.db.ExecContext(c, "INSERT INTO accounts_totp_attempts (id, account_id, expires_at) VALUES ($1, $2, $3)", attempt.Id, attempt.AccountId, attempt.ExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to add totp attempt: %w", err)
	}
	return nil
}

func (w *TOTPsAttemptsWrapper) GetById(c context.Context, id string) (*entities.TOTPAttempt, error) {
	var attempt entities.TOTPAttempt
	err := w.db.QueryRowContext(c, "SELECT id, account_id, expires_at").Scan(&attempt.Id, &attempt.AccountId, &attempt.ExpiresAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get totp attempt by id: %w", err)
	}
	return &attempt, nil
}

func (w *TOTPsAttemptsWrapper) Modify(c context.Context, attempt *entities.TOTPAttempt) error {
	_, err := w.db.ExecContext(c, "UPDATE accounts_totp_attempts SET expires_at = $2 WHERE id = $1", attempt.Id, attempt.ExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to moditfy totp attempt: %w", err)
	}
	return nil
}

func (w *TOTPsAttemptsWrapper) RemoveById(c context.Context, id string) error {
	_, err := w.db.ExecContext(c, "DELETE FROM accounts_totp_attempts WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to remove totp attempt by id: %w", err)
	}
	return nil
}
