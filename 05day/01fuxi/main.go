package main

import (
	"encoding/json"
	"fmt"
)

// 复习结构体
type tmp struct {
	x int
	y int
}

type person struct {
	name string
	age  int
}

// 构造(结构体变量的)函数，返回值是对应结构体类型
func newPerson(name string, age int) person {
	return person{name: name, age: age}
}

// 方法和接收者
// 方法 = 函数 + 接收者
// 制定了接收者，只能该类型的变量才可以调用
// 接收者一般使用类型的第一个字母小写
func (p *person) dream() {
	fmt.Printf("%v 在白日做梦~\n", p.name)
}

// 结构体是一个值类型
// 值类型在函数传递过程中，传递的是副本

// 指针接收者
// 1、需要修改结构体变量的值
// 2、结构体比较大，传递过程中开销较大
// 3、方法中有一个使用了指针接收者，后面的方法最好支持
func (p *person) guonian() {
	p.age++
	fmt.Printf("%v 在过年，年龄是：…%d~\n", p.name, p.age)
}

func main() {

	var a = tmp{
		10, 20,
	}
	fmt.Println(a)

	var b = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(b)

	var p1 person
	p1.name = "运单"
	p1.age = 40
	p2 := person{"李磊", 20} // 结构体实例化
	fmt.Println(p1, p2)

	// 调用构造函数生成person类型的变量
	p3 := newPerson("hahah ", 20)
	fmt.Println(p3)
	p3.dream()
	p3.guonian()

	// 结构体嵌套
	type address struct {
		province string
		city     string
	}

	type student struct {
		name    string
		address // 匿名 嵌套结构体
	}

	stu1 := student{
		name: "徐骁",
		address: address{
			province: "四川",
			city:     "成都",
		},
	}

	fmt.Println(stu1)
	fmt.Println(stu1.province) // 现在内部查找变量、内部没有从匿名结构体中查找变量

	// 序列化和反序列化
	// 1、结构体中变量不要大写!!!不大写别人就访问不到
	type point struct {
		X int `json:"hengzou"`
		Y int `json:"zouzhou"`
	}

	poi1 := point{
		X: 100,
		Y: 200,
	}
	j_v, err := json.Marshal(poi1)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	fmt.Printf("%v\n", string(j_v))

	// 反序列化
	// 反序列化时要传递指针
	var poi2 point
	po_v := `{"hengzou": 120, "zouzhou": 500}`
	err = json.Unmarshal([]byte(po_v), &poi2)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
		return
	}
	fmt.Println(poi2)

	type myInt int
	type mineInt int
	// 自定义类型和类型别名只在编写是有效，编译完成之后就不存在，内置的类型byte和rune都属于类型别名

}
