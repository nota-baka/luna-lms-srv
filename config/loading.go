package config

import (
	"fmt"
	"os"

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
