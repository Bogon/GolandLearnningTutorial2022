package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
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

func Init(addrs []string, maxSize int) (err error) {
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
	// 初始化 logdatachan
	logDataChan = make(chan *logData, maxSize)
	// 开启后台的 goroutine 从通道中读取数据发送到 kafka
	go SendToKafka()
	return
}

// SendToChan 给外部暴露一个函数，该函数把数据发送到一个 channel 中
func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

// SendToKafka 向 kafka 中发送消息
// 真正向 kafka 中发送数据
func SendToKafka() {

	for {
		select {
		case ld := <-logDataChan:
			// 构造一个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			// 发送到 kafka
			// 发送消息
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, error:", err)
				return
			}
			fmt.Printf("pid:%v, offset:%v \n", pid, offset)
			fmt.Println("发送成功")
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}

}
