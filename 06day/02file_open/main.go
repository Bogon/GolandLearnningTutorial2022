package main

import (
	"fmt"
	"os"
)

// 文件操作

func openFile() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}
	defer file.Close()

}

func openFile2() {
	file, err := os.OpenFile("./sb.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}
	defer file.Close()
	file.WriteString("a")
	file.WriteString("b")
	file.WriteString("1G")

	file.Seek(1, 0) // 光标移动
	var ret [1]byte

	_, err = file.Read(ret[:])
	if err != nil {
		fmt.Println("read file failed, err: ", err)
		return
	}
	fmt.Printf("读取到的内容：%s", string(ret[:]))

}

func main() {

	openFile2()

}
