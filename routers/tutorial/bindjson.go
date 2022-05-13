package tutorial

import (
	"Firstgin/controllers/tutorial"
	"github.com/gin-gonic/gin"
)

func JsonRouter(router *gin.Engine) {
	router.POST("/json", tutorial.BindJson)
}
