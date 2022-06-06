package main

import (
	"fmt"
)

// 指针
func main() {

	// 1、&: 取地址
	// 2、*: 根据地址取值

	n := 18
	p := &n
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Printf("%T \n", p)

	// 取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T \n", m)

	// var a *int // nil pointer
	var a = new(int)
	fmt.Println(a)
	*a = 100
	fmt.Println(*a)

	var a1 *int
	fmt.Println(a1)

	// make 和 new 的区别
	// 1、make 和 new 都是用来申请内存的
	// 2、new 很少用，一般用来给基本数据类型申请内存，string/int...，返回对应类型的指针。
	// 3、make 用来给 slice、map、chan 申请内存，make函数返回的是对应这三个类型本身的指针。

}
