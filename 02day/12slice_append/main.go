package main

import "fmt"

// 关于 append 删除切片中某个元素
func main() {

	a1 := [...]int{1, 2, 3, 6, 31, 1}
	s1 := a1[:]
	fmt.Println(s1)

	// 删除 索引为1的2
	s1 = append(s1[0:1], s1[2:]...)
	fmt.Println(s1)
	fmt.Println(a1)

}
