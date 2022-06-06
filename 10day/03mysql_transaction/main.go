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

// 事务
func transactionDemo() {
	// 1.开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("open transaction failed, error: %v\n", err)
		return
	}
	// 执行多个sql
	sqlStr1 := "update user set age=age-2 where id=2"
	sqlStr2 := "update user set age=age+2 where id=3"
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行 sql1 错误，要执行回滚操作，error：", err)
		return
	}

	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("执行 sql2 错误，要执行回滚操作，error：", err)
		return
	}
	// 上面执行完成，提交事务
	err = tx.Commit()
	if err != nil {
		fmt.Println("事务提交失败，error：", err)
		return
	}
	fmt.Println("事务提交成功")
}

func main() {
	initDB()
	transactionDemo()
}
