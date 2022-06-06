package main

import "fmt"

// 自定义类型和类型别名

// type 后面是类型
type myInt int

func main() {

	var n myInt
	n = 100
	fmt.Printf("%T\n", n)
	fmt.Println(n)
}
