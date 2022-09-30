package rest

import (
	"net/http"
	"strconv"

	"github.com/bitwyre/template-golang/pkg/helper"
	"github.com/bitwyre/template-golang/pkg/model"
	"github.com/gin-gonic/gin"
)

func (h *Rest) GetUserRestHandler(c *gin.Context) {
	getId := c.Param("userid")

	userId, err := strconv.Atoi(getId)
	if err != nil {
		helper.HttpErrorResponse(http.StatusBadRequest, model.BaseErrorResponseSchema{
			Code:    "BAD_REQUEST",
			Message: "Should be number",
			Errors:  nil,
		}, c)
		return
	}

	s, err := h.IUserService.GetUser(userId)
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
