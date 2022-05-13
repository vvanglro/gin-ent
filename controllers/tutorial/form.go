package tutorial

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FormApi(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("userpassword")
	c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
}
