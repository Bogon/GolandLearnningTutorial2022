package main

import (
	"fmt"
)

// 切片 slice

func main() {
	// 切片的定义
	var s1 []int // 定义一个存在 int 类型的切片
	var s2 []string

	fmt.Printf("%T, %T\n", s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"1", "2", "3"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 长度和容量
	fmt.Printf("len s1: %d, cap s1: %d\n", len(s1), cap(s1))
	fmt.Printf("len s2: %d, cap s2: %d\n", len(s2), cap(s2))

	// 通过数组获取切片
	// 切片第一个元素到最后一个元素的数量 就是切片的容量
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 6, 5, 43, 5, 3, 3, 2, 3, 42}
	s3 := a1[0:4] // 左闭右开
	fmt.Println(s3)
	s4 := a1[1:7] // 左闭右开
	fmt.Println(s4)
	fmt.Printf("len s4: %d, cap s4: %d\n", len(s4), cap(s4))
	s5 := a1[:7] // 左闭右开
	fmt.Printf("len s5: %d, cap s5: %d\n", len(s5), cap(s5))
	fmt.Println(s5)
	s6 := a1[1:] // 左闭右开
	fmt.Printf("len s6: %d, cap s6: %d\n", len(s6), cap(s6))
	fmt.Println(s6)
	s7 := a1[:] // 左闭右开
	fmt.Printf("len s7: %d, cap s7: %d\n", len(s7), cap(s7))
	fmt.Println(s7)

	// 切片
	// 切片底层是指向一个数组
	// 切片的长度就是切片中元素的个数
	// 切片的容量就是切片指向数组中的第一个元素到最后一个元素的个数

	// 切片是引用类型，都指向了底层的一个数组
	// 切片再切割
	// 修改底层数组的值
	a1[3] = 1500
	s8 := s3[2:]
	fmt.Println(s8)
	fmt.Println(a1)
	fmt.Printf("len s8: %d, cap s8: %d\n", len(s8), cap(s8))

	// make 函数创建切片

}
