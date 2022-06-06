package main

import (
	"fmt"
	"net"
)

// tcp server

func processConn(cnn net.Conn) {
	defer cnn.Close()
	// 3.与客户端通信
	var tmp [128]byte
	for {
		n, err := cnn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed, err: ", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}

func main() {

	fmt.Println()
	// 1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp serevr on 127.0.0.1:20000 failed, error:", err)
		return
	}
	for {
		// 2.等待别人来链接
		cnn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, error:", err)
			return
		}
		go processConn(cnn)
	}

}
