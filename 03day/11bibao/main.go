package main

import (
	"fmt"
	"strings"
)

// 闭包
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")

	fmt.Println(jpgFunc("test"))
	fmt.Println(jpgFunc("test.jpg"))
	fmt.Println(jpgFunc("test.jpg.txt"))
	fmt.Println(txtFunc("test"))
	fmt.Println(txtFunc("test.txt"))
	fmt.Println(txtFunc("test.jpg"))

}
