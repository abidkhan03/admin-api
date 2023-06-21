package dbmigrate

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spongeling/admin-api/internal/errors"
)

type Config struct {
	MigrationsSource string `yaml:"source"`
	MigrationVersion int    `yaml:"version"`
	Down             bool   `yaml:"down"`
}

func (c *Config) UpdateFromArguments() {
	flag.StringVar(&c.MigrationsSource, "source", "", "Migrations source")
	flag.IntVar(&c.MigrationVersion, "n", -1, "Migrate to a specific version (optional)")
	flag.BoolVar(&c.Down, "down", false, "Migrate to the last version (optional)")
	flag.Parse()

	remainingArgs := flag.Args()
	if len(remainingArgs) > 0 {
		fmt.Println("Unknown arguments:", strings.Join(remainingArgs, " "))
		os.Exit(-1)
	}
}

func (c *Config) UpdateFromEnv() {
	if x := os.Getenv("DB_MIGRATIONS_SOURCE"); x != "" {
		c.MigrationsSource = x
	}
	if x, err := strconv.ParseInt(os.Getenv("DB_MIGRATION_VERSION"), 10, 64); err == nil {
		c.MigrationVersion = int(x)
	}
}

func (c *Config) Validate() error {
	if c.MigrationsSource == "" {
		return errors.New(errors.InvalidArgument, "empty migration source. -source flag required")
	}

	return nil
}
