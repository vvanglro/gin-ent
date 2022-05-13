package tutorial

import (
	"Firstgin/models/ent"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine, jwtMiddleware *jwt.GinJWTMiddleware) {
	apiRouter := router.Group("/api")

	// 登录 api ，直接用 jwtMiddleware 中的 `LoginHandler` 既可
	apiRouter.POST("/login", jwtMiddleware.LoginHandler)

	// 刷新 token api
	apiRouter.GET("/refresh_token", jwtMiddleware.RefreshHandler)

	// 当 api 需要验证的时候只需要使用 jwtMiddleware.MiddlewareFunc() 中间件既可
	authRouter := apiRouter.Group("/auth")
	authRouter.Use(jwtMiddleware.MiddlewareFunc())
	{
		authRouter.GET("/hello", func(ctx *gin.Context) {
			claims := jwt.ExtractClaims(ctx)
			fmt.Println(claims)
			user, _ := ctx.Get("id")
			fmt.Println(user)
			ctx.JSON(200, gin.H{
				"userID": claims["id"],
				//user.(*global.User)类型断言
				"userName": user.(*ent.User).Name,
				"text":     "Hello World.",
			})
		})

	}
}
