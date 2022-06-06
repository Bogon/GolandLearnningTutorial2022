package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2/json"
	"net/http"
)

func main() {
	// Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "沙河娜扎")
		//username := c.Query("username")
		address := c.Query("address")
		// 组成 JSON 返回给 对方
		c.JSON(http.StatusOK, gin.H{
			"status":   0,
			"username": username,
			"address":  address,
		})
	})

	r.POST("/user/search", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")
		// 组成 JSON 返回给 对方
		c.JSON(http.StatusOK, gin.H{
			"status":   0,
			"username": username,
			"address":  address,
		})
	})

	r.POST("/json", func(c *gin.Context) {
		// 注意：下面为了举例子方便，暂时忽略了错误处理
		b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
		// 定义map或结构体
		var m map[string]interface{}
		// 反序列化
		_ = json.Unmarshal(b, &m)
		c.JSON(http.StatusOK, m)
	})

	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	r.Run()
}
