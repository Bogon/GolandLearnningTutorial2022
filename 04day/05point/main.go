package main

import "fmt"

func main() {

	var n int
	n = 100

	m := &n
	fmt.Printf("type n: %T, type m: %T\n", n, m)
	//  打印内存地址
	fmt.Printf("%p\n", &n)
	fmt.Printf("%p\n", m)
	fmt.Printf("%p\n", &m)
	fmt.Printf("%v\n", m)
}
