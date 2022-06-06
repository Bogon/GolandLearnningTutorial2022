package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i == 8 {
			continue
		} else if i == 17 {
			fmt.Println("停止打印")
			break
		} else {
			fmt.Println("当前索引：", i)
		}
	}
}
