package test_config

import (
	"github.com/bitwyre/template-golang/pkg/app"
	"github.com/bitwyre/template-golang/pkg/lib"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type ClientTestSuite struct {
	Gin *gin.Engine
	Db  *gorm.DB
}

func BootstrapAppTest() *ClientTestSuite {
	_ = godotenv.Load()
	lib.InitAppConfig(true)
	lib.GetLogger()

	var r = gin.Default()
	r.Use(lib.GinLogger(), gin.Recovery())
	instance := app.BootstrapApp(r)

	return &ClientTestSuite{
		Gin: r,
		Db:  instance.Db,
	}
}
