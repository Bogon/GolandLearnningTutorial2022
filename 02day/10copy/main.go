package main

import "fmt"

func main() {

	a1 := []int{1, 2, 3}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)

	a1[1] = 100
	fmt.Println(a1, a2, a3)

	//  删除元素
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1, a2, a3)
}
