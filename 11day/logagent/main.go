package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"learning.goland.com/studyGoland/11day/logagent/conf"
	"learning.goland.com/studyGoland/11day/logagent/etcd"
	"learning.goland.com/studyGoland/11day/logagent/kafka"
	"learning.goland.com/studyGoland/11day/logagent/taillog"
	"time"
)

var (
	cfg = new(conf.AppConf)
)

//func run() {
//	// 1. 读取日志
//	for {
//		select {
//		case line := <-taillog.ReadChan():
//			kafka.SendToKafka(cfg.Topic, line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//	}
//	// 2. 发送到 kafka
//}

// logAgent 入口
func main() {
	// 0. 加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("ini mapto config failed, error: ", err)
		return
	}
	//cfg, err := ini.Load("./conf/config.ini")
	//fmt.Println(cfg.Section("kafka").Key("address"))
	//fmt.Println(cfg.Section("kafka").Key("topic"))
	//fmt.Println(cfg.Section("taillog").Key("path"))

	// 1.  初始化一个kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Println("init kafa failed, error: ", err)
		return
	}
	fmt.Println("kafka 初始化成功……")
	// 2.初始化 etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("etcd init failed, error: %v \n", err)
		return
	}
	fmt.Println("etcd 初始化成功……")
	// 2.1 从 etcd 中获取 日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Printf("etcd.GetConf failed, error: %v \n", err)
		return
	}
	fmt.Printf("etcd.GetConf success conf: %#v \n", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index: %v, value:%#v, \n", index, value)
	}
	// 2.2 派发一个哨兵监听日志收集项的变化(有变化及时通知我的 logagent 实现热加载配置)

	// 3. 收集日志发送 Kafka
	taillog.Init(logEntryConf)

	// 4. 具体的业务
	//run()

}
