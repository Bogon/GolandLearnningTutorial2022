package main

import "fmt"

func main() {

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("当前索引：", i)
	// }

	// var i = 0
	// for i < 10 {
	// 	fmt.Println("1当前索引：", i)
	// 	i++
	// }

	// // 无线循环
	// for i < 10 {
	// 	fmt.Println("1当前索引：", i)
	// }

	// for range 循环
	var value = "Hello，上海！"
	for index, v := range value {
		fmt.Printf("字符串索引：%d, 字符串值：%c \n", index, v)
	}
}
