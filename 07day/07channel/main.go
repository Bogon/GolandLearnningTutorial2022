package main

import (
	"fmt"
	"sync"
)

// channel 练习
// 1、启动一个 goroutine ，生成100个数发送到ch1
// 2、启动一个 goroutine ，从ch1中取值，计算奇平方放到ch2
// 3、在main函数中从ch2中取出打印出来

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		i, ok := <-ch1
		if !ok {
			break
		}

		ch2 <- i * i
	}
	once.Do(func() { close(ch2) }) // 确保某个操作只执行一次
}

func main() {
	ch1 := make(chan int, 50)
	ch2 := make(chan int, 100)
	wg.Add(3)
	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)
	wg.Wait()

	fmt.Println("----------------------------------------------")
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
