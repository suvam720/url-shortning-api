package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suvam720/url-shortner/pkg/helpers"
	"github.com/suvam720/url-shortner/pkg/request"
)

func CreateShortUrl(c *gin.Context) {
	var url request.RequestBody
	if err := c.ShouldBindJSON(&url); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	randomStr := helpers.Random(6)

	if err := helpers.CreateUrl(&url, randomStr); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"shortUrl": randomStr,
	})
}
