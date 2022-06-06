package main

import "fmt"

// go语言的中 return 语句不是原子操作，在底层是分为两步来执行：
// 1、第一步：返回值赋值
// 2、第二步：真正的RET返回
// 函数中如果存在defer，那么derfer执行的时机是在第一步和第二步之间

func main() {

	fmt.Println()

}
