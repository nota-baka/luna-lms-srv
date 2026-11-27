package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"Luna.LMS/server/entities"
)

type DevicesWrapper struct {
	db *sql.DB
}

func (w *DevicesWrapper) Add(c context.Context, device *entities.Device) error {
	_, err := w.db.ExecContext(c, "INSERT INTO Devices (id, account_id, last_online_at, session_token_hash, session_expires_at) VALUES ($1, $2, $3, $4, $5)",
		device.Id, device.AccountId, device.LastOnlineAt, device.TokenHash, device.TokenExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to add device: %w", err)
	}
	return nil
}

func (w *DevicesWrapper) GetById(c context.Context, id string) (*entities.Device, error) {
	// TODO: implement me
	panic("implement me")
}

func (w *DevicesWrapper) GetByAccountId(c context.Context, id string) ([]entities.Device, error) {
	// TODO: implement me
	panic("implement me")
}

func (w *DevicesWrapper) Modify(c context.Context, device *entities.Device) error {
	// TODO: implement me
	panic("implement me")
}

func (w *DevicesWrapper) RemoveById(c context.Context, id string) error {
	// TODO: implement me
	panic("implement me")
}
