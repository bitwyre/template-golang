package seeder

import (
	"github.com/bitwyre/template-golang/pkg/datastore/mysql"
	"github.com/bitwyre/template-golang/pkg/lib"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type ClientSeeder struct {
	db *gorm.DB
}

func Exec() {
	_ = godotenv.Load()
	lib.InitAppConfig(false)

	mysqlClient := mysql.MySQLDriver()
	mysqlClient.MySQLMigration()

	AuthenticationDataSeeder(mysqlClient.Db).CreateMany(10)
	UserDataSeeder(mysqlClient.Db).CreateMany(10)
}
