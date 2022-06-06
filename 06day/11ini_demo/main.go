package main

import "fmt"

// MysqlConfig mysql 配置文件
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func loadIni(v interface{}) {

}

func main() {
	var mc MysqlConfig
	loadIni(&mc)
	fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)
}
