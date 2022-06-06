package main

import "fmt"

// fmt 占位符
func main() {

	var n = 100

	// 打印类型
	fmt.Printf("%T\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	var s = "1313"
	fmt.Printf("%s\n", s)
	fmt.Printf("%#v\n", s)
}
