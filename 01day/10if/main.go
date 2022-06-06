package main

import "fmt"

// if 条件判断
func main() {
	age := 10

	if age > 10 {
		fmt.Println("年龄大于", age, "岁！")
	} else if age == 10 {
		fmt.Println("年龄等于", age, "岁！")
	} else {
		fmt.Println("年龄小于", age, "岁！")
	}

}
