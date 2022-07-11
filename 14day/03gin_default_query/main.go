package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "lilei")
		c.String(http.StatusOK, fmt.Sprintf("Hello %v", name))
	})

	r.Run(":8000")
}
