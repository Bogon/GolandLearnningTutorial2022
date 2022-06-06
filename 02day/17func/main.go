package main

import "fmt"

// 函数
// 函数定义 一段代码的封装，让代码结构清晰和整洁
func sum(x int, y int) (ret int) {
	return x + y
}

// 无返回值
func f2(x int, y int) {
	fmt.Println(x + y)
}

// 不命名返回值
func f3(x int, y int) int {
	return x + y
}

//  返回值可以命名也可以不命名
func f4(x int, y int) (ret int) {
	ret = x + y
	return // 使用命名返回值可以在return 省略
}

// 多个返回值
func f5() (int, string) {
	return 1, "A"
}

//  参数类型简写，当前面几个类型一致时可以只在最后一个参数后面添加类型，前面的类型可以省略
func f6(x, y int) int {
	return x + y
}

// 可变长参数	可变长参数必须放在函数的末尾
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

func main() {

	// 函数定义
	r := sum(10, 20)
	fmt.Println(r)

	fmt.Println(f5)

	f7("2")
	f7("2", 21, 123, 123, 1, 23, 12, 3, 2)

}
