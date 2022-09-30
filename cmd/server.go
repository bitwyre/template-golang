package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/bitwyre/template-golang/pkg/app"
	"github.com/bitwyre/template-golang/pkg/lib"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Server() {
	_ = godotenv.Load()
	lib.InitAppConfig(false)
	lib.GetLogger()

	if lib.AppConfig.App.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
	}

	r := gin.Default()
	r.Use(lib.GinLogger(), gin.Recovery())

	// Trusted Proxy when on production
	if lib.AppConfig.App.Env == "prod" {
		r.TrustedPlatform = gin.PlatformCloudflare
	}

	log.Println("🔰 Application Environment: ", lib.AppConfig.App.Env)
	app.BootstrapApp(r)

	err := r.Run(fmt.Sprintf(`:%d`, lib.AppConfig.App.ServerPort))
	if err != nil {
		fmt.Println(err)
		return
	}
}
