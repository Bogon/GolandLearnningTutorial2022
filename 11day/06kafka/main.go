package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// kafka demo

// 基于sarama 的第三方库开发的 kafka 客户端
func main() {

	config := sarama.NewConfig()
	// tailf 使用
	config.Producer.RequiredAcks = sarama.WaitForAll // 发送完整数据需要
	// leader 和 follower 确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个 partition
	config.Producer.Return.Successes = true                   // 成功交付的信息将在success channel 返回
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, error:", err)
		return
	}
	fmt.Println("连接 kafka 成功")
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, error:", err)
		return
	}
	fmt.Printf("pid:%v, offset:%v \n", pid, offset)

}
