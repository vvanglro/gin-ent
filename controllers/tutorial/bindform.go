package tutorial

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FormLogin struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User     string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func BindForm(c *gin.Context) {
	var form FormLogin
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	err := c.Bind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if form.User != "root" || form.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
