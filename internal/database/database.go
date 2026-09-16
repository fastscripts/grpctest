package database

import (
	"context"
	"fmt"
	"time"

	"github.com/fastscripts/grpctest/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//https://dev.to/sagarmaheshwary/go-microservices-boilerplate-series-from-hello-world-to-production-part-2-428b

type Database struct {
	db *gorm.DB
}

type DatabaseService interface {
	DB() *gorm.DB
	Ping(ctx context.Context) error
	Close() error
}

func NewDatabaseService(cfg *config.Config) (DatabaseService, error) {
	var db *gorm.DB
	var err error

	switch cfg.Database.Driver {
	case "postgres":
		db, err = gorm.Open(postgres.Open(cfg.Database.DSN), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.Database.DSN), &gorm.Config{})
	case "mysql":
		db, err = gorm.Open(mysql.Open(cfg.Database.DSN), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Database.Driver)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get db instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(cfg.Database.PoolMaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.PoolMaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.PoolMaxConnLifetimeSeconds) * time.Second)

	//@todo logger verwenden
	fmt.Printf("Database connected: %s\n", cfg.Database.Driver)

	return &Database{db: db}, nil
}

func (d *Database) DB() *gorm.DB {
	return d.db
}

func (d *Database) Ping(ctx context.Context) error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get db instance: %v", err)
	}
	return sqlDB.PingContext(ctx)
}

func (d *Database) Close() error {
	if d == nil || d.db == nil {
		return fmt.Errorf("cannot close: database is not initialized")
	}
	sqlDB, err := d.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get db instance: %v", err)
	}
	return sqlDB.Close()
}
