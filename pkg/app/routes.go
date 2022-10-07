package app

import (
	"github.com/bitwyre/template-golang/pkg/app/middleware"
	"github.com/bitwyre/template-golang/pkg/handler/rest"
	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.Engine, rest *rest.Rest) {
	public(r, rest)
	private(r, rest)

}

func public(r *gin.Engine, rest *rest.Rest) {
	public := r.Group("/public")
	public.GET("/health", rest.HealthCheck)
}

func private(r *gin.Engine, rest *rest.Rest) {
	private := r.Group("/private")

	private.Use(middleware.BasicAPIKeyMiddleware())
	private.Use(middleware.JWTMiddleware())

	private.GET("/user/:userid", rest.GetUserRestHandler)
}
