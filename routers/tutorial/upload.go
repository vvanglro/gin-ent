package tutorial

import (
	"Firstgin/controllers/tutorial"
	"github.com/gin-gonic/gin"
)

func UploadRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/upload", tutorial.Upload)
	}
	v2 := router.Group("/v2")
	{
		v2.POST("/uploads", tutorial.Uploads)
	}
}
