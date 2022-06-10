package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suvam720/url-shortner/pkg/helpers"
)

func GetAllUrl(c *gin.Context) {
	urls, err := helpers.GetAllUrl()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{"data": urls})
}
