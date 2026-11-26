package main

import (
	"fmt"
	"log"

	"Luna.LMS/server/config"
	"Luna.LMS/server/db"
	"Luna.LMS/server/db/postgres"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	dbw, err := prepareDB(&conf.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := dbw.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

func prepareDB(c *config.DatabaseConfig) (db.DB, error) {
	switch c.Type {
	case config.DatabaseTypePostgres:
		postgres, err := postgres.NewPostgresWrapper(&c.Postgres)
		if err != nil {
			return nil, fmt.Errorf("failed to prepare database handler: %w", err)
		}
		return postgres, err
	default:
		return nil, fmt.Errorf("failed to prepare database handler: unhandled database type %q", c.Type)
	}
}
