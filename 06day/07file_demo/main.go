package main

import (
	"fmt"
	"os"
)

// 1、文件对象的类型
// 2、获取文件对象的详细信息

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	fmt.Printf("%T\n", file)
	// 获取文件对象的详细信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file info  failed, err:", err)
		return
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())

}
