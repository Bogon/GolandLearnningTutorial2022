package taillog

import (
	"context"
	"fmt"
	"github.com/nxadm/tail"
	"learning.goland.com/studyGoland/logagent/kafka"
)

// TailTask 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	// 为了能够实现退出 t.run()
	ctx        context.Context
	cancelFunc context.CancelFunc
}

// NewTailTask 创建 TailTask 实例
func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancle := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancle,
	}
	tailObj.init() // 根据路径打开对应的日志
	return
}

func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的那个位置打开
		MustExist: false,                                // 文件不存在报错
		Poll:      true,                                 //
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, error: ", err)
	}
	// 当 goroutine 执行的函数退出的时候，goroutine 就退出了
	go t.run() //直接采集日志到kafka
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task: %s_%s 结束了… \n", t.path, t.topic)
			return
		case line := <-t.instance.Lines: // 从 tailobj 的通道中一行一行的读取日志数据
			// 3.2 发往 Kafka
			//kafka.SendToKafka(t.topic, line.Text) // 函数调用函数
			// 先把 日志数据发送到有个通道中
			kafka.SendToChan(t.topic, line.Text)
			// kafka 那个包中 有单独的 goroutine 去取日志数据数据发送到 kafka
		}
	}
}
