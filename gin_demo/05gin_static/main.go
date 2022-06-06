package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")

	// 加载……
	r.Run(":8082")
}
