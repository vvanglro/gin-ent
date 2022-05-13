package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
}
