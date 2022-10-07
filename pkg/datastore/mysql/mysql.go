package mysql

import (
	"fmt"
	"log"

	"github.com/bitwyre/template-golang/pkg/datastore/mysql/entity"
	"github.com/bitwyre/template-golang/pkg/lib"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln("🔴  Couldn't connect to database:" + err.Error())
		return nil
	}

	log.Println("🟢 MySQL database connected")

	return &ClientInstanceMySQL{Db: db}
}

func (client *ClientInstanceMySQL) MySQLMigration() {
	err := client.Db.AutoMigrate(
		&entity.User{},
		&entity.Authentication{},
	)
	if err != nil {
		log.Fatalln("🔴 Database migration failed", err)
	}
}
