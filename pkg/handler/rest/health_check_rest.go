package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Rest) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
