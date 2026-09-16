package service

import (
	"context"

	"github.com/fastscripts/grpctest/internal/database"
	"github.com/fastscripts/grpctest/internal/database/models"
	"gorm.io/gorm"
)

type UserService interface {
	FindByID(ctx context.Context, id uint) (*models.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db database.DatabaseService) UserService {
	return &userService{db: db.DB()}
}

func (s *userService) FindByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	if err := s.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
