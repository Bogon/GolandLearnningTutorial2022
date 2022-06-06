package main

import "fmt"

// 匿名字段
// 结构体的字段是唯一的
// 字段少，比较简单的场景
// 不常用！！！
type person struct {
	string
	int
}

func main() {

	p1 := person{
		"李磊",
		1230,
	}

	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)

}
