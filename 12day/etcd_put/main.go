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
	key := "/logagent/collect_config"
	value := `[{"path":"/Users/zhangqi/go/src/learning.goland.com/studyGoland/11day/logagent/logs/ngnix.log","topic":"web_log"},{"path":"/Users/zhangqi/go/src/learning.goland.com/studyGoland/11day/logagent/logs/redis.log","topic":"redis_log"},{"path":"/Users/zhangqi/go/src/learning.goland.com/studyGoland/11day/logagent/logs/mysql.log","topic":"mysql_log"}]`
	_, err = cli.Put(ctx, key, value)
	cancel()
	if err != nil {
		// handle error!
		fmt.Printf("get from etcd failed, err: %v \n", err)
		return
	}
	// use the response
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key, clientv3.WithPrefix())
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
