package main

import (
	"bufio"
	"fmt"
	"os"
)

func useBufInfo() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Println(s)
}

func main() {
	// logger.Trace()
	// logger.Debug()
	// logger.Warning()
	// logger.Info()
	// logger.Error("日志的内容")

	// 写日志
	fmt.Fprintln(os.Stdout, "这是一天日志记录")
	file, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	fmt.Fprintln(file, "这是一天日志记录")
}
