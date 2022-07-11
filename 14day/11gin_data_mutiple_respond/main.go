package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

/// 多种响应方式

func main() {
	r := gin.Default()

	// 1. json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "someJSON success",
			"status": 0,
			"data":   nil,
		})
	})

	// 2. 结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}

		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 12

		c.JSON(http.StatusOK, msg)
	})

	// 3. xml
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "XML", "status": 200, "data": nil})
	})

	// 4. YAML 响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "YAML success",
			"status":  200,
			"data":    nil,
		})
	})

	// 5. protobuffer, google 开发的高效的存储和读取的工具
	r.GET("someProtoBuf", func(c *gin.Context) {
		resp := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传递数据 protobuffer 格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  resp,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	r.Run(":8000")
}
