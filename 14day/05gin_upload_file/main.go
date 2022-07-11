package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		// 从表单中获取文件
		file, _ := c.FormFile("file")
		fmt.Println("file name:", file.Filename)
		// 将图片保存到项目根目录
		_ = c.SaveUploadedFile(file, file.Filename)
		// 打印信息
		c.String(http.StatusOK, fmt.Sprintf(`%s upload success!`, file.Filename))
	})
	r.Run(":8000")
}
