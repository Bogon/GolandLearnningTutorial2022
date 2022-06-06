package main

import "fmt"

// 闭包

func f1(f func()) {
	fmt.Println("this si f1!")
	f()
}

// 目的，想要在 f1中执行f2
func f2(x, y int) {
	fmt.Println("this si f2!")
	fmt.Println(x + y)
}

// 定义一个函数对 f2 包装
func f3(f func(int, int), m, n int) func() {
	fmt.Println("this si f3!")
	tmp := func() {
		f(m, n)
	}
	return tmp
}

func main() {

	ret := f3(f2, 100, 200)
	f1(ret)
}
