package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 使用 map 拼接数据返回
	// 使用 map[string]interface{}
	r.GET("/index", func(c *gin.Context) {
		// 方式一：自己拼接 JOSN
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world~",
			"code":    200,
			"data":    nil,
		})
	})

	// 使用 struct 返回
	r.GET("/more", func(c *gin.Context) {
		// 方式二：使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string `json:"message"`
			Age     int    `json:"age"`
		}
		msg.Name = "张珊"
		msg.Message = "大学一年级，年龄 20岁，主修人工智能领域"
		msg.Age = 20
		c.JSON(http.StatusOK, msg)
	})

	// 开启网络请求
	r.Run(":8084")
}
