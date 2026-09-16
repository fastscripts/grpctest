package seeder

import (
	"github.com/fastscripts/grpctest/internal/database/models"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {

	users := []models.User{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
		{Name: "Charlie", Email: "charlie@example.com"},
	}
	for _, u := range users {
		err := db.Create(&u).Error
		if err != nil {
			return err
		}
	}
	return nil
}
