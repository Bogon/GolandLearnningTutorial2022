package main

import (
	"fmt"
	"sync"
)

var b chan int // 需要指定通道中元素的类型

var wg sync.WaitGroup

func noBufChan() {
	fmt.Println(b)
	wg.Add(1)
	b = make(chan int) // 不带缓冲区通道的初始化
	go func() {
		x := <-b
		defer wg.Done()
		fmt.Println("数从后台读取通道中数据：", x)
	}()
	b <- 10 // Hang住了
	fmt.Println("数据发送到的通道中……")
	wg.Wait()
}

func bufChan() {
	fmt.Println(b)
	b = make(chan int, 1) // 不带缓冲区通道的初始化
	b <- 10               // Hang住了
	fmt.Println("数据发送到的通道中……")
	x := <-b
	fmt.Println("数从后台读取通道中数据：", x)

	close(b)
}

func main() {
	// noBufChan()
	bufChan()
}
