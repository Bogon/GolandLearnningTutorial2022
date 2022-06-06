package main

import (
	"fmt"

	calc "learning.goland.com/studyGoland/05day/10calc"
)

var x = 100

const pi = 3.14

func init() {
	fmt.Println("自动执行……")
	fmt.Println(x, pi)
}

// 起别名 calc "learning.goland.com/studyGoland/05day/10calc"
// 默认文件件的名字就是报名

// 包中init函数的执行时机
// 全局声明 => init() => main()

// 包的路径从 GOPATH/src  后面的路径开始解析
// 对外的像被别的包调用，首字母需要大写
// 单行导入和多行导入
// 导入包的时候可以指定别名
// 导入包不想使用内部的标识符，需要使用匿名导入
// 每个包导入的时候都会执行定义的 init 函数，没有参数、返回值
// 多个包定义了 init函数，调用init函数按照导入包的逆序进行调用

func main() {
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}
