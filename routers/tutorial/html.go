package tutorial

import (
	"Firstgin/controllers/tutorial"
	"github.com/gin-gonic/gin"
)

func HtmlRouter(router *gin.Engine) {
	router.GET("/formhtml", tutorial.FormHtml)
}
