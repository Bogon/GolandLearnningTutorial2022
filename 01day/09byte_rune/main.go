package main

import "fmt"

func main() {
	// 字符串长度
	s1 := "Hello，中国"
	len := len(s1)
	fmt.Println(len)

	// 字符串修改
	s2 := "西红柿"
	s3 := []rune(s2) // 把字符串强制转换成 rune 字符串
	s3[0] = '蓝'
	fmt.Println(string(s3)) // 把rune切片强制转换成字符串

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1: %T, c2: %T\n", c1, c2)

	c3 := "H"
	c4 := 'H'
	fmt.Printf("c3: %T, c4: %T\n", c3, c4)
	fmt.Printf("c4: %d\n", c4)

	// 类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)
}
