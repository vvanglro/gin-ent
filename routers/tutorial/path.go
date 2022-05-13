package tutorial

import (
	"Firstgin/controllers/tutorial"
	"github.com/gin-gonic/gin"
)

func PathRouter(router *gin.Engine) {
	router.GET("/param/:name/*action", tutorial.ParamFunc)
}
