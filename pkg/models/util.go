package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindRequestBody(c *gin.Context, data interface{}) bool {
	if err := c.ShouldBindJSON(data); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Type  string
			Title string
		}{
			Type:  "INVALID_RESOURCE_ID",
			Title: "Invalid resource id provided",
		})
		c.Abort()
		return false
	}

	return true
}
