package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database *Database `validate:"required"`
}

type Database struct {
	DSN                        string `validate:"required"`
	Driver                     string `validate:"required,,oneof=postgres sqlite mysql"`
	PoolMaxOpenConns           int    `validate:"gte=0"`
	PoolMaxIdleConns           int    `validate:"gte=0"`
	PoolMaxConnLifetimeSeconds int    `validate:"gte=0"`
}

func NewConfig() (*Config, error) {

	err := godotenv.Load("./config/.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := &Config{
		Database: &Database{
			DSN:                        os.Getenv("DATABASE_DSN"),
			Driver:                     os.Getenv("DATABASE_DRIVER"),
			PoolMaxOpenConns:           10, //@TODO: Make configurable via env variable
			PoolMaxIdleConns:           5,
			PoolMaxConnLifetimeSeconds: 300,
		},
	}
	return cfg, nil
}
