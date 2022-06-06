package main

import (
	"fmt"
	"sync"
	"time"
)

// rwlock

var (
	x      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func read() {
	defer wg.Done()
	lock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	lock.Unlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	x += 2
	time.Sleep(time.Millisecond * 5)
	lock.Unlock()
}

func main() {

	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}

	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
