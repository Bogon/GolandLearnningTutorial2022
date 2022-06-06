package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type logData struct {
	topic string
	data  string
}

// 专门向kafka中写日志
var (
	client      sarama.SyncProducer // 声明一个全局的链接 kafka
	logDataChan chan *logData
)

func Init(addrs []string) (err error) {
	config := sarama.NewConfig()
	// tailf 使用
	config.Producer.RequiredAcks = sarama.WaitForAll // 发送完整数据需要
	// leader 和 follower 确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个 partition
	config.Producer.Return.Successes = true                   // 成功交付的信息将在success channel 返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config) /// []string{"127.0.0.1:9092"}
	if err != nil {
		fmt.Println("producer closed, error:", err)
		return
	}
	return
}

func SendToChan() {

}

// 向 kafka 中发送消息
func SendToKafka(topic, data string) {
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	// 发送到 kafka
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, error:", err)
		return
	}
	fmt.Printf("pid:%v, offset:%v \n", pid, offset)
	fmt.Println("发送成功")

}
