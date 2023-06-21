package main

import (
	"log"

	"github.com/spongeling/admin-api/dbmigrate"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spongeling/admin-api/shared"
)

func main() {
	// config
	err := shared.LoadConfig(".env")
	if err != nil {
		log.Fatalf("error loading config %v", err)
	}
	cfg, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = dbmigrate.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func readConfig() (dbmigrate.Config, error) {
	var cfg dbmigrate.Config

	cfg.UpdateFromArguments()
	cfg.UpdateFromEnv()

	return cfg, cfg.Validate()
}
