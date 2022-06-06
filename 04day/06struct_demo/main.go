package main

import "fmt"

// 3、结构体占用一块连续的内存空间
type x struct {
	a int8 // 8bit -> 1byte
	b int8
	c int8
}

func main() {

	a := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}

	fmt.Printf("%p\n", &(a.a))
	fmt.Printf("%p\n", &(a.b))
	fmt.Printf("%p\n", &(a.c))
}
