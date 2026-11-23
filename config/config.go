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

type DatabaseType int

const (
	DatabaseTypePostgres DatabaseType = iota
)

func (t *DatabaseType) UnmarshalText(text []byte) error {
	textStr := string(text)
	switch textStr {
	case "postgres":
		*t = DatabaseTypePostgres
	default:
		return fmt.Errorf("unknown DatabaseType value: %s", textStr)
	}
	return nil
}

// DatabaseConfig is configuration of database
type DatabaseConfig struct {
	Type     DatabaseType           `toml:"type"`
	Postgres DatabasePostgresConfig `toml:"postgres"`
}

// ServerSSLConfig is configuration of SSL/TLS for server
type ServerSSLConfig struct {
	Enabled bool `toml:"enabled"`
}

type ServerMode int

const (
	ServerModeDebug ServerMode = iota
	ServerModeProd
)

func (m *ServerMode) UnmarshalText(text []byte) error {
	textStr := string(text)
	switch textStr {
	case "debug":
		*m = ServerModeDebug
	case "prod":
		*m = ServerModeProd
	default:
		return fmt.Errorf("unknown ServerMode value: %s", textStr)
	}
	return nil
}

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
	// TODO: add debug log
	// speaker.DebugLog(&c)
	return &c, nil
}
