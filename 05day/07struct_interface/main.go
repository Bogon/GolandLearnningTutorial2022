package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口还可以嵌套

// 如果一个变量实现了这个接口的所有方法，就说明这个变量实现了所有的接口

// 空接口 咩有方法
// 空接口没有必要写名字，通常写成下面的样子
// interface {}

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat()
}

type cat struct {
	name string
	feet int64
}

func (c *cat) move() {
	fmt.Println("走猫步……")
}

func (c *cat) eat(st string) {
	fmt.Printf("吃%s……", st)
}

func main() {

	fmt.Println()

}
