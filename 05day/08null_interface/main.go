package main

import (
	"fmt"
)

// 空接口
// interface: 关键字
// interface{}: 空接口类型

// 空接口可以存储任意类型的值

func show(i interface{}) {
	fmt.Printf("type: %T value: %v\n", i, i)
}

func main() {
	m1 := make(map[string]interface{}, 16)
	fmt.Println(m1)
	m1["name"] = "nicke"
	m1["age"] = 20
	m1["hobbay"] = []int{1, 2, 3, 3}
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)

}
