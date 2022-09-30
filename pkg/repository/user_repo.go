package repository

import (
	"log"

	"github.com/bitwyre/template-golang/pkg/datastore/postgres/entity"
	"gorm.io/gorm"
)

type IUserRepo interface {
	FindById(id int) (entity.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (r *userRepo) FindById(id int) (entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepo) DeleteById(user *entity.User) error {
	err := r.db.Delete(&user).Error
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}