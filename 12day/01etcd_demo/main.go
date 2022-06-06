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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "bandelu", "dsb")
	cancel()
	if err != nil {
		// handle error!
		fmt.Printf("get from etcd failed, err: %v \n", err)
		return
	}
	// use the response
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "bandelu", clientv3.WithPrefix())
	cancel()
	if err != nil {
		// handle error!
		fmt.Printf("get from etcd failed, err: %v \n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s \n", ev.Key, ev.Value)
	}
}
