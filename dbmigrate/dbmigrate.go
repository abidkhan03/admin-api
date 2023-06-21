package dbmigrate

import (
	"log"

	"github.com/spongeling/admin-api/shared"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Run(cfg Config) error {
	// database
	connStr := shared.GetDBConnectionString()

	log.Println("Starting database migrations...")

	// get source url
	log.Println("Migrating source: " + cfg.MigrationsSource)

	// get database url
	log.Println("Migrating target: " + connStr)

	// new migration instance
	m, err := migrate.New(cfg.MigrationsSource, connStr)
	if err != nil {
		return err
	}

	// migration to a specific version
	if cfg.MigrationVersion >= 0 {
		err = m.Migrate(uint(cfg.MigrationVersion))
	} else if cfg.Down {
		err = m.Down()
	} else {
		err = m.Up()
	}
	if err == migrate.ErrNoChange {
		log.Println("No change in database...")
		return nil
	} else if err != nil {
		return err
	}

	log.Println("Database migrations completed...")
	return nil
}
