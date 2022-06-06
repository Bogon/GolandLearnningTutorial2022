package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// redis
var redisdb *redis.Client

// 初始化连接
func initClient() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err = redisdb.Ping().Result()
	return
}

func main() {

	err := initClient()
	if err != nil {
		fmt.Println("init redis db failed, error: ", err)
		return
	}
	fmt.Println("连接Redis成功！")

	// zset
	key := "rank"
	items := []redis.Z{
		redis.Z{Score: 90, Member: "PHP"},
		redis.Z{Score: 91, Member: "SWIFT"},
		redis.Z{Score: 92, Member: "OC"},
		redis.Z{Score: 93, Member: "C++"},
		redis.Z{Score: 94, Member: "C"},
	}

	// 把元素追加到key
	redisdb.ZAdd(key, items...)

	// 添加新的key
	newScore, err := redisdb.ZIncrBy(key, 10.0, "Golang").Result()
	if err != nil {
		fmt.Println("zincrby failed, error: ", err)
		return
	}
	fmt.Printf("Golangis score is %f . \n", newScore)

}
