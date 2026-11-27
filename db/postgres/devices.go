package postgres

import (
	"context"
	"database/sql"

	"Luna.LMS/server/entities"
)

type DevicesWrapper struct {
	db *sql.DB
}

func (w *DevicesWrapper) Add(c context.Context, device *entities.Device) error {
	// TODO: implement me
	panic("implement me")
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
