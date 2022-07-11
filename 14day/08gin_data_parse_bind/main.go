package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	// binding:"required" 必须字段，如果传递的是空值，则会报错
	User     string `json:"user" xml:"user" form:"user" uri:"user" binding:"required" ini:"user"`
	Password string `json:"password" xml:"password" form:"password" uri:"password" binding:"required" ini:"password"`
}

func main() {
	r := gin.Default()

	/// json 绑定
	r.POST("loginJSON", func(c *gin.Context) {
		// 设置接收的变量
		var json Login
		// 把 request 中body 中数据 按照 json 格式解析到 结构体中 Login 中
		// gin.H 小工具分装了生成json的方法
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  err.Error(),
				"data": nil,
			})
			return
		}
		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 304,
				"msg":  "登录信息错误",
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "登录成功",
			"data": nil,
		})
	})
	r.Run(":8000")
}
