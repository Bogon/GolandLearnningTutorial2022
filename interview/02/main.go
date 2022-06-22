package main

import "fmt"

func main() {
	// cannot use num2 (variable of type []int) as type []interface{} in argument to append
	/*
		golang 的类型系统规定 interface{} 是空接口, 任何类型的对象都默认实现了空接口;
		注意，[]interface{} 和 interface{} 是不同的类型, 所以在传参时并不会做类似的隐式类型转换.

		// 数组的类型是类型加上长度标识一个数组类型

		// 新术语：协变与逆变
			子类型可以隐性的转换为父类型
	*/
	var num1 []interface{}
	num2 := []int{1, 2, 3}
	num3 := append(num1, num2...)
	fmt.Println(len(num3))
}
