package models

import (
	"OnlineDeck/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindRequestBody(c *gin.Context, data interface{}) bool {
	if err := c.ShouldBindJSON(data); err != nil {
		c.JSON(http.StatusBadRequest,
			errors.HttpError{
				Type:   "INVALID_BODY_PARAMS",
				Title:  "Invalid params in body",
				Detail: err.Error(),
			})
		c.Abort()
		return false
	}

	return true
}
