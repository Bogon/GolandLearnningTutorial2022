package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 为什么要 context
func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
FORLOOP:
	for {
		fmt.Println("run …")
		time.Sleep(time.Millisecond * 500)

		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func f2(ctx context.Context) {
	defer wg.Done()

FORLOOP:
	for {
		fmt.Println("pipipi …")
		time.Sleep(time.Millisecond * 500)

		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func main() {
	ctx, cancle := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancle()
	wg.Wait()
	// 如何通知子 goroutine 退出

}
