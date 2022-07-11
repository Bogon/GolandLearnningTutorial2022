package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/form", func(c *gin.Context) {
		// 表单设置默认值
		c.DefaultPostForm("type", "alert")
		// 接收表单值
		username := c.PostForm("username")
		password := c.PostForm("password")
		hobby := c.PostFormArray("hobby")

		c.String(http.StatusOK, fmt.Sprintf(`{"username": %v, "password": %v, "hobby": %v}`, username, password, hobby))
	})
	r.Run(":8000")
}
