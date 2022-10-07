package seeder

import (
	"log"

	"github.com/bitwyre/template-golang/pkg/datastore/mysql/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

type userDataSeeder struct{ *ClientSeeder }

func UserDataSeeder(db *gorm.DB) *userDataSeeder {
	return &userDataSeeder{&ClientSeeder{db: db}}
}

func (s *userDataSeeder) CreateOne() *entity.User {
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

func (s *userDataSeeder) CreateMany(count int) []*entity.User {
	var users []*entity.User

	for i := 1; i <= count; i++ {
		user := s.CreateOne()
		users = append(users, user)
	}

	return users
}
