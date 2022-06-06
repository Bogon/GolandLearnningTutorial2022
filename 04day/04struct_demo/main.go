package main

import "fmt"

// 结构体是值类型
type person struct {
	name   string
	gender string
}

// 在go语言中，传递传递的永远是副本
func f(x person) {
	x.gender = "女"
}

func f2(x *person) {
	// (*x).gender = "女" // 根据内存地址获取对象，修改数据
	x.gender = "女" // 语法糖，可根据指针找对应的变量
}

func main() {

	var p person
	p.name = "lilei"
	p.gender = "男"
	f(p)
	fmt.Println(p)
	fmt.Println(p.gender)
	f2(&p)
	fmt.Println(p)
	fmt.Println(p.gender)

	// new 返回的该类型的指针
	var p2 = new(person)
	p2.name = "你咩"
	p2.gender = "保密"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)
	fmt.Printf("%p\n", &p2)
	fmt.Println(p2)

	// 2、结构体指针2
	// 2.1、key-value 形式初始化
	// 2.2、值列表的形式进行初始化，值得顺序要和字段定义的顺序一直
	var p3 = person{
		name:   "和黑",
		gender: "男",
	}
	fmt.Printf("%v\n", p3)

	p4 := person{
		"哈哈", "女",
	}
	fmt.Printf("%v\n", p4)

}
