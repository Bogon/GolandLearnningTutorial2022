package main

import (
	"fmt"
)

// 类型断言

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了……")
	} else {
		fmt.Println("传进来的是一个字符串：", str)
	}
}

func assign2(a interface{}) {
	fmt.Printf("%T\n", a)

	switch t := a.(type) {
	case string:
		fmt.Println("传进来的是一个字符串：", t)
	case int:
		fmt.Println("传进来的是一个整形：", t)
	case bool:
		fmt.Println("传进来的是一个布尔型：", t)
	default:
		fmt.Println("未能识别的类型：", t)
	}
}

func main() {

	assign2(100)
	assign2("100")
	assign2(false)

	type cat struct{}

	assign2(cat{})

}
