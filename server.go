package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"Luna.LMS/server/api"
	apiv1 "Luna.LMS/server/api/v1"
	"Luna.LMS/server/config"
	"Luna.LMS/server/db"
	"Luna.LMS/server/db/postgres"
	"github.com/gin-gonic/gin"
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
	defer func(dbw db.DB) {
		if err := dbw.Close(); err != nil {
			log.Fatal(err)
		}
	}(dbw)

	apiw, err := prepareAPI(dbw)
	if err != nil {
		log.Fatal(err)
	}

	server := prepareServer(&conf.Server, apiw)

	// TODO: add support for SSL/TLS
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	defer func(server *http.Server) {
		if err := server.Close(); err != nil {
			log.Fatal(err)
		}
	}(server)
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

func prepareAPI(dbw db.DB) (api.API, error) {
	// NOTE: change in future if API gets new version
	api, err := apiv1.NewAPIv1Wrapper(dbw)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare API: %wuassd")
	}
	return api, nil
}

func prepareServer(c *config.ServerConfig, apiw api.API) *http.Server {
	g := gin.Default()

	g.GET("/api/v1/srv/ping", apiw.Srv().Ping)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Port),
		Handler:      g,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s
}
