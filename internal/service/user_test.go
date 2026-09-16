package service

import (
	"context"
	"testing"

	"github.com/fastscripts/grpctest/internal/config"
	"github.com/fastscripts/grpctest/internal/database"
	"github.com/fastscripts/grpctest/internal/database/models"
)

func Test_UserService_FindByID(t *testing.T) {

	db, err := database.NewDatabaseService(&config.Config{
		Database: &config.Database{
			DSN:    ":memory:",
			Driver: "sqlite",
		},
	})
	if err != nil {
		t.Fatalf("failed to create database service: %v", err)
	}
	defer db.Close()

	u := &models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	// Insert the user into the database
	if err := db.DB().Create(u).Error; err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Now test the FindByID method

	userService := NewUserService(db)
	retrievedUser, err := userService.FindByID(context.Background(), u.ID)
	if err != nil {
		t.Fatalf("FindByID returned an error: %v", err)
	}

	if retrievedUser == nil {
		t.Fatalf("FindByID returned nil, expected a user")
	}

	if retrievedUser.ID != u.ID || retrievedUser.Name != u.Name || retrievedUser.Email != u.Email {
		t.Fatalf("retrieved user does not match created user. Got %+v, expected %+v", retrievedUser, u)
	}
}
