package StudentMgr

import "fmt"

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
