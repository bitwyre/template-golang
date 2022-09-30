package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepo IUserRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepo: newUserRepository(db),
	}
}
