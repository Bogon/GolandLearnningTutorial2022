package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client

func main() {

	// 1、与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial 127.0.0.1:20000 fialed, error: ", err)
		return
	}
	// 2、发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请说话：")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			break
		}
		conn.Write([]byte(text))
	}
	conn.Close()
}
