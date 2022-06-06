package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("Hello，哔哔")
}

func f2() int {
	fmt.Println("Hello，哔哔")
	return 10
}

// 函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

// 函数也可以作为返回值
func f4() func() int {
	return f2
}

func main() {

	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)

	f3(f2)

	c := f4()
	fmt.Printf("%T\n", c)
	fmt.Println(c())

}
