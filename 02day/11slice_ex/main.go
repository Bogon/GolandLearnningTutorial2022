package main

import (
	"fmt"
	"sort"
)

func main() {

	var a = make([]int, 5, 10) // 创建切片 长度为5 容量 为10
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	fmt.Println(cap(a))

	// 排序
	var a1 = [...]int{2, 4, 1, 7, 0}
	sort.Ints(a1[:])
	fmt.Println(a1)

}
