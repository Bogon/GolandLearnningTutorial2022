package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1. 创建路由
	// 默认使用了 2个中间件：Logger(), Recovery()
	// 也可以创建没有中间件的路由: r := gin.New()
	r := gin.Default()
	// 2. 绑定路由规则，执行的函数
	// *gin.Context 内部封装了 request 和 response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, `{"code": 200, "msg": "success", "data": {"msg": "Hello, world!"}}`)
	})

	r.POST("/xxxPost", func(c *gin.Context) {
		c.String(http.StatusOK, `{"code": 200, "msg": "success", "data": {"name": "沙河娜扎", "age": 18, "married": true}}`)
	})

	r.PUT("/xxxPUT", modify)

	r.DELETE("/xxxDELETE", deleteD)

	// 设置监听端口，默认8080端口
	r.Run(":8000")
}

// deleteD ...
func deleteD(c *gin.Context) {
	c.String(http.StatusOK, `{"code": 200, "msg": "success", "data": {"name": "曼陀罗", "age": 120, "isKilled": false}}`)
}

// > The function `modify` takes a pointer to a `gin.Context` and returns nothing
func modify(c *gin.Context) {
	c.String(http.StatusOK, `{"code": 200, "msg": "success", "data": {"name": "曼陀罗", "age": 120, "isKilled": false}}`)
}
