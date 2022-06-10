package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suvam720/url-shortner/pkg/helpers"
)

func RedirectUrl(c *gin.Context) {
	url := c.Param("url")
	u, err := helpers.FindUrl(url)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	c.JSON(http.StatusAccepted, gin.H{"url": u})
	// c.Redirect(http.StatusPermanentRedirect, u)
}
