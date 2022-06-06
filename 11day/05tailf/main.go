package main

import (
	"fmt"
	"time"

	"github.com/nxadm/tail"
)

// tailf 用法
func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的那个位置打开
		MustExist: false,                                // 文件不存在报错
		Poll:      true,                                 //
	}

	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, error: ", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)

	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail file close reopen, Filename", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)

	}
}
