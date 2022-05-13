package tutorial

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ParamFunc(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	//截取
	action = strings.Trim(action, "/")
	c.String(http.StatusOK, name+" is "+action)

}
