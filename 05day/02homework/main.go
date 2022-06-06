package main

import (
	"fmt"
	"os"
)

// 学生管理系统
// 保存一些数据
// 三个功能
type Student struct {
	id   int64
	name string
}

// 造一个学生管理者
type StudentMgr struct {
	allStudent map[int64]Student
}

// 构造一个Student对象
func (s StudentMgr) newStudent(name string, id int64) Student {
	return Student{
		name: name,
		id:   id,
	}
}

// 打印学生信息
func (s *Student) info() {
	fmt.Printf("学号：%d 姓名：%v\n", s.id, s.name)
}

// 显示所有学生
func (s StudentMgr) showAllStudent() {
	for _, v := range s.allStudent {
		v.info()
	}
}

// 增加一个学生
func (s StudentMgr) addStudent(name string, id int64) {
	stu := s.newStudent(name, id)
	s.allStudent[id] = stu
	fmt.Printf("学号：%d 姓名：%v 学生信息录入成功!\n", id, name)
}

// 删除一个学生
func (s StudentMgr) deleteStudent(id int64) {
	delete(s.allStudent, id)
	fmt.Printf("学号：%d 学生信息删除成功!\n", id)
}

var smr StudentMgr

// 菜单
func showMenu() {
	fmt.Println("===============================================")
	fmt.Println("================ welcome sms! =================")
	fmt.Println("===============================================")
	fmt.Println(`
	学生信息关系系统菜单如下：
		1. 查看所有学生
		2. 添加学生
		3. 删除学生
		4. 退出
	`)
}

func main() {

	smr = StudentMgr{
		allStudent: make(map[int64]Student, 40),
	}

	for {
		showMenu()
		// 等待用户输入
		var commond int64
		fmt.Print("请输入你的指令：")
		fmt.Scan(&commond)
		fmt.Printf("您选择了 %d 选项!\n", commond)
		fmt.Println("===============================================")
		fmt.Println("==================== End ======================")
		fmt.Println("===============================================")
		fmt.Println()

		switch commond {
		case 1:
			smr.showAllStudent()

		case 2:
			// 输入学号，名称
			var id int64
			var name string
			fmt.Print("请输入学生学号：")
			fmt.Scan(&id)
			fmt.Print("请输入学生姓名：")
			fmt.Scan(&name)
			smr.addStudent(name, id)

		case 3:
			// 输入学号，名称
			var id int64
			fmt.Print("请输入学生学号：")
			fmt.Scan(&id)
			smr.deleteStudent(id)

		case 4:
			os.Exit(1)

		default:
			fmt.Println("滚~")
		}
	}
}
