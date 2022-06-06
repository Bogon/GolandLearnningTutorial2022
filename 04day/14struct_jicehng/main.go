package main

import "fmt"

// 结构体中模拟实现其他语言中的“继承”
type animal struct {
	name string
}

// move
func (a *animal) move() {
	fmt.Println(a.name + "开始移动~")
}

type dog struct {
	feet uint8
	animal
}

func (d *dog) wang() {
	fmt.Printf("%v, 狂吠了……", d.name)
}

func main() {

	// a1 := animal{
	// 	name: "小花",
	// }
	// a1.move()

	d1 := dog{
		animal: animal{
			name: "哈雷",
		},
	}
	d1.wang()
	d1.move()

}
