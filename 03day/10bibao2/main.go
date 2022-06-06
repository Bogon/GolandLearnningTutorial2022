package main

import "fmt"

// 闭包是什么
// 闭包是一个函数，这个函数包含了一个他外部作用于的一个变量

// 底层原理
// 1、函数可以作为返回值
// 2、函数内部变量查找顺序，先在自己内部找，找不到向外部找

func addr() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret := addr()
	res := ret(200)
	fmt.Println(res)
}
