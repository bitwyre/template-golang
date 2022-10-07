package app

import (
	"log"

	"github.com/bitwyre/template-golang/pkg/datastore/mysql"
	grpcHandler "github.com/bitwyre/template-golang/pkg/handler/grpc"
	"github.com/bitwyre/template-golang/pkg/handler/rest"
	"github.com/bitwyre/template-golang/pkg/lib"
	"github.com/bitwyre/template-golang/pkg/repository"
	"github.com/bitwyre/template-golang/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ClientApp struct {
	Db   *gorm.DB
	Grpc *grpc.Server
}

func BootstrapApp(r *gin.Engine) *ClientApp {
	err := InitOpenTelemetry()
	if err != nil {
		log.Fatalln("🔴  Couldn't connect to jaeger:" + err.Error())
	}

	sql := mysql.MySQLDriver()
	sql.MySQLMigration()
	if err := sql.Db.Use(otelgorm.NewPlugin()); err != nil {
		logrus.Fatal(err)
	}

	r.Use(otelgin.Middleware(lib.AppConfig.App.ServiceName))

	repo := repository.NewRepository(sql.Db)
	rootService := service.NewService(repo)
	rootRestHandler := rest.NewRest(rootService)
	NewRoutes(r, rootRestHandler)

	// Init GRPC Server
	grpcServer := grpc.NewServer()
	rpc := grpcHandler.NewGRPC(rootService)
	grpcHandler.RegisterGRPCService(grpcServer, rpc)

	return &ClientApp{
		Db:   sql.Db,
		Grpc: grpcServer,
	}
}
