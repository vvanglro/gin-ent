package routers

import (
	"Firstgin/controllers"
	"Firstgin/controllers/user"
	"Firstgin/global"
	"Firstgin/routers/tutorial"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {

	f := global.FS
	router := gin.Default()
	// 限制每次处理文件所占用的最大内存，上传后的文件
	//https://www.jianshu.com/p/9ec31f232f5f
	router.MaxMultipartMemory = 8 << 20 //字节单位 8左移运算符20 就是8388608  这里是8mb等于8388608字节
	tempL := template.Must(template.New("").ParseFS(f, "templates/html/*.html"))
	router.SetHTMLTemplate(tempL)
	// example: /public/static/js/a.js
	router.StaticFS("/public", http.FS(f))
	//setUpConfig(router)
	setUpRouter(router)

	return router
}

func setUpRouter(router *gin.Engine) {
	jwtMiddleware, _ := user.AuthInit()
	router.GET("/", controllers.HelloIndex)
	tutorial.PathRouter(router)
	tutorial.QueryRouter(router)
	tutorial.FormRouter(router)
	tutorial.HtmlRouter(router)
	tutorial.UploadRouter(router)
	tutorial.JsonRouter(router)
	tutorial.BindFormRouter(router)
	tutorial.BindUrlRouter(router)
	tutorial.AuthRouter(router, jwtMiddleware)
	UserRouter(router, jwtMiddleware)
}
