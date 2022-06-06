package main

import (
	"fmt"
)

func main() {

	// fmt.Print("山河")
	// fmt.Print("美丽")
	// fmt.Println("--------")
	// fmt.Print("山河")
	// fmt.Print("美丽")

	// printf("格式化字符串", 值)
	// %T: 查看类型
	// %d: 十进制数
	// %b: 二进制数
	// %o: 八进制数
	// %x: 十六进制数
	// %c: 显示字符
	// %s: 字符串
	// %p: 指针
	// %v: 值
	// %f: 浮点数
	// %t: 布尔类型

	//

	// 通用占位符
	// %v: 值
	// %#v: 类型+值
	// var m1 = make(map[string]int, 1)
	// m1["lixiang"] = 200
	// fmt.Printf("%v\n", m1)
	// fmt.Printf("%#v\n", m1)

	// printBaiFenBi(90)

	// 获取用户输入
	// var s string
	// fmt.Scan(&s)
	// fmt.Println("用户输入的内容是：", s)

	var (
		name  string
		age   int
		class string
	)

	// fmt.Scanf("%s %d %s", &name, &age, &class)
	// fmt.Println("用户姓名：", name, "用户年龄：", age, "用户班级：", class)

	fmt.Scanln(&name, &age, &class)
	fmt.Println("用户姓名：", name, "用户年龄：", age, "用户班级：", class)
}

func printBaiFenBi(num int) {
	fmt.Printf("%d%%\n", num)
}
