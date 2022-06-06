package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()
	fmt.Println("connect etcd success…")

	// watch
	// 排一个哨兵一直监视着 liminghui, 这个 key 的变化，新增修改删除
	ch := cli.Watch(context.Background(), "liminghui")

	// 从通道尝试取值
	for wresp := range ch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %v, Key: %v, Value: %v \n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
		}
	}
}
