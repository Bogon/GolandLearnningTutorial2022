package main

import "fmt"

// defer
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("哈哈哈哈")
	defer fmt.Println("嘿嘿嘿黑")
	defer fmt.Println("嘻嘻嘻嘻")
	fmt.Println("end")
}

func main() {

	deferDemo()

}
