package routers

import (
	"Firstgin/controllers/user"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine, jwtMiddleware *jwt.GinJWTMiddleware) {
	apiRouter := router.Group("/user")
	apiRouter.Use(jwtMiddleware.MiddlewareFunc())
	apiRouter.POST("/add", user.Controller{}.CreateUser)
	apiRouter.POST("/delete", user.Controller{}.DeleteUser)
	apiRouter.POST("/update", user.Controller{}.UpdateUser)
	apiRouter.POST("/retrieve", user.Controller{}.RetrieveUser)
}
