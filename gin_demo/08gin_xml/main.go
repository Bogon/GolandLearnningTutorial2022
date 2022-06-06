package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/getXML", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"name"`
			Message string `json:"message"`
			Age     int    `json:"age"`
		}
		msg.Name = "沙河娜扎"
		msg.Message = "沙河娜扎，女，家住北京市昌平区沙河高教园，本科、主攻金融。"
		msg.Age = 21
		c.XML(http.StatusOK, msg)
	})

	r.GET("/getJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"name"`
			Message string `json:"message"`
			Age     int    `json:"age"`
		}
		msg.Name = "沙河娜扎"
		msg.Message = "沙河娜扎，女，家住北京市昌平区沙河高教园，本科、主攻金融。"
		msg.Age = 21
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/getYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})

	r.GET("/getProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})

	r.Run(":8085")
}
