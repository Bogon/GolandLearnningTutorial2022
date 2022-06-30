package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

/// 初始化 ES 准备接收 kafka 这边发送过来的数据

var (
	client *elastic.Client
	ch     chan *LogData
)

// LogData is a struct that contains a string field called data.
// @property {string} data - The data to be logged.
type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

// Init
// > Initialize the client with the given address
func Init(address string, chanSize, nums int) (err error) {

	ch = make(chan *LogData, chanSize)

	// 创建连接 elastic client
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		return err
	}
	fmt.Println("connect elastic success…")
	for i := 0; i < nums; i++ {
		go sendToES()
	}
	return
}

// SendToES 发送数据到 ES

// SendToESChan It sends the log data to the channel.
func SendToESChan(msg *LogData) {
	ch <- msg
}

// It's a goroutine that waits for messages to be sent to the channel, and when it receives a message, it sends it to Elasticsearch
func sendToES() {
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().Index(msg.Topic).BodyJson(msg).Do(context.Background())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Indexed user %s to index %s, type %s \n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}
}
