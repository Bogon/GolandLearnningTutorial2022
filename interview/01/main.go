package main

import "fmt"

func main() {

	// 这道题目的核心点在于对 rune 字面量的理解和数组的语法。
	m := [...]int{
		'a': 1,
		'b': 2,
		'c': 3,
	}

	m['a'] = 3
	fmt.Println(len(m))
	fmt.Println(m)
}
