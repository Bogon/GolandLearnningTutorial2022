package main

import (
	"fmt"
	"time"
)

// goroutine

func hello(i int) {
	fmt.Println("hello, world~ id: ", i)
}

// 程序启动之后会启动一个 main goroutine
func main() {
	for i := 0; i < 100; i++ {
		// go hello(i) // 单独开启一个 goroutine 去执行 hello 函数(任务)
		go func(i int) {
			fmt.Println(i) // 用的是函数传递进来的，不是外面的i
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
	// main 函数结束了，由mian函数启动的goroutine 也就结束了
}
