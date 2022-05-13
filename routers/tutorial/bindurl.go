package tutorial

import (
	"Firstgin/controllers/tutorial"
	"github.com/gin-gonic/gin"
)

func BindUrlRouter(router *gin.Engine) {
	router.POST("/bindurl/:user/:password", tutorial.BindUrl)
}
