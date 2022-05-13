package tutorial

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传文件错误")
	}
	saveErr := c.SaveUploadedFile(file, file.Filename)
	if saveErr != nil {
		c.String(500, "保存文件错误")
	}
	c.String(200, "文件上传成功")
}

func Uploads(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
	}
	// 获取所有文件
	files := form.File["files"]
	// 遍历所有文件
	for _, file := range files {
		// 逐个存
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			return
		}
	}
	c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
}
