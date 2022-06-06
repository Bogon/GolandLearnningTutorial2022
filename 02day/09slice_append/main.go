package main

import "fmt"

// append 为切片添加元素
func main() {
	s1 := []string{"北京", "上海", "广州"}
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
	s1[2] = "深圳" // 编译错误，索引越界
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))

	// 调用 append 函数必须用原来的切片变量接收返回值
	s1 = append(s1, "郑州")
	// append 最近元素，原来底层的数组放不下的时候，go底层会把底层数组换一个
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))

	s2 := []string{"西安", "武汉", "南京"}
	s1 = append(s1, s2...) // ... 表示拆开
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
}
