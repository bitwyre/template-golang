package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:1;not_null"`
	UserCode  string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Status    int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano"`
	DeletedAt gorm.DeletedAt
}
