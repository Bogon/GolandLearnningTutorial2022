package main

import (
	"fmt"
	"time"
)

func f1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 时间戳
	fmt.Println(now.Unix())
	// 把时间戳转换成时间
	ret := time.Unix(1652253557, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Month())
	fmt.Println(ret.Day())
	fmt.Println(ret.Hour())
	fmt.Println(ret.Minute())
	fmt.Println(ret.Second())

	/// 时间间隔
	fmt.Println(time.Second)
	// now + 1小时
	fmt.Println(now.Add(24 * time.Hour))
	// 时间间隔
	newTime, err := time.Parse("2006-01-02", "2022-05-12")
	if err != nil {
		fmt.Println("parse time failed， err:", err)
		return
	}
	newTime = newTime.UTC()
	fmt.Println(newTime.Sub(now))

	// 定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	// 时间格式化
	// go 语言诞生时间 2006-1-2 3-4-5
	fmt.Println(now.Format("2006-1-2"))
	// 2019/02/03 11:21:12
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	// 2019/02/03 11:21:12 AM
	fmt.Println(now.Format("2006/01/02 3:04:05 PM"))
	// 2019/02/03 11:21:12.000
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))

	/// 按照对用的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2000-08-02")
	if err != nil {
		fmt.Println("parse time failed， err:", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	// sleep
	n := 5
	fmt.Println("开始sleep了")
	fmt.Println(time.Now())
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(time.Now())
	fmt.Println("5s过去了…")
}

// 时区
func f2() {
	now := time.Now()
	fmt.Println(now)

	// 明天的这时间

}

func main() {

	// f1()

}
