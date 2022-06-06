package main

import (
	"fmt"
)

// 结构体嵌套

type address struct {
	province string
	city     string
}

type compony struct {
	name string
	addr address
}

type person struct {
	name    string
	age     int
	address // 匿名嵌套结构体
	workPlace
}

type workPlace struct {
	province string
	city     string
}

func main() {

	p1 := person{
		name: "李磊",
		age:  20,
		address: address{
			province: "北京",
			city:     "顺义区",
		},
	}

	fmt.Println(p1)
	fmt.Println(p1.name)
	// fmt.Println(p1.province)
	// fmt.Println(p1.city) // 现在自己的结构体中找字段，没有话向匿名结构体中找字段
	fmt.Println(p1.address.city)
	fmt.Println(p1.address.province) // 现在自己的结构体中找字段，没有话向匿名结构体中找字段
	// fmt.Println(p1.addr.province)
	// fmt.Println(p1.addr.city)

}
