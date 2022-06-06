package main

import "fmt"

// 全局常量
// 常量定义之后在程序运行期间不能改变
const pi = 3.14

// 批量声明，如果某一行声明没有复制，默认与上一个值一致
const (
	statusOk = 200
	notFound = 404
)

// iota 是常量计数器，只能在常量的表达式中使用
// iota 在const 关键字首次出现的时候会被置为0，const 中每新增一条常量声明，iota 计数一次(iota 可理解为const语句块中行索引)
// 使用iota 可以简化定义，在定义枚举时很有用。

const (
	a1 = iota
	a2
	a3
)

const (
	b1 = iota // 只要const出现，iota就会被置0
	b2
	_
	b3
)

// 插队
const (
	c1 = iota // 只要const出现，iota就会被置0
	c2 = 100
	c3 = iota
	c4
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2

	d3, d4 = iota + 1, iota + 2
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("a1: ", a1)
	fmt.Println("a2: ", a2)
	fmt.Println("a3: ", a3)

	fmt.Println("b1: ", b1)
	fmt.Println("b2: ", b2)
	fmt.Println("b3: ", b3)

	fmt.Println("c1: ", c1)
	fmt.Println("c2: ", c2)
	fmt.Println("c3: ", c3)
	fmt.Println("c4: ", c4)

	fmt.Println("d1: ", d1)
	fmt.Println("d2: ", d2)
	fmt.Println("d3: ", d3)
	fmt.Println("d4: ", d4)
}
