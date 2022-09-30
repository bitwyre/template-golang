package seeder

import (
	"log"

	"github.com/bitwyre/template-golang/pkg/datastore/postgres/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

func UserDataSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db}
}

func (s *Seeder) CreateOne() *entity.User {
	user := entity.User{
		UserCode: uuid.New().String(),
		Email:    faker.Internet().Email(),
		Status:   1,
	}

	err := s.db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	return &user
}

func (s *Seeder) CreateMany(count int) []*entity.User {
	var users []*entity.User

	for i := 1; i <= count; i++ {
		user := s.CreateOne()
		users = append(users, user)
	}

	return users
}
