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
	// TODO: implement me
	panic("implement me")
}

func (w *TOTPsWrapper) GetByAccountId(c context.Context, id string) (*entities.TOTP, error) {
	// TODO: implement me
	var totp entities.TOTP
	err := w.db.QueryRowContext(c, "SELECT account_id, secret_enc, last_used_step, failed_attempts, locked_until, confirmed FROM accounts_totp WHERE account_id = $1",
		id).Scan(&totp.AccountId, &totp.SecretEnc, &totp.LastUsedStep, &totp.FailedAttempts, &totp.LockedUntil, &totp.Confirmed)
	if err != nil {
		return nil, fmt.Errorf("failed to get totp by account id: %w", err)
	}
	return &totp, nil
}

func (w *TOTPsWrapper) Modify(c context.Context, totp *entities.TOTP) error {
	// TODO: implement me
	panic("implement me")
}

func (w *TOTPsWrapper) RemoveById(c context.Context, id string) error {
	// TODO: implement me
	panic("implement me")
}
