package main

import "fmt"

// 构造函数
type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

// 构造函数
// 构造函数返回的结构体还是结构体指针
// 结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newDog(name string) *dog {
	return &dog{
		name: name,
	}
}

func main() {

	p1 := newPerson("提莫", 18)
	p2 := newPerson("卡萨丁", 20)
	fmt.Println(p1, p2)

	d1 := newDog("米粒")
	fmt.Println(d1)

}
