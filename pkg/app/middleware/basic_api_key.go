package middleware

import (
	"net/http"

	"github.com/bitwyre/template-golang/pkg/helper"
	"github.com/bitwyre/template-golang/pkg/lib"
	"github.com/bitwyre/template-golang/pkg/model"
	"github.com/gin-gonic/gin"
)

func BasicAPIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-Api-Key")

		if lib.AppConfig.App.Env == "prod" {
			if apiKey == "" || apiKey != lib.AppConfig.App.BasicApiKey {
				helper.HttpErrorResponse(http.StatusUnauthorized, model.BaseErrorResponseSchema{
					Code:    "UNAUTHORIZED",
					Message: "Invalid authentication header",
				}, c)
				c.Abort()
			}
		}

		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Next()
	}
}
