package main

import (
	"fmt"
	"os"
)

// os.Args 获取命令行参数

func main() {

	fmt.Printf("type: %T, value: %#v", os.Args, os.Args)

}
