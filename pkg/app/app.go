package app

import (
	"github.com/bitwyre/template-golang/pkg/app/middleware"
	"github.com/bitwyre/template-golang/pkg/datastore/mysql"
	"github.com/bitwyre/template-golang/pkg/handler/rest"
	"github.com/bitwyre/template-golang/pkg/lib"
	"github.com/bitwyre/template-golang/pkg/repository"
	"github.com/bitwyre/template-golang/pkg/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClientApp struct {
	Db *gorm.DB
}

func BootstrapApp(r *gin.Engine) ClientApp {
	sql := mysql.MySQLDriver()
	sql.MySQLMigration()

	r.Use(middleware.SetUpCors(lib.AppConfig.App.FrontEndURL))

	repo := repository.NewRepository(sql.Db)
	rootService := service.NewService(repo)
	rootRestHandler := rest.NewRest(rootService)

	NewRoutes(r, rootRestHandler)

	return ClientApp{
		Db: sql.Db,
	}
}
