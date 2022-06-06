package main

import "fmt"

func main() {

	s1 := make([]int, 5, 10)
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d \n", s2, len(s2), cap(s2))

	// 切片的本质
	// 切片就是一个框，框住了一块连续的内存。
	// 切片属于引用类型，真正的数据都保存在底层的数组里。

	// 切片不能直接比较，一个 nil 值的切片并没有底层数组，一个 nil 值的切片的长度和容量都是0，
	// 但是我们不能说一个长度和容量都是0的切片一定是nil。

	// 判断一个切片是否是空的，要使用 len(s) == 0 来判断，不能使用 s == nil 来判断。

	// 切片的赋值和拷贝
	s3 := []int{1, 3, 5}
	s4 := s3
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s3, s4)

	// 切片的遍历
	// 切片的遍历和数组的遍历是一样的
	for i := 0; i < len(s3); i++ {
		fmt.Println("i: ", i, "value: ", s3[i])
	}

	for i, value := range s3 {
		fmt.Println("i: ", i, "value: ", value)
	}

}
