package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

//// 重定向

func main() {
	r := gin.Default()

	// 同步和异步
	r.GET("/long_async", func(c *gin.Context) {
		// 异步处理
		// goroutine 可以很方便的开启衣蛾异步任务
		// 在启动新的 goroutine 时，不应该使用原始上下文，必须使用使用它的原始副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(time.Second * 3)
			log.Println("异步处理：" + copyContext.Request.URL.Path)

		}()
	})
	// 同步处理
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(time.Second * 3)
		log.Println("同步处理：" + c.Request.URL.Path)
	})
	_ = r.Run(":8000") // 不处理返回结果
}
