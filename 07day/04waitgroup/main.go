package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wait group

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int() // int64 随机数
		r2 := rand.Intn(10)
		fmt.Println(r1, "  ", r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {

	// f()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}

	// 怎么知道 goroutine 执行结束？
	wg.Wait() // 等待 wg的计数器为0

}
