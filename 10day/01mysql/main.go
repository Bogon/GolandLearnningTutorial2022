package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// go 链接 mysql

func query() {

}

func insert() {

}

func main() {

	// 数据库信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/goday10"
	/// 连接数据库
	db, err := sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {                   // dsn 格式不正确
		fmt.Printf("dsn：%s invilid, error:%v \n", dsn, err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed, error:%v \n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功")

}
