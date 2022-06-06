package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p person
	p.name = "lilei "
	p.age = 20
	p.gender = "男"
	p.hobby = []string{"游泳", "打架", "飞行"}
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)

	// 匿名结构体 多用于临时场景
	var s struct {
		x string
		y int
	}

	s.x = "嘿嘿嘿"
	s.y = 120
	fmt.Println(s)
	fmt.Printf("%T\n", s)

}
