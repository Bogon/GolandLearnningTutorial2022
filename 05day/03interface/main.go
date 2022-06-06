package main

import "fmt"

// 引出接口的实例

// 定义一个能叫的类型
type speaker interface {
	speak() // 只要实现了 sperk 方法的变量都是speaker类型 方法的签名
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func (p person) speak() {
	fmt.Println("啊啊啊啊~")
}

func da(x speaker) {
	x.speak() // 只要实现了 sperk 方法的变量都是speaker类型
}

func main() {
	var c cat
	var d dog
	var p person

	da(c)
	da(d)
	da(p)
}
