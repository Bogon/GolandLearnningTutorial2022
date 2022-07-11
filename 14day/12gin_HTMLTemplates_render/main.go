package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/// HTML 模板渲染
// 本质上就是字符串替换

func main() {
	r := gin.Default()

	// html 模板渲染
	// 记载模板文件
	r.LoadHTMLGlob("templates/*")
	//r.LoadHTMLFiles("templtes/index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		// 根据文件名渲染
		// 最终 json 将 title 替换
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Hello Goland",
		})
	})
	r.Run(":8000")
}
