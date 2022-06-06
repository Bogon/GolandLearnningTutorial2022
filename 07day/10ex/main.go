package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 使用 goroutine 和 channel 实现一个计算 int64 随机数个位数和的程序
// 1、开启一个 goroutine 循环生成 int64 类型的随机数，发送到 jobchan
// 2、开启24个goroutine 从 jobchan 中取出随机数计算各位数的和，将结果发送到 resultChan
// 3、主 goroutine 从resultChan中取出结果并打印到终端输出

type Job struct {
	value int64
}

type Result struct {
	job    *Job
	result int64
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)
var wg sync.WaitGroup

func zhoulin(zl chan<- *Job) {
	// 1、开启一个 goroutine 循环生成 int64 类型的随机数，发送到 jobchan
	for {
		x := rand.Int63()
		newJob := &Job{
			value: x,
		}
		zl <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func baodelu(jo <-chan *Job, resultChan chan<- *Result) {
	defer wg.Done()
	// 2、开启24个goroutine 从 jobchan 中取出随机数计算各位数的和，将结果发送到 resultChan
	for {
		job := <-jo
		sum := int64(0)
		n := job.value
		for n > 10 {
			sum += n % 10
			n = n / 10
		}
		newResult := &Result{
			job:    job,
			result: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	wg.Add(1)
	go zhoulin(jobChan)
	// 开启24 goroutine 执行保存
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go baodelu(jobChan, resultChan)
	}
	// 3、主 goroutine 从resultChan中取出结果并打印到终端输出
	for result := range resultChan {
		fmt.Printf("value: %d, sum: %d\n", result.job.value, result.result)
	}
	wg.Wait()

}
