package tutorial

import (
	"Firstgin/controllers/tutorial"
	"github.com/gin-gonic/gin"
)

func QueryRouter(router *gin.Engine) {
	router.GET("/urlparam", tutorial.UrlParam)
}
