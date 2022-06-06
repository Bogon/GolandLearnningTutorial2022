package main

import (
	"fmt"
	"strconv"
	"sync"
)

// go语言中内置的map 是并发不安全的

// 同时也提供了一个开箱即用的 并发安全的 map
var m = sync.Map{}
var wg sync.WaitGroup

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=%v, v=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
