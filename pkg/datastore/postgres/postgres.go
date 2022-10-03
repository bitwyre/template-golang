package postgres

import (
	"fmt"
	"log"

	"github.com/bitwyre/template-golang/pkg/datastore/postgres/entity"

	"github.com/bitwyre/template-golang/pkg/lib"
	pgDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ClientInstancePG struct {
	Db *gorm.DB
}

func PGDriver() *ClientInstancePG {
	var env = lib.AppConfig.App
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta`, env.PgHost, env.PgUser, env.PgPassword, env.PgDB, env.PgPort)

	db, err := gorm.Open(pgDriver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln("🔴  Postgres Error:" + err.Error())
		return nil
	}

	log.Println("🚀 Postgres database connected")

	return &ClientInstancePG{Db: db}
}

func (client *ClientInstancePG) AutoMigrate() {
	err := client.Db.AutoMigrate(
		&entity.User{},
	)
	if err != nil {
		log.Fatalln("🔴 Database Migration Failed", err)
	}
}
