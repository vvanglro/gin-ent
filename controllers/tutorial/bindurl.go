package tutorial

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UrlLogin 定义接收数据的结构体
type UrlLogin struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User     string `uri:"user"  binding:"required"`
	Password string `uri:"password"  binding:"required"`
}

func BindUrl(c *gin.Context) {
	var urlLogin UrlLogin
	err := c.ShouldBindUri(&urlLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if urlLogin.User != "root" || urlLogin.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
