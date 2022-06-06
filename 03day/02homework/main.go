package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 判断字符串中汉字的数量
	// 难点是判断一个字符是汉字
	s1 := "沙盒Hello"
	// 1、字符串字符
	// 2、判断是否是汉字
	// 3、记录汉字的数量
	count := 0
	for _, value := range s1 {
		fmt.Println(value)
		if unicode.Is(unicode.Han, value) {
			count++
		}
	}
	fmt.Println("汉字的数量：", count)

	// 判断单词出现的次数
	// how do you do
	s := "how do you do"
	// 1、切割到切片中
	s2 := strings.Split(s, " ")
	// 存储到map中
	m1 := make(map[string]int, 10)
	// 累加出现的次数
	for _, w := range s2 {
		// 如果 m1 中不存在，那就存储到m1中，如果存在，就+1
		_, ok := m1[w]
		if ok {
			m1[w]++
		} else {
			m1[w] = 1
		}
	}

	fmt.Println(m1)

	// 回文判断
	// 字符串从左向右读和从右向左读是一样的
	// 上海自来水来自海上
	// 山西运煤车煤运西山
	// 黄山落叶松叶落山黄

	s4 := "上海自来水来自海上"
	r := make([]rune, 0, len(s4))
	for _, c := range s4 {
		r = append(r, c)
	}
	fmt.Println("[]rune = ", r)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println(r[i], "  -  ", r[len(s4)-1-i])
			fmt.Println(i, "  -  ", len(r)-1-i)
			fmt.Println("不是回文!")
			return
		}
	}
	fmt.Println("是回文!")
}
