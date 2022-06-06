package main

import "fmt"

// 类型断言

func main() {

	var a interface{} // 定义一个 空接口变量

	a = 100
	// 判断 a 保存的值具体的类型
	// 类型断言
	v, ok := a.(int)
	if !ok {
		fmt.Println("猜错了…")
	} else {
		fmt.Println(v)
	}

	fmt.Println(a.(int))

	// switch
	switch a.(type) {
	case int8:
		fmt.Println("int8")
	case int16:
		fmt.Println("int16")
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("滚~")
	}

}
