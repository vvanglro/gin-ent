package tutorial

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FormHtml(c *gin.Context) {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	// https://stackoverflow.com/questions/68836237/how-to-get-full-server-url-from-any-endpoint-handler-in-gin
	fmt.Println("the url:", scheme+"://"+c.Request.Host+c.Request.URL.Path)
	c.HTML(http.StatusOK, "form.html", gin.H{"formapi": scheme + "://" + c.Request.Host + "/formapi"})
}
