package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// go 链接 mysql

var db *sql.DB // 一个连接池对象

func initDB() (err error) {
	// 数据库信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	/// 连接数据库
	db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {                  // dsn 格式不正确
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}
	// 设置数据库连接池最大连接数
	db.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)
	return nil
}

type user struct {
	id   int
	name string
	age  int
}

// 单行查询
func queryOne(id int) {
	var u1 user
	// 1、单条记录查询的sql
	sqlStr := "select id, name, age from user where id=?;"
	// 2、执行
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age) // 从连接池里拿出一个连接数据库查询单条记录
	// 3、拿到结果
	// rowObj.Scan(&u1.id, &u1.name, &u1.age) // 必须对 rowObj 对象调用scan方法，这个方法有自动释放连接的方法

	fmt.Printf("u1:%#v \n", u1)
}

// 多行查询
func queryMore(n int) {
	// 1. sql语句
	sqlStr := "select id, name, age from user where id > ?;"
	// 2. 执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s failed, error: %v \n", sqlStr, err)
		return
	}
	// 一定不要忘记关闭
	defer rows.Close()

	// 循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age) // 从连接池里拿出一个连接数据库查询单条记录
		if err != nil {
			fmt.Println("scan failed, error: ", err)
			return
		}
		fmt.Printf("u1:%#v \n", u1)
	}
}

//  插入数据
func insert() {
	sqlStr := `insert into user(name, age) values("张巡礼", 30)`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("exec insert %s failed, error: %v", sqlStr, err)
		return
	}
	// 如果是插入数据，能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get insert id failed, error: %v", err)
		return
	}
	fmt.Printf("insert succesed id: %d \n", id)
}

// 更新数据
func update() {
	sqlStr := "update user set age = ? where id = ?"
	ret, err := db.Exec(sqlStr, 21, 2)
	if err != nil {
		fmt.Printf("update user %s failed, error: %v\n", sqlStr, err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Println("get affect ids failed, error:", err)
		return
	}
	fmt.Printf("修改完成记录：%d 条.\n", n)
}

// 删除数据
func delete(id int) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("update user %s failed, error: %v\n", sqlStr, err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Println("get affect ids failed, error:", err)
		return
	}
	fmt.Printf("操作完成影响记录：%d 条.\n", n)
}

// 预处理插入
func prepareInsert() {
	sqlStr := "insert into user(name, age) values(?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare sql %s failed, error:%v\n", sqlStr, err)
		return
	}
	defer stmt.Close()
	var m = map[string]int{
		"李欣欣": 19,
		"去田林": 20,
		"刘兰兰": 26,
		"王孙兰": 20,
		"李晓磊": 23,
		"王大拿": 50,
	}

	for k, v := range m {
		stmt.Exec(k, v)
	}

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("init db failed, err:", err)
		return
	}
	fmt.Println("连接数据库成功")

	// 单行查询
	// queryOne(1)
	// queryOne(2)

	// 多行查询
	// queryMore(0)

	// 插入
	// insert()

	// queryMore(0)
	// update()
	// queryMore(0)

	// queryMore(0)
	// delete(1)
	// queryMore(0)

	queryMore(0)
	prepareInsert()
	queryMore(0)
}
