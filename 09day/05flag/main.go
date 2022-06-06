package main

import (
	"flag"
	"fmt"
	"time"
)

// 获取命令行参数
func main() {

	// 创建一个标志位函数
	name := flag.String("name", "洒水", "请输入一个名字")
	age := flag.Int("age", 20, "请输入真实年龄")
	married := flag.Bool("married", false, "请输入是否结婚")
	cTime := flag.Duration("ct", time.Second, "结婚多久了")

	// var name string
	// flag.StringVar(&name, "name", "洒水", "请输入一个名字")

	// 使用flag
	flag.Parse()
	// fmt.Println(name)
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)

	fmt.Println(flag.Args())  // 返回命令行参数后的其他参数，以[]string 类型
	fmt.Println(flag.NArg())  // 返回命令行后的其他参数个数
	fmt.Println(flag.NFlag()) // 返回使用的命令行参数个数

}
