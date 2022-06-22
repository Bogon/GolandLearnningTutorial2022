package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"learning.goland.com/studyGoland/11day/logagent/conf"
	"learning.goland.com/studyGoland/11day/logagent/etcd"
	"learning.goland.com/studyGoland/11day/logagent/kafka"
	"learning.goland.com/studyGoland/11day/logagent/taillog"
	"learning.goland.com/studyGoland/11day/logagent/utils"
	"sync"
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

	// 打印区别 %+v 打印添加字段名
	fmt.Printf("cfg.Section(\"kafka\"): %+v \n", cfg.KafkaConf)

	// 1.  初始化一个kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
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
	// 为了实现每个 logagent 都能拉取自己独有的配置，所以要以自己的ip地址作为区分
	ip, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ip)
	// 2.1 从 etcd 中获取 日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("etcd.GetConf failed, error: %v \n", err)
		return
	}
	fmt.Printf("etcd.GetConf success conf: %#v \n", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index: %v, value:%#v, \n", index, value)
		fmt.Printf("index: #{index}, value: #{value} \n")
	}
	// 2.2 派发一个哨兵监听日志收集项的变化(有变化及时通知我的 logagent 实现热加载配置)

	// 3. 收集日志发送 Kafka
	taillog.Init(logEntryConf)
	var wg sync.WaitGroup
	newConfChan := taillog.NewConfChan() // 从 taillog 包中获取对外暴露的通道
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan) // 哨兵发现最新的配置信息会通知上面的通道
	wg.Wait()
	// 4. 具体的业务
	//run()

}
