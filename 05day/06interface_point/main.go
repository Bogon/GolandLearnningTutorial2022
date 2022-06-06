package main

import "fmt"

// 使用值接收者和指针接收者的区别
// 1、使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存储
// 2、指针接收者实现接口只能存结构体指针类型的变量

// 类型和接口的关系
// 多个类型可以实现一个接口
// 一个类型可以实现多个接口

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int64
}

type chicken struct {
	feet int64
}

func (c chicken) move() {
	fmt.Println("鸡动~")
}

func (c chicken) eat(st string) {
	fmt.Printf("鸡吃%s~\n", st)
}

// 使用值接收者实现的接口的方法
func (c cat) move() {
	fmt.Println("猫动~")
}

func (c cat) eat(st string) {
	fmt.Printf("猫吃%s~\n", st)
}

func main() {

	var a1 animal

	c1 := cat{
		name: "小金桔",
		feet: 15,
	}
	c2 := &cat{
		name: "假老练",
		feet: 15,
	}
	fmt.Println(c1, c2, a1)

	a1 = c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)

}
