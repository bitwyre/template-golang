package rest

import (
	"net/http"
	"strconv"

	"github.com/bitwyre/template-golang/pkg/helper"
	"github.com/bitwyre/template-golang/pkg/model"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (h *Rest) GetUserRestHandler(c *gin.Context) {
	getId := c.Param("userid")

	span := trace.SpanFromContext(c.Request.Context())
	span.SetAttributes(attribute.String("userId", getId))

	userId, err := strconv.Atoi(getId)
	if err != nil {
		helper.HttpErrorResponse(http.StatusBadRequest, model.BaseErrorResponseSchema{
			Code:    "BAD_REQUEST",
			Message: "Should be number",
			Errors:  nil,
		}, c)
		return
	}

	s, err := h.IUserService.GetUser(userId, c.Request.Context())
	if err != nil {
		helper.HttpErrorResponse(http.StatusBadRequest, model.BaseErrorResponseSchema{
			Code:    "NOT_FOUND",
			Message: "User is not found",
			Errors:  nil,
		}, c)
		return
	}

	helper.HttpSuccess(s, c)
}
