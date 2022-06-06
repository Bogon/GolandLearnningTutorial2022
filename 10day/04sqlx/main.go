package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	sqlx "github.com/jmoiron/sqlx"
)

// go 链接 mysql

var db *sqlx.DB // 一个连接池对象

func initDB() (err error) {
	// 数据库信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	/// 连接数据库
	db, err = sqlx.Connect("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {                      // dsn 格式不正确
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
	ID   int
	Name string
	Age  int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("init db failed, error: ", err)
		return
	}

	sqlStr := `select id, name, age from user where id=3;`
	var u1 user
	db.Get(&u1, sqlStr)
	fmt.Printf("%#v\n", u1)

	var userList = make([]user, 0, 10)
	sqlStr2 := "select id, name, age from user;"
	db.Select(&userList, sqlStr2)
	fmt.Printf("%#v\n", userList)

}
