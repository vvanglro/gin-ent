package main

import (
	"Firstgin/conf"
	"Firstgin/models"
	"Firstgin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	InitEmbed()
	conf.ConfigInit()
	models.EntInit()
	gin.SetMode(gin.DebugMode) // 设置为生产模式
	r := routers.InitRouter()
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	println(err)
}
