package main

import (
	"fmt"
	"math"
)

// 浮点数
func main() {
	f1 := math.MaxFloat32 // 32位浮点数

	fmt.Println(f1)

	f2 := 1.2342
	fmt.Printf("%T\n", f2) // 默认 go 语言中的小数都是 64位

	f3 := float32(1.2342)
	fmt.Printf("%T\n", f3) // 显式声明 32 位 浮点数
	// f2 = f3 float32 位类型的值不能直接赋值给 float64 位类型的变量
}
