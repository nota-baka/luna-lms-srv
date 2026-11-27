package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"Luna.LMS/server/entities"
)

type TOTPsWrapper struct {
	db *sql.DB
}

func (w *TOTPsWrapper) Add(c context.Context, totp *entities.TOTP) error {
	_, err := w.db.ExecContext(c, "INSERT INTO accounts_totp (account_id, secret_enc, last_used_step, failed_attempts, locked_until, confirmed) VALUES ($1, $2, $3, $4, $5, $6)", totp.AccountId, totp.SecretEnc, totp.LastUsedStep, totp.FailedAttempts, totp.LockedUntil, totp.Confirmed)
	if err != nil {
		return fmt.Errorf("failed to add totp: %w", err)
	}
	return nil
}

func (w *TOTPsWrapper) GetByAccountId(c context.Context, id string) (*entities.TOTP, error) {
	var totp entities.TOTP
	err := w.db.QueryRowContext(c, "SELECT account_id, secret_enc, last_used_step, failed_attempts, locked_until, confirmed FROM accounts_totp WHERE account_id = $1",
		id).Scan(&totp.AccountId, &totp.SecretEnc, &totp.LastUsedStep, &totp.FailedAttempts, &totp.LockedUntil, &totp.Confirmed)
	if err != nil {
		return nil, fmt.Errorf("failed to get totp by account id: %w", err)
	}
	return &totp, nil
}

func (w *TOTPsWrapper) Modify(c context.Context, totp *entities.TOTP) error {
	_, err := w.db.ExecContext(c, "UPDATE accounts_totp SET last_used_step = $2, failed_attempts = $3, locked_until = $4, confirmed = $5 WHERE account_id = $1",
		totp.AccountId, totp.LastUsedStep, totp.FailedAttempts, totp.LockedUntil, totp.Confirmed)
	if err != nil {
		return fmt.Errorf("failed to modify totp: %w", err)
	}
	return nil
}

func (w *TOTPsWrapper) RemoveById(c context.Context, id string) error {
	_, err := w.db.ExecContext(c, "DELETE FROM accounts_totp WHERE account_id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to remove totp by account id: %w", err)
	}
	return nil
}
