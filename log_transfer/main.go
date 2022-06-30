package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"learning.goland.com/studyGoland/log_transfer/conf"
	"learning.goland.com/studyGoland/log_transfer/es"
	"learning.goland.com/studyGoland/log_transfer/kafka"
)

var (
	cfg = new(conf.LogTransferCfg)
)

/// log transfer
/// 从日志数据冲kafka中读取出来发送到ES

func main() {
	// 0.0 加载配置文件
	err := ini.MapTo(cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Println("cfg init failed, err: ", err)
		return
	}
	fmt.Printf("cfg: %v \n", cfg)
	// 1. 初始化 ES
	// 1.1 初始化一个链接 ES 的 client
	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanSize, cfg.ESCfg.Nums)
	if err != nil {
		fmt.Println("init ES client failed, error: ", err)
		return
	}
	fmt.Println("init es success……")
	// 2. 初始化 kafka
	// 2.1 链接 kafka ，创建分区接收者
	// 2.2 每个分区消费者分别取出数据，通过 SendToES 发送到 ES
	// 2.3 对外提供一个想 ES 中 写入数据的函数
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Println("init kafka consumer failed, error: ", err)
		return
	}

	select {}
}
