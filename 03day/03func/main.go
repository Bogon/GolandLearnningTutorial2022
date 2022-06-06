package main

import (
	"fmt"
)

// 函数： 一段代码的封装
func f1() {
	fmt.Println("hello")
}

func f2(s string) {
	fmt.Println("hello ", s)
}

// 带参数和返回值的函数
func f3(x int, y int) int {
	sum := x + y
	return sum
}

// 参数类型简写
func f4(x, y int) int {
	return x + y
}

// 可变参数
func f5(x int, y ...int) int {
	fmt.Println(y)
	return x
}

// 类型返回值
func f6(x, y int) (ret int) {
	ret = x + y
	return // 可以省略 返回值变量
}

// 多个返回值
func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {

	f1()

	f2("历代")
	f2("云南")
	f2("贵州")
	f2("哈哈哈")

	sum := f3(10, 20)
	fmt.Println(sum)

	// 在一个命名函数中不能在声明函数
}
