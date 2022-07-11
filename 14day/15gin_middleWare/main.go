package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 全局中间件/局部中间件

/// 定义中间件
func middleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := time.Now()
		fmt.Println("中间件开始执行了…")
		// 设置变量到 context 的key中，后面可以通过 get() 获取
		c.Set("request", "中间件")
		// 执行中间件
		c.Next()
		// 中间件执行完后续的一些事
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕…", status)

		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	r := gin.Default()

	r.Use(middleWare())
	// {} 为了代码 规范
	{
		// 同步和异步
		r.GET("/middleware", func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("get key request, value: ", req)

			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"msg":    "success",
				"data":   req,
			})
		})
	}

	// 单个路由后面跟的单个/局部中间件
	r.GET("/middleware2", middleWare(), func(c *gin.Context) {
		req, _ := c.Get("request")
		fmt.Println("get key request, value: ", req)

		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "success",
			"data":   req,
		})
	})
	_ = r.Run(":8000") // 不处理返回结果
}
