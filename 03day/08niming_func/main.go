package main

import "fmt"

// 匿名函数

var f1 = func(x, y int) int {
	return x + y
}

func main() {
	fmt.Printf("%T\n", f1)
	fmt.Println(f1)
	fmt.Println(f1(10, 20))
	// 匿名函数的用途
	// 命名函数内部无法在声明有名字的函数

	// 匿名函数声明之后立即执行
	func() {
		fmt.Println("匿名函数立即执行")
	}()
}
