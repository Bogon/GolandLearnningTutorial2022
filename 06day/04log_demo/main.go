package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// log demo

func main() {
	file, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	log.SetOutput(file)

	for {
		log.Println("这是一条测试日志……")
		time.Sleep(3 * time.Second)
	}

}

/**
日志库 需求分析
1、支持向不同的地方输出日志
2、日志分级别
	1、Debug
	2、Trace
	2、Info
	3、Warning
	4、Error
	5、Fatal
3、日志要支持开关控制
4、日志要有时间、行号、文件名、日志级别、日志信息
5、日志文件要切割
*/
