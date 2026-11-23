package config

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

// ServerConfig is configuration of server
type ServerConfig struct {
	Mode ServerMode      `toml:"mode"`
	Port int             `toml:"port"`
	SSL  ServerSSLConfig `toml:"ssl"`
}

// EncryptionConfig is configuration of encryption
type EncryptionConfig struct {
	Key []byte `toml:"key"`
}

// Config is full application configuration
type Config struct {
	Server     ServerConfig     `toml:"server"`
	Encryption EncryptionConfig `toml:"encryption"`
	Database   DatabaseConfig   `toml:"database"`
}
