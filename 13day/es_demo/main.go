package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

/// ES Insert Data Demo

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

// 链式调用 举例
func (p *Person) run() *Person {
	fmt.Printf("%s在嗷嗷叫的跑步……\n", p.Name)
	return p
}

func (p *Person) eat() *Person {
	fmt.Printf("%s在哐哐哐的吃……\n", p.Name)
	return p
}

func main() {

	// 链式调用举例
	dsb := Person{
		Name:    "dsb",
		Age:     28,
		Married: true,
	}
	dsb.run().eat()

	// 创建连接 elastic client
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.39.95:9200"))
	if err != nil {
		panic(err)
	}
	fmt.Println("connect elastic success…")

	p1 := Person{Name: "琉璃瓦", Age: 26, Married: true}
	put1, err := client.Index().Index("student").BodyJson(p1).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Indexed user %s to index %s, type %s \n", put1.Id, put1.Index, put1.Type)

}
