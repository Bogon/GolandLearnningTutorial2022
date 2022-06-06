package main

import "fmt"

// go 语言声明变量的方式, 推荐使用驼峰式命名
// var student_name string
var studentName string

// var StudentName string

// 声明变量
// var name string
// var age int
// var isShow bool

// 批量声明
var (
	name   string
	age    int
	isShow bool
)

func main() {
	name = "Hello man"
	age = 20
	isShow = false

	var lalalal string

	fmt.Println("Mine name is", name, "；My age is", age, "；My avatar is show or Not:", isShow, "。")

	lalalal = "啦啦啦啦"

	fmt.Printf("lallala is %s", lalalal)

	// 声明变量的同时 并给出赋值
	var s1 string = "s1"
	// 类型推到
	var s2 = "s2"
	// 简短变量声明
	s3 := "s3"

	fmt.Println(s1, s2, s3)
}

/*
备注：
	1、函数外的每一条语句都必须是以关键字开始
	2、:= 简短声明变量不能在函数外使用
	3、_ 多用于占位符，表示忽略值
*/
