package main

import "fmt"

// 接口实例2
// 不管是什么类型的车，都可以跑
type car interface {
	run()
}

func draw(c car) {
	c.run()
}

// byd
type byd struct{}

func (b byd) run() {
	fmt.Println("路上跑的是比亚迪！")
}

// bmw
type bmw struct{}

func (b bmw) run() {
	fmt.Println("路上跑的是宝马！")
}

func runOnRoad(c car) {
	c.run()
}

func main() {

	var b byd
	runOnRoad(b)

	var b1 bmw
	runOnRoad(b1)

}
