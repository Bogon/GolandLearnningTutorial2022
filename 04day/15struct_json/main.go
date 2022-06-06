package main

import (
	"encoding/json"
	"fmt"
)

// 结构体/Json
type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age" db:"age" ini:"age"`
}

func main() {
	p1 := person{
		Name: "李磊",
		Age:  20,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", string(b))

	// 反序列化
	str := `{"name": "李磊", "age": 20}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)

}
