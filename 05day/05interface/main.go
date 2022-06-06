package main

import "fmt"

// 接口的实现

// 接口保存分为两部分：
// 1、值的类型和值本身
// 这样就可以实现接口变量可以保存不同的值
// 动态类型

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

func (c cat) move() {
	fmt.Println("猫动~")
}

func (c cat) eat(st string) {
	fmt.Printf("猫吃%s~\n", st)
}

func main() {
	// 定义一个接口类型的变量
	var a1 animal
	fmt.Printf("%T\n", a1)

	a1 = cat{
		name: "小花",
		feet: 20,
	}
	fmt.Printf("%T\n", a1)
	a1.move()
	a1.eat("鱼")

	a1 = chicken{
		feet: 30,
	}
	fmt.Printf("%T\n", a1)
	a1.move()
	a1.eat("饲料")
}
