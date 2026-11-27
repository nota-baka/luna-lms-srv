package db

import (
	"context"

	"Luna.LMS/server/entities"
)

type Accounts interface {
	Add(c context.Context, acc *entities.Account) error
	GetById(c context.Context, id string) (*entities.Account, error)
	GetByDeviceId(c context.Context, id string) (*entities.Account, error)
	GetByEmail(c context.Context, email string) (*entities.Account, error)
	Modify(c context.Context, acc *entities.Account) error
	RemoveById(c context.Context, id string) error
}

type TOTPs interface {
	Add(c context.Context, totp *entities.TOTP) error
	GetByAccountId(c context.Context, id string) (*entities.TOTP, error)
	Modify(c context.Context, totp *entities.TOTP) error
	RemoveById(c context.Context, id string) error
}

type TOTPsAttempts interface {
	Add(c context.Context, attempt *entities.TOTPAttempt) error
	GetById(c context.Context, id string) (*entities.TOTPAttempt, error)
	Modify(c context.Context, attempt *entities.TOTPAttempt) error
	RemoveById(c context.Context, id string) error
}

type Devices interface {
	Add(c context.Context, device *entities.Device) error
	GetById(c context.Context, id string) (*entities.Device, error)
	GetByAccountId(c context.Context, id string) ([]entities.Device, error)
	Modify(c context.Context, device *entities.Device) error
	RemoveById(c context.Context, ids string) error
}

type DB interface {
	Accounts() Accounts
	TOTPs() TOTPs
	TOTPsAttempts() TOTPsAttempts
	Devices() Devices
	Close() error
}
