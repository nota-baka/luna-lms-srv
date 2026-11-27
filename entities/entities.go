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
	Id           string
	Role         AccountRole
	Name         string
	Email        string
	PasswordHash []byte
}

type Device struct {
	Id             string
	AccountId      string
	LastOnlineAt   time.Time
	TokenHash      []byte
	TokenExpiresAt time.Time
}

type TOTP struct {
	AccountId      string
	SecretEnc      string
	LastUsedStep   int64
	FailedAttempts int
	LockedUntil    time.Time
	Confirmed      bool
}

type TOTPAttempt struct {
	Id        string
	AccountId string
	ExpiresAt time.Time
}
