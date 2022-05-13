package tutorial

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UrlParam(c *gin.Context) {
	// URL参数可以通过DefaultQuery()或Query()方法获取
	// DefaultQuery()若参数不村则，返回默认值，Query()若不存在，返回空串
	// API ? name=zs
	name := c.DefaultQuery("name", "古藤")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))

}
