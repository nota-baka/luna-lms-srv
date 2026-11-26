package config

import "fmt"

// DatabasePostgresConfig is configuration of PostgreSQL database
type DatabasePostgresConfig struct {
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	User         string `toml:"user"`
	Password     string `toml:"password"`
	DatabaseName string `toml:"database_name"`
}

type DatabaseType string

const (
	DatabaseTypeUndefined DatabaseType = "undefined"
	DatabaseTypePostgres  DatabaseType = "postgres"
)

// DatabaseConfig is configuration of database
type DatabaseConfig struct {
	Type     DatabaseType           `toml:"type"`
	Postgres DatabasePostgresConfig `toml:"postgres"`
}

// ServerSSLConfig is configuration of SSL/TLS for server
type ServerSSLConfig struct {
	Enabled bool `toml:"enabled"`
}

type ServerMode string

const (
	ServerModeUndefined ServerMode = "undefined"
	ServerModeDebug     ServerMode = "debug"
	ServerModeProd      ServerMode = "prod"
)

// ServerConfig is configuration of server
type ServerConfig struct {
	Mode ServerMode      `toml:"mode"`
	Port int             `toml:"port"`
	SSL  ServerSSLConfig `toml:"ssl"`
}

// EncryptionConfig is configuration of encryption
type EncryptionConfig struct {
	Key string `toml:"key"`
}

// Config is full application configuration
type Config struct {
	Server     ServerConfig     `toml:"server"`
	Encryption EncryptionConfig `toml:"encryption"`
	Database   DatabaseConfig   `toml:"database"`
}

// NewConfig creates, loads and returns new configuration instance
func NewConfig() (*Config, error) {
	var c Config
	if err := loadFromTOML(&c); err != nil {
		return nil, fmt.Errorf("failed to load config from TOML: %w", err)
	}
	if err := loadFromEnv(&c); err != nil {
		return nil, fmt.Errorf("failed to load config from env: %w", err)
	}
	if err := validate(&c); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}
	// TODO: add debug log
	// speaker.DebugLog(&c)
	return &c, nil
}
