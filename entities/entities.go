package entities

import "time"

type AccountRole string

const (
	AccountRoleUndefined     AccountRole = "undefined"
	AccountRoleStudent       AccountRole = "student"
	AccountRoleCourseManager AccountRole = "course_manager"
	AccountRoleModerator     AccountRole = "moderator"
	AccountRoleAdministrator AccountRole = "administrator"
)

type Account struct {
	id           string
	role         AccountRole
	name         string
	email        string
	passwordHash []byte
}

type Device struct {
	id             string
	accountId      string
	lastOnlineAt   time.Time
	tokenHash      []byte
	tokenExpiresAt time.Time
}

type TOTP struct {
	accountId      string
	secretEnc      string
	lastUsedStep   int64
	failedAttempts int
	lockedUntil    time.Time
	confirmed      bool
}
