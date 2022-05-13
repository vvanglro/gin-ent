package tutorial

import (
	"Firstgin/controllers"
	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

func BindJson(c *gin.Context) {
	var json Login
	// 将request的body中的数据，自动按照json格式解析到结构体
	err := c.ShouldBindJSON(&json)
	if err != nil {
		controllers.Response.Fail(c, 10013, "输入错误", err.Error(), controllers.HttpCode(401))
		return
	}
	if json.User != "root" || json.Password != "admin" {
		controllers.Response.Fail(c, 10014, "输入错误", "")
		return
	}
	controllers.Response.Success(c, json)
	return
}
