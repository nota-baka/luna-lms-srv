package config

import (
	"encoding/base64"
	"errors"
)

// validate validates configuration values
func validate(c *Config) error {
	if c.Server.Mode == ServerModeUndefined {
		return errors.New("server mode is required")
	}
	if c.Server.Port < 1 || c.Server.Port > 65535 {
		return errors.New("server port must be in range [1, 65535]")
	}
	if keyBytes, err := base64.StdEncoding.DecodeString(c.Encryption.Key); err != nil || len(keyBytes) != 32 {
		return errors.New("encryption key must be 32 byte long key formatted to base64 string")
	}
	if c.Database.Type == DatabaseTypeUndefined {
		return errors.New("database type is required")
	}
	if c.Database.Type == DatabaseTypePostgres {
		if c.Database.Postgres.Host == "" {
			return errors.New("postgres host is required")
		}
		if c.Database.Postgres.Port < 1 || c.Database.Postgres.Port > 65535 {
			return errors.New("postgres port must be in range [1. 65535]")
		}
		if c.Database.Postgres.User == "" {
			return errors.New("postgres user is required")
		}
		if c.Database.Postgres.Password == "" {
			return errors.New("postgres password is required")
		}
		if c.Database.Postgres.DatabaseName == "" {
			return errors.New("postgres database name is required")
		}
	}
	return nil
}
