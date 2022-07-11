package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/// URI 数据解析与绑定

type Login struct {
	// binding:"required" 必须字段，如果传递的是空值，则会报错
	User     string `json:"username" xml:"username" form:"username" uri:"username" binding:"required" ini:"username"`
	Password string `json:"password" xml:"password" form:"password" uri:"password" binding:"required" ini:"password"`
}

func main() {
	r := gin.Default()

	/// json 绑定
	r.GET("/:username/:password", func(c *gin.Context) {
		// 设置接收的变量
		var login Login
		// Bind() 默认解析并绑定form格式数据
		// 根据请求头中 content-type 自动推断
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  err.Error(),
				"data": nil,
			})
			return
		}

		if login.User != "root" || login.Password != "admin" {
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
