package main

import (
	"fmt"
	"strconv"
)

//strconv

func main() {
	str := "1200"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parse int error, ", err)
		return
	}
	fmt.Printf("%#v, %T\n", ret1, ret1)

	i := int32(97)
	ret2 := int64(i)

	fmt.Println(string(ret2))

	// 字符串转整形
	fmt.Println(strconv.Atoi(str))

	// 整形转字符串
	fmt.Println(strconv.Itoa(97))

	// 字符串解析出boo
	b_str := "false"
	ret_bool, _ := strconv.ParseBool(b_str)
	fmt.Println(ret_bool)
}
