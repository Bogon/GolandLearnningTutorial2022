package main

import (
	"learning.goland.com/studyGoland/06day/mylogger"
)

// 测试自己写的日志库
var log mylogger.Logger // 声明一个全局的接口变量

func main() {
	// log = mylogger.NewConsoleLog("debug")	// 终端日志模式
	log = mylogger.NewFileLogger("debug", "./", "file_text.log", 1024) // 写入文件模式

	name := "力王"
	age := 100
	for {
		log.Debug("这是一条 Debug 日志…… name = %s, age = %d", name, age)
		log.Trace("这是一条 Trace 日志……")
		log.Info("这是一条 Info 日志……")
		log.Warning("这是一条 Warning 日志……")
		log.Error("这是一条 Error 日志……")
		log.Fatal("这是一条 Fatal 日志……")
		// time.Sleep(time.Second)
	}

}
