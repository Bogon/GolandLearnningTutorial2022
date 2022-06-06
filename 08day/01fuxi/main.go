package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel

func sendNum(ch chan<- int) {
	for {
		num := rand.Intn(10)
		ch <- num
		time.Sleep(time.Second * 4)
	}
}

func main() {
	ch := make(chan int, 1)
	// ch <- 100
	// ch <- 200
	// <-ch          // 从通道中取出来
	// x, ok := <-ch // 再去阻塞
	// fmt.Println(x, ok)

	go sendNum(ch)
	for {
		x, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(x)
	}

	for x := range ch {
		fmt.Println(x)
	}
}
