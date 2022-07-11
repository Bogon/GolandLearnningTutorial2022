package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleWare 中间件：登录、进入主页
func AuthMiddleWare(c *gin.Context) {

}

func main() {
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		// 设置 cookie
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		// 返回信息
		c.String(http.StatusOK, "login success!")
	})

	r.GET("/home", AuthMiddleWare, func(c *gin.Context) {
		// 验证 cookie

	})

	_ = r.Run(":8000")
}
