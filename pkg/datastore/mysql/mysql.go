package mysql

import (
	"fmt"
	"log"

	"github.com/bitwyre/template-golang/pkg/datastore/mysql/entity"
	"github.com/bitwyre/template-golang/pkg/lib"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var datetimePrecision = 0

type ClientInstanceMySQL struct {
	Db *gorm.DB
}

func MySQLDriver() *ClientInstanceMySQL {
	var env = lib.AppConfig.App
	dsn := fmt.Sprintf(`%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`, env.SqlUser, env.SqlPassword, env.SqlHost, env.SqlPort, env.SqlDB)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultDatetimePrecision: &datetimePrecision,
	}))

	if err != nil {
		log.Fatalln("🔴  MySQL Error:" + err.Error())
		return nil
	}

	log.Println("🚀 MySQL database connected")

	return &ClientInstanceMySQL{Db: db}
}

func (client *ClientInstanceMySQL) MySQLMigration() {
	err := client.Db.AutoMigrate(
		&entity.User{},
	)
	if err != nil {
		log.Fatalln("🔴 Database Migration Failed", err)
	}
}
