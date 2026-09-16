package database

import (
	"context"
	"testing"

	"github.com/fastscripts/grpctest/internal/config"
)

func Test_Database_SupportedDriver(t *testing.T) {

	db, err := NewDatabaseService(&config.Config{
		Database: &config.Database{
			DSN:    ":memory:",
			Driver: "sqlite",
		},
	})
	if err != nil {
		t.Fatalf("expected no error for supported driver, got %v", err)
	}
	if db == nil {
		t.Fatalf("expected non-nil database for supported driver, got nil")
	}

	gormDB := db.DB()
	if gormDB == nil {
		t.Fatalf("expected non-nil gorm.DB for supported driver, got nil")
	}

	err = db.Ping(context.Background())
	if err != nil {
		t.Fatalf("expected no error on ping for supported driver, got %v", err)
	}

	err = db.Close()
	if err != nil {
		t.Fatalf("expected no error on close for supported driver, got %v", err)
	}

}

func Test_Database_UnsupportedDriver(t *testing.T) {

	db, err := NewDatabaseService(&config.Config{
		Database: &config.Database{
			DSN:    "test",
			Driver: "unsupported_driver",
		},
	})
	if err == nil {
		t.Fatalf("expected error for unsupported driver, got nil")
	}
	if db != nil {
		t.Fatalf("expected nil database for unsupported driver, got %v", db)
	}

}

func Test_Database_InvalidDSN(t *testing.T) {

	// For sqlite, an empty string is invalid
	config := &config.Config{
		Database: &config.Database{
			DSN:    "",
			Driver: "sqlite",
		},
	}
	db, err := NewDatabaseService(config)
	if err == nil {
		t.Fatalf("expected error for invalid DSN, got nil")
	}
	if db != nil {
		t.Fatalf("expected nil database for invalid DSN, got %v", db)
	}

}

func Test_Database_InvalidDB(t *testing.T) {
	db := &Database{db: nil}
	err := db.Close()
	if err == nil {
		t.Fatalf("expected error for closing uninitialized database, got nil")
	}
}
