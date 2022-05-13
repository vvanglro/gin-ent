package tutorial

import (
	"Firstgin/controllers/tutorial"
	"github.com/gin-gonic/gin"
)

func FormRouter(router *gin.Engine) {
	router.POST("/formapi", tutorial.FormApi)
}
