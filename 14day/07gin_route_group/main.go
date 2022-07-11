package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	// 路由组 route group
	v1 := r.Group("/v1")
	// 书写规范
	{
		v1.GET("/login", login)
		v1.POST("/submit", submit)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/user", submit)
	}
	r.Run(":8000")
}

// It takes a pointer to a gin.Context and returns nothing
func login(c *gin.Context) {
	name := c.DefaultQuery("name", "lilei")
	c.String(http.StatusOK, fmt.Sprintf("hello %v \n", name))
}

// `submit` is a function that takes a pointer to a gin.Context and returns nothing
func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "nancy")
	c.String(http.StatusOK, fmt.Sprintf("hello %v \n", name))
}
