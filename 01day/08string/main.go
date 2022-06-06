package main

import (
	"fmt"
	"strings"
)

// 字符串

func main() {
	path := "Users/zhangqi/go/src/learning.goland.com/studyGoland/01day/08string"
	fmt.Println(path)

	// 字符使用单引号
	s := "i'm ok"
	fmt.Println(s)

	// 多行字符串
	s2 := `
			世情薄
			人情恶
			雨送黄昏花易落
		  `
	fmt.Println(s2)

	// 字符串拼接
	s3 := "立项"
	s4 := "sdd"

	ss := s3 + s4
	fmt.Println(ss)

	ss1 := fmt.Sprintf("%s%s", s3, s4)
	fmt.Println(ss1)

	fmt.Printf("%s%s\n", s3, s4)

	// 字符串分割
	ret := strings.Split(path, "/")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(s3, "立"))
	fmt.Println(strings.Contains(s4, "s"))

	// 前缀
	fmt.Println(strings.HasPrefix(path, "Users"))
	// 后缀
	fmt.Println(strings.HasSuffix(path, "string"))

	// 获取索引
	fmt.Println(strings.Index(path, "s"))
	fmt.Println(strings.LastIndex(path, "s"))

	// 拼接字符串
	fmt.Println(strings.Join(ret, "+"))

}
