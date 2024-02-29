package configs

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	User     string `env:"USER"`
	Name     string `env:"DB_NAME"`
	Password string `env:"PASSWORD"`
}

func ConnectToDatabase(config DatabaseConfig) (*sql.DB, error) {
	url := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Name, config.Password)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err = runMigrations(db); err != nil {
		return nil, err
	}

	return db, nil
}

func runMigrations(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	goose.SetTableName("_migrations")
	return goose.Up(db, "scheme/migrations")
}
