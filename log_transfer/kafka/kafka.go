package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"learning.goland.com/studyGoland/log_transfer/es"
)

/// 初始化连接 kafka 的一个 client

// Init We create a consumer, get a list of partitions for the topic, and then start a goroutine for each partition that
// consumes messages from that partition
func Init(addr []string, topic string) error {
	consumer, err := sarama.NewConsumer(addr, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return err
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition, error: %v\n", err)
		return err
	}
	fmt.Printf("分区列表：%v \n", partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d, error: %v\n", partition, err)
			return err
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				// 直接发送给 ES
				msgValue := string(msg.Value)
				if msgValue != "" { // 不为空的时候发送数据
					fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v \n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
					ld := es.LogData{Topic: topic, Data: msgValue}
					//_ = es.SendToES(topic, ld) // 函数调用函数 -> 改成异步
					// 优化一下：将数据放到 chan 中
					es.SendToESChan(&ld)
				}

			}
		}(pc)
	}
	select {}
	return err
}
