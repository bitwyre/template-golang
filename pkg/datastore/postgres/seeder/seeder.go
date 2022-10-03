package seeder

import (
	"github.com/bitwyre/template-golang/pkg/datastore/postgres"
	"github.com/bitwyre/template-golang/pkg/lib"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func Exec() {
	_ = godotenv.Load()
	lib.InitAppConfig(false)

	pgClient := postgres.PGDriver()
	pgClient.AutoMigrate()

	UserDataSeeder(pgClient.Db).CreateMany(10)
}
