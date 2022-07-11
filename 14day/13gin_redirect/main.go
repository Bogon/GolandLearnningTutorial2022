package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/// 重定向

func main() {
	r := gin.Default()
	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		// 重定向: 支持内部和外部重定向
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	_ = r.Run(":8000") // 不处理返回结果
}
