package main

import (
	"context"
	"fmt"
	"time"
)

// context.WithDeadline

func main() {

	d := time.Now().Add(time.Millisecond * 2000)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管 context 会过期，但在任何情况下调用它的 cancel 函数都是很好的实践
	// 如果不是这样，可能会使上下文机器父类存活的时间超过必要的时间
	defer cancel()

	select {
	case <-time.After(time.Second * 1):
		fmt.Println("周林")
	case <-ctx.Done():
		fmt.Println(ctx.Err())

	}

}
