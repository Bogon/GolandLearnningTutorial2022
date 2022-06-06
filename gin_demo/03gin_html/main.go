package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建的一个 gin 实例
	r := gin.Default()
	// 加载/渲染 HTML 信息
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	r.Run(":8080")
}
