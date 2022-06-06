package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	router := gin.Default()
	// 创建模板信息
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// 加载模板文件
	router.LoadHTMLFiles("./index.tmpl")

	// 请求地址
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://xuebaonline.com'> Char's 博客</a>")
	})

	router.GET("/subPath", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href=`https://www.liwenzhou.com`> 推荐的博客地址</a>")
	})

	router.Run(":8080")
}
