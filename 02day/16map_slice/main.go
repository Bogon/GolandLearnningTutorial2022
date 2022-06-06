package main

import "fmt"

func main() {

	// 元素类型为 map 的切片
	// panic: runtime error: index out of range [0] with length 0
	var s1 = make([]map[int]string, 1, 10) // panic: assignment to entry in nil map
	s1[0] = make(map[int]string, 1)
	s1[0][100] = "A"
	fmt.Println(s1)

	// 值为切片类型的 map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = make([]int, 10, 10)
	fmt.Println(m1)
}
