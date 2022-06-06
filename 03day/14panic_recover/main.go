package main

import "fmt"

func funcA() {
	fmt.Println("A")
}

func funcB() {
	// 刚刚打开数据库
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("打开数据库链接")
	}()
	panic("出现严重的错误啦！！！") // 程序崩溃退出
	fmt.Println("B")
}

func funcC() {
	fmt.Println("C")
}

func main() {
	funcA()
	funcB()
	funcC()
}
