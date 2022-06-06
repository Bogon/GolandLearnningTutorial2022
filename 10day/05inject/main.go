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

// sql 注入
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id,name,age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)

	var users []user
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Println("select users failed, error: %v\n", err)
		return
	}

	for _, u := range users {
		fmt.Printf("%#v\n", u)
	}

}

func main() {

	err := initDB()
	if err != nil {
		fmt.Println("init db failed, error: ", err)
		return
	}

	sqlInjectDemo("李云辉' or 1=1 #")

}
