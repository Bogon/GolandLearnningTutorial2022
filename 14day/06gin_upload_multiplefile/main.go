package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 限制表单上传的大小 8MB 默认表单大小： 32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		// 从表单中获取文件
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form data failed, error: %v", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 循环遍历所有的图片
		for _, file := range files {
			//保存所有图片
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload mutiple file failed, error: %v", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("upload %v file success!", len(files)))
	})
	r.Run(":8000")
}
