package calc

import "fmt"

func init() {
	fmt.Println("我是init函数，自动执行……")
}

// 包中的标识符(变量名，函数名，结构体接口等)， 小写的表示私有只能在当前包中使用
// 首字母大写的可以被外部可见
func Add(x, y int) (ret int) {
	ret = x + y
	return
}
