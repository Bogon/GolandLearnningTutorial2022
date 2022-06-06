package main

import "fmt"

func main() {
	var name string
	name = "123"
	fmt.Println(name)

	// 数组初始化
	var ages [20]int // 声明了一个变量，它是 [20]int 类型
	ages = [20]int{1, 213, 12, 3, 123}
	fmt.Println(ages)

	var age2 = [...]int{1, 2, 123, 123, 1, 23, 1, 23, 1, 2}
	fmt.Println(age2)

	var age3 = [...]int{1: 100, 99: 200}
	fmt.Println(age3)

	// 数组的遍历
	// 二维数组的定义
	var a1 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a1)
	// 多维数组只有最外层可以使用 ...

	// 数组是值类型
	x := [3]int{1, 2, 3}
	fmt.Println(x)
	y := x
	y[1] = 100
	fmt.Println(x)
	fmt.Println(y) // 修改副本，不用不影响原来的值
	change(x)
	fmt.Println(x)

	// 切片的定义 slice
	var s1 []int // 切片声明之后是没有分配内存的，使用之前需要分配内存
	fmt.Println(s1)
	fmt.Println(s1 == nil)

	s1 = []int{1, 2, 3, 4, 2, 2}
	fmt.Println(s1)

	// make 初始化并分配内存
	s2 := make([]bool, 2, 4)
	fmt.Println(s2)

	s3 := make([]int, 0, 4)
	fmt.Println(s3)
	fmt.Println(s3 == nil)

	// 切片的定义
	// 切片：指针、长度、容量

	s4 := []int{1, 2, 3, 4}
	s4[2] = 100
	s5 := s4
	fmt.Println(s4)
	fmt.Println(s5)

	// 切片的扩容策略
	// 1、如果申请的容量大于原来的2倍，直接扩容至新申请的容量
	// 2、如果小于1024，那么直接2倍
	// 3、如果大于1024，就按照1.25倍扩容
	// 4、具体存储的值类型不同，扩容的策略也会有所调整。

	var s6 []int
	s6 = make([]int, 1)
	s6[0] = 100
	fmt.Println(s6)

	s7 := []int{1, 2, 3}
	s8 := s7
	var s9 []int
	copy(s9, s8)
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)

	// 指针
	// go语言中指针，只能读不能修改，不能修改指针变量指向的地址
	addr := "云南"
	addrP := &addr
	fmt.Println(addrP)
	fmt.Printf("%T \n", addrP)
	addrV := *addrP
	fmt.Println(addrV)

	// map
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["lilei"] = 120
	fmt.Println(m1)
	fmt.Println(m1["lilei"])
	fmt.Println(m1["lilei2"]) // 如果不存在，返回对应类型的 0 值

	score, ok := m1["lilei2"]
	if ok {
		fmt.Println(score)
	} else {
		fmt.Println("查无此key！")
	}
	delete(m1, "klilei") // 如果删除的key不存在，则什么都不做
	delete(m1, "lilei")
	fmt.Println(m1)
	fmt.Println(m1 == nil)
}

// func
func change(a [3]int) {
	a[1] = 100
}
