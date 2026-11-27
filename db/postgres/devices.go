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

// TODO: rename device to dev
func (w *DevicesWrapper) Add(c context.Context, device *entities.Device) error {
	_, err := w.db.ExecContext(c, "INSERT INTO devices (id, account_id, last_online_at, session_token_hash, session_expires_at) VALUES ($1, $2, $3, $4, $5)",
		device.Id, device.AccountId, device.LastOnlineAt, device.TokenHash, device.TokenExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to add device: %w", err)
	}
	return nil
}

func (w *DevicesWrapper) GetById(c context.Context, id string) (*entities.Device, error) {
	var dev entities.Device
	err := w.db.QueryRowContext(c, "SELECT id, account_id, last_online_at, session_token_hash, session_expires_at FROM devices WHERE id = $1", id).Scan(&dev.Id, &dev.AccountId, &dev.LastOnlineAt, &dev.TokenHash, &dev.TokenExpiresAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get device by id: %w", err)
	}
	return &dev, nil
}

func (w *DevicesWrapper) GetByAccountId(c context.Context, id string) ([]entities.Device, error) {
	// TODO: implement me
	var devs []entities.Device
	rows, err := w.db.QueryContext(c, "SELECT d.id, d.account_id, d.last_online_at, d.session_token_hash, d.session_expires_at FROM devices AS d JOIN accounts AS a ON a.id = d.account_id WHERE a.id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to get devices by account id: %w", err)
	}
	for rows.Next() {
		var dev entities.Device
		err := rows.Scan(&dev.Id, &dev.AccountId, &dev.LastOnlineAt, &dev.TokenHash, &dev.TokenExpiresAt)
		if err != nil {
			return nil, fmt.Errorf("failed to read device from row: %w", err)
		}
		devs = append(devs, dev)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to prepare device row to read: %w", err)
	}
	return devs, nil
}

// TODO: rename device to dev
func (w *DevicesWrapper) Modify(c context.Context, device *entities.Device) error {
	_, err := w.db.ExecContext(c, "UPDATE devices SET last_online_at = $2, session_token_hash = $3, session_expires_at = $4 WHERE id = $1", device.Id, device.LastOnlineAt, device.TokenHash, device.TokenExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to modify device: %w", err)
	}
	return nil
}

func (w *DevicesWrapper) RemoveById(c context.Context, id string) error {
	_, err := w.db.ExecContext(c, "DELETE FROM devices WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to remove device by id: %w", err)
	}
	return nil
}
