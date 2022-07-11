package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// MyTime 中间件：定义一个中间件，记录请求的的时间
func MyTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		since := time.Since(start)
		fmt.Println("程序用时:", since)

	}
}

func main() {

	r := gin.Default()
	r.Use(MyTime())

	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)

		shoppingGroup.GET("/home", shopHomeHandler)
	}

	_ = r.Run(":8000")
}

// ShopIndexHandler is a function that takes a pointer to a gin.Context and returns nothing.
func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "get index success",
		"data":   nil,
	})
}

// `shopHomeHandler` is a function that takes a pointer to a gin.Context and returns nothing
func shopHomeHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "get home success",
		"data":   nil,
	})
}
