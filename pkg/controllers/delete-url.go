package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suvam720/url-shortner/pkg/helpers"
)

func DeleteUrl(c *gin.Context) {
	url := c.Param("url")
	if err := helpers.DeleteUrl(url); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "url does not exist..",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted sucessfully..",
	})
}
