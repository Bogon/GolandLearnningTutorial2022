package main

import "fmt"

// map

func main() {
	// 显示声明
	var m1 map[string]int         // panic: assignment to entry in nil map
	fmt.Println(m1 == nil)        // 还没有在内存中开辟空间
	m1 = make(map[string]int, 10) // 要估算好空间，避免在程序运行过程中动态开辟空间
	m1["int"] = 0
	m1["string"] = 1
	m1["map"] = 2
	m1["slice"] = 3
	m1["chan"] = 4
	m1["float"] = 5
	m1["double"] = 6
	m1["bool"] = 7
	fmt.Println(m1)

	fmt.Println(m1["bool"])

	fmt.Println(m1["double_int"]) // 如果不存在这个key拿到对应值类型的零值

	// 预定俗成用ok接收返回的bool值
	value, ok := m1["double_int"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("查无此key")
	}

	// map 遍历
	for k, v := range m1 {
		fmt.Printf("k: %v, v: %v \n", k, v)
	}

	// 只遍历 key
	for k := range m1 {
		fmt.Printf("key: %v \n", k)
	}

	// 只遍历 value
	for _, value := range m1 {
		fmt.Printf("value: %v \n", value)
	}

	// 删除
	delete(m1, "int")
	fmt.Println(m1)
	// 删除 不存在的 key
	delete(m1, "double_int")
	fmt.Println(m1)

}
