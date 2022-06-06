package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("当前位置：%d, %d \n", i, j)
			if j == 2 {
				break /// 仅仅跳出本次循环
			} else if j == 5 {
				// 设置退出标签 跳出所有的嵌套的循环
				goto breakTag
			}
		}
	}

	return

breakTag:
	fmt.Println("结束for 循环")
}
