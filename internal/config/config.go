package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database *Database `validate:"required"`
	Redis    *Redis    `validate:"required"`
}

type Database struct {
	DSN                        string `validate:"required"`
	Driver                     string `validate:"required,,oneof=postgres sqlite mysql"`
	PoolMaxOpenConns           int    `validate:"gte=0"`
	PoolMaxIdleConns           int    `validate:"gte=0"`
	PoolMaxConnLifetimeSeconds int    `validate:"gte=0"`
}

type Redis struct {
	Addr         string `validate:"required"`
	Password     string `validate:"required"`
	DB           int    `validate:"gte=0"`
	DialTimeout  int    `validate:"gte=0"`
	ReadTimeout  int    `validate:"gte=0"`
	WithTimeout  int    `validate:"gte=0"`
	PoolSize     int    `validate:"gte=0"`
	MinIdleConns int    `validate:"gte=0"`
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
			PoolMaxOpenConns:           10,
			PoolMaxIdleConns:           5,
			PoolMaxConnLifetimeSeconds: 300,
		},
		Redis: &Redis{
			Addr:         getEnvString("REDIS_ADDR", "localhost:6379"),
			Password:     getEnvString("REDIS_PASSWORD", ""),
			DB:           getEnvInt("REDIS_DB", 0),
			DialTimeout:  getEnvInt("REDIS_DIAL_TIMEOUT", 5),
			ReadTimeout:  getEnvInt("REDIS_READ_TIMEOUT", 5),
			WithTimeout:  getEnvInt("REDIS_WITH_TIMEOUT", 5),
			PoolSize:     getEnvInt("REDIS_POOL_SIZE", 10),
			MinIdleConns: getEnvInt("REDIS_MIN_IDLE_CONNS", 5),
		},
	}
	return cfg, nil
}

func getEnvString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return parsed
}
