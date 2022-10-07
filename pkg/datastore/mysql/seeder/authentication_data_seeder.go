package seeder

import (
	"log"
	"time"

	"github.com/bitwyre/template-golang/pkg/datastore/mysql/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

type authenticationDataSeeder struct{ *ClientSeeder }

func AuthenticationDataSeeder(db *gorm.DB) *authenticationDataSeeder {
	return &authenticationDataSeeder{&ClientSeeder{
		db: db,
	}}
}

func (s *authenticationDataSeeder) CreateOne() *entity.Authentication {
	authentication := entity.Authentication{
		UserUuid:          uuid.New().String(),
		ExpireTime:        time.Now().String(),
		TokenType:         "Bearer",
		ExpiresIn:         0000000000,
		AccessToken:       faker.RandomString(10),
		RefreshToken:      faker.RandomString(10),
		RSAPublicKey:      faker.RandomString(10),
		IPAddress:         faker.Internet().IpV4Address(),
		Country:           faker.Address().Country(),
		LastAuthenticated: 0000000000,
		Revoked:           0,
	}

	err := s.db.Create(&authentication).Error
	if err != nil {
		log.Fatal(err)
	}

	return &authentication
}

func (s *authenticationDataSeeder) CreateMany(count int) []*entity.Authentication {
	var authentication []*entity.Authentication

	for i := 1; i <= count; i++ {
		auth := s.CreateOne()
		authentication = append(authentication, auth)
	}

	return authentication
}
