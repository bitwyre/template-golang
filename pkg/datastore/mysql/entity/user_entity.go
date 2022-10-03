package entity

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement:1;not_null"`
	UserCode  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Status    int    `gorm:"not null"`
	CreatedAt time.Time
}
