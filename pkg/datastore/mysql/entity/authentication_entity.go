package entity

import (
	"time"

	"gorm.io/gorm"
)

type Authentication struct {
	ID                int       `gorm:"primaryKey;autoIncrement:1;not_null"`
	UserUuid          string    `gorm:"not null; size:36"`
	ExpireTime        string    `gorm:"not null; size:70"`
	TokenType         string    `gorm:"not null; size:36"`
	ExpiresIn         uint      `gorm:"type:bigint(20) unsigned not null;size:20"`
	AccessToken       string    `gorm:"not null; size:1000"`
	RefreshToken      string    `gorm:"not null; size:1000"`
	RSAPublicKey      string    `gorm:"not null; size:500"`
	IPAddress         string    `gorm:"size:25"`
	Country           string    `gorm:"size:50"`
	LastAuthenticated int       `gorm:"size:20"`
	Revoked           int       `gorm:"not null; tinyint"`
	CreatedAt         time.Time `gorm:"autoCreateTime:nano"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime:nano"`
	DeletedAt         gorm.DeletedAt
}
