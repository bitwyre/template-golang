package app

import (
	"github.com/bitwyre/template-golang/pkg/handler/rest"
	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.Engine, rest *rest.Rest) {
	// Rest API Routes
	r.GET("/health", rest.HealthCheck)
	r.GET("/user/:userid", rest.GetUserRestHandler)

	// You can include middleware.BasicAPIKeyMiddleware() to use API Key Middleware
}
