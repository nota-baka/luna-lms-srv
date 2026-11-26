package db

import (
	"context"

	"Luna.LMS/server/entities"
)

type Accounts interface {
	Add(c context.Context, acc *entities.Account) error
	GetById(c context.Context, id string) (*entities.Account, error)
	GetByIds(c context.Context, ids []string) ([]entities.Account, error)
	GetByDeviceId(c context.Context, id string) (*entities.Account, error)
	GetByEmail(c context.Context, id string) (*entities.Account, error)
	Modify(c context.Context, acc *entities.Account) error
	Remove(c context.Context, acc *entities.Account) error
}

type TOTPs interface {
	Add(c context.Context, totp *entities.TOTP) error
	GetByAccountId(c context.Context, id string) (*entities.TOTP, error)
	Modify(c context.Context, totp *entities.TOTP) error
	Remove(c context.Context, totp *entities.TOTP) error
}

type Devices interface {
	Add(c context.Context, device *entities.Device) error
	GetById(c context.Context, id string) (*entities.Device, error)
	GetByAccountId(c context.Context, id string) ([]entities.Device, error)
	Modify(c context.Context, device *entities.Device) error
	Remove(c context.Context, device *entities.Device) error
}

type DB interface {
	Accounts() Accounts
	TOTPs() TOTPs
	Devices() Devices
	Close() error
}
