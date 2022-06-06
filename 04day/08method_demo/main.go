package main

import (
	"fmt"
)

type person struct {
	name   string
	age    int
	gender string
}

// 方法
// 标识符：变量名、函数名、类型名、方法名
// go语言中如果首字母是大写的，就表示对外部可见(暴露的，公有的)
type dog struct {
	name string
}

// 构造函数
func newdog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法式作用域特定类型的函数
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("dog.name: %v 旺~\n", d.name)
}

func newPerson(name string, age int, gender string) person {
	return person{
		name:   name,
		age:    age,
		gender: gender,
	}
}

// 指针接收者：传递内存地址
func (p *person) guonian() {
	p.age++
}

func (p *person) dream() {
	fmt.Println("I无好丢为第五二号")
}

// 什么时机应该传递指针接收者
// 1、需要修改接收者中的值
// 2、接收者是拷贝代价比较大的对象
// 3、保证一致性，如果某个方法使用了指针接收者，那么其他的犯非法也应该使用指针接收者

func main() {
	// d1 := newdog("泰迪")
	// d1.wang()

	p1 := newPerson("王磊", 18, "男")
	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)
	p1.dream()

}
