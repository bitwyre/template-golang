package helper

import (
	"net/http"

	"github.com/bitwyre/template-golang/pkg/model"
	"github.com/gin-gonic/gin"
)

func restResponse(httpCode int, responseSchema interface{}, errorSchema model.BaseErrorResponseSchema, c *gin.Context) {
	var traceId = c.GetHeader("X-Trace-Id")
	var isSuccess = true
	var errData interface{}

	if errorSchema.Code != "" {
		isSuccess = false
		errData = errorSchema
	}

	c.JSON(httpCode, model.BaseResponseSchema{
		Success: isSuccess,
		TraceId: traceId,
		Error:   errData,
		Results: &responseSchema,
	})
}

func HttpSuccess(responseSchema interface{}, c *gin.Context) {
	var errorSchema = model.BaseErrorResponseSchema{}
	restResponse(http.StatusOK, &responseSchema, errorSchema, c)
}

func HttpErrorResponse(httpCode int, errorSchema model.BaseErrorResponseSchema, c *gin.Context) {
	restResponse(httpCode, nil, errorSchema, c)
}
