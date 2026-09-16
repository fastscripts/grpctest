package models

import "time"

type User struct {
	ID        int64  `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
