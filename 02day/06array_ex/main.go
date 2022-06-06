package main

import "fmt"

// array 数组的练习
// 1、求数组 [1, 3,5 ,7, 8] 所有元素的和
// 2、找出数组中和为指定值的两个元素的下标，比如数组[1, 3, 5, 7, 8] 中找出和为8的两个元素下标分别为(0, 3)和(1, 2)。

func main() {
	// 1
	a1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a1 {
		sum += v
	}
	fmt.Println("sum =", sum)

	// 2
	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			sum1 := a1[i] + a1[j]
			if sum1 == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}
