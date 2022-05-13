package tutorial

import (
	"Firstgin/controllers/tutorial"
	"github.com/gin-gonic/gin"
)

func BindFormRouter(router *gin.Engine) {
	router.POST("/bindform", tutorial.BindForm)
}
