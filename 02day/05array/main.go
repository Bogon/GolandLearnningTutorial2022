package main

import "fmt"

// 数组

// 存放元素的容器
// 必须指定存放的元素类型和容量(长度)
// 数组的长度是类型的一部分
func main() {
	var a1 [3]bool
	var a2 [4]bool

	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)

	// 数组的初始化
	// 如果不初始化，默认都是 0值 (布尔值是：false，整型和浮点型都是0，字符串："")
	fmt.Println(a1, a2)
	// 初始化方式1
	a1 = [3]bool{true, false, true}
	fmt.Println(a1)
	// 初始化方式2
	a10 := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a10)
	// 根据初始值自动推断数组的长度是多少
	a11 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 8, 67, 5, 3, 2, 12, 4, 56, 7, 4}
	fmt.Println(a11)
	// 初始化方式3 根据索引初始化值
	a13 := [5]int{0: 1, 4: 2}
	fmt.Println(a13)

	// 数组的遍历
	// 1、根据索引进行遍历
	citys := [...]string{"北京", "洛阳", "郑州", "西安", "重庆"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	//2、for range 遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 多维数组
	// [[1,2], [3, 4]]
	var a14 [3][2]int
	a14 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a14)

	// 多维数组的遍历
	for _, v1 := range a14 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 4
	fmt.Println(b1, b2)

}
