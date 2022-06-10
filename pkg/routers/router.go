package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/suvam720/url-shortner/pkg/controllers"
)

func Router() *gin.Engine {

	r := gin.Default()

	r.GET("api/url", controllers.GetAllUrl)
	r.POST("api/shorten", controllers.CreateShortUrl)
	r.GET("api/:url", controllers.RedirectUrl)
	r.DELETE("api/:url", controllers.DeleteUrl)

	return r
}
