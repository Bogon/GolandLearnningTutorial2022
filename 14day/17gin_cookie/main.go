package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 会话控制
// 1、cookie 是什么
// 2、cookie的用途
// 2.1、保存用户登录状态

// 目标：测试服务端发送cookie给客户端，客户端请求携带cookie给服务端

func main() {
	r := gin.Default()

	// 服务端给客户端 cookie
	r.GET("/cookie", func(c *gin.Context) {
		// 获取客户端是否携带 cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 设置 cookie (name string, value string,
			// maxAge int 单位为秒,
			// path string, cookie 所在目录
			// domain string, 域名
			// secure bool,	是否只能通过 https 来访问
			// httpOnly bool 是否允许别人使用 js 来获取 cookie
			c.SetCookie("key_cookie", "value_cookie", 12000, "/cookie", "localhost", false, false)
		}

		fmt.Println("cookie 的值是：", cookie)
	})
	_ = r.Run(":8000")
}
