package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/// 路由
	r := gin.Default()
	// API传递参数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		// 添加参数
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})
	r.Run(":8000")
}
