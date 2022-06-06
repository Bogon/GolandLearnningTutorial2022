package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 打开文件
func readFile() {
	fileobj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	// 记得关闭文件
	defer fileobj.Close()

	// 读文件
	var tmp = make([]byte, 128)
	for {
		n, err := fileobj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读取完成！")
			return
		}

		if err != nil {
			fmt.Println("read from file failed, err: ", err)
			return
		}
		fmt.Printf("读取了 %d 个字节\n", n)
		fmt.Println(string(tmp))

		if n < 128 {
			return
		}
	}
}

// 读取文件 使用 bufferio读取文件
func readFileByBuffer() {
	fileobj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	// 记得关闭文件
	defer fileobj.Close()
	// 创建一个从文件中读取内容的对象
	reader := bufio.NewReader(fileobj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("read line failed, err:", err)
			return
		}

		fmt.Print(line)
	}
}

func readFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("iouitl read file failed, err:", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	// readFile()

	// readFileByBuffer()

	readFileByIoutil()
}
