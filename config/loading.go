package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	gap "github.com/muesli/go-app-paths"
)

// loadFromTOML loads configuration from TOML files
func loadFromTOML(c *Config) error {
	scope := gap.NewScope(gap.User, "luna-lms-srv")
	tomlFiles, err := scope.LookupConfig("config.toml")
	if err != nil {
		return fmt.Errorf("failed to lookup TOML config files: %w", err)
	}
	for _, tomlFile := range tomlFiles {
		f, err := os.ReadFile(tomlFile)
		if err != nil {
			return fmt.Errorf("failed to read TOML file: %w", err)
		}
		err = toml.Unmarshal(f, c)
		if err != nil {
			return fmt.Errorf("failed to unmarshal TOML: %w", err)
		}
	}
	return nil
}

// loadFromEnv loads configuration from environment variables
func loadFromEnv(c *Config) error {
	if serverMode := os.Getenv("LUNASRV_SERVER_MODE"); serverMode != "" {
		c.Server.Mode = ServerMode(serverMode)
	}
	if serverPort := os.Getenv("LUNASRV_SERVER_PORT"); serverPort != "" {
		if p, err := strconv.Atoi(serverPort); err != nil {
			return fmt.Errorf("failed to convert server port to int: %w", err)
		} else {
			c.Server.Port = p
		}
	}
	if serverSSLEnabled := os.Getenv("LUNASRV_SERVER_SSL_ENABLED"); serverSSLEnabled != "" {
		c.Server.SSL.Enabled = true
	}
	if encryptionKey := os.Getenv("LUNASRV_ENCRYPTION_KEY"); encryptionKey != "" {
		c.Encryption.Key = encryptionKey
	}
	if databaseType := os.Getenv("LUNASRV_DATABASE_TYPE"); databaseType != "" {
		c.Database.Type = DatabaseType(databaseType)
	}
	if databasePostgresHost := os.Getenv("LUNASRV_DATABASE_POSTGRES_HOST"); databasePostgresHost != "" {
		c.Database.Postgres.Host = databasePostgresHost
	}
	if databasePostgresPort := os.Getenv("LUNASRV_DATABASE_POSTGRES_PORT"); databasePostgresPort != "" {
		if p, err := strconv.Atoi(databasePostgresPort); err != nil {
			return fmt.Errorf("failed to convert postgres port to int: %w", err)
		} else {
			c.Database.Postgres.Port = p
		}
	}
	if databasePostgresUser := os.Getenv("LUNASRV_DATABASE_POSTGRES_USER"); databasePostgresUser != "" {
		c.Database.Postgres.User = databasePostgresUser
	}
	if databasePostgresPassword := os.Getenv("LUNASRV_DATABASE_POSTGRES_PASSWORD"); databasePostgresPassword != "" {
		c.Database.Postgres.Password = databasePostgresPassword
	}
	if databasePostgresDatabaseName := os.Getenv("LUNASRV_DATABASE_POSTGRES_DATABASE_NAME"); databasePostgresDatabaseName != "" {
		c.Database.Postgres.DatabaseName = databasePostgresDatabaseName
	}
	return nil
}
