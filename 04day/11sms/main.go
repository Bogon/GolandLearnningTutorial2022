package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看\新增\删除学生
*/

var (
	allStudent map[int64]*student
)

type student struct {
	id   int64
	name string
}

// 展示所有学生
func showAllStudent() {
	// 把所有学生打印出来
	for k, v := range allStudent {
		fmt.Printf("学号：%d, 姓名：%v\n", k, v.name)
	}
}

// 新增学生
func addStudent() {
	// 添加一个学生
	// 获取用户输入
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号：")
	fmt.Scan(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scan(&name)

	s := student{
		name: name,
		id:   id,
	}
	allStudent[id] = &s
	fmt.Printf("学号：%d, 姓名：%v 学生信息录入成功!\n", id, name)
}

// 删除学生
func deleteStudent() {
	var id int64
	fmt.Print("请输入要删除学生学号：")
	fmt.Scan(&id)
	delete(allStudent, id)
	fmt.Printf("学号：%d  学生信息删除成功!\n", id)
}

func main() {
	allStudent = make(map[int64]*student, 40) // 初始化（开辟内存空间）
	for {
		// 1、打印菜单
		fmt.Println("欢迎光临学生管理系统！")
		fmt.Println(`
			1. 查看所有学生
			2. 新增学生
			3. 删除学生
			4. 退出
		`)
		fmt.Print("请输入你要干什么：")
		// 2、等待客户选择做什么
		var commond int
		fmt.Scan(&commond)
		fmt.Printf("你选择了%d这个选项！\n", commond)
		// 3、执行对应的函数
		switch commond {
		case 1:
			showAllStudent()

		case 2:
			addStudent()

		case 3:
			deleteStudent()

		case 4:
			os.Exit(1)

		default:
			fmt.Println("请输入有效指令!")
		}
	}
}
