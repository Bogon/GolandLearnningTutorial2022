package taillog

import (
	"fmt"
	"learning.goland.com/studyGoland/logagent/etcd"
	"time"
)

var tskMgr *tailLogMgr

// task 管理者
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry // 把当前日志收集项配置信息保存起来
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf, // 把当前的日志收集项保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), // 无缓冲区的通道
	}
	for _, logEntry := range logEntryConf {
		//初始化的时候起了多少个 tailObj，记录一下，为后续做判断方便
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[mk] = tailObj
	}

	go tskMgr.run()
}

// 监听自己的 newConfChan 有了新配置就做对应处理
// 1. 配置新增
// 2. 配置删除
// 3. 配置变更
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := tskMgr.tskMap[mk]
				if ok {
					// 原来就有，不需要操作
					continue
				} else {
					// 原来不存在，就初始化配置项，新增
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}
			// 找出 t.logEntry 中有，newConf中没有，找出来删除
			for _, c1 := range t.logEntry {
				isDelete := true
				for _, c2 := range newConf {
					if c1.Path == c2.Path && c1.Topic == c2.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					// 把c1对应的这个tailObj给停掉
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					t.tskMap[mk].cancelFunc()
				}
			}
			// 2. 配置删除
			fmt.Println("新的配置来了！", newConf)

		default:
			time.Sleep(time.Second)
		}
	}
}

// NewConfChan 定义一个函数，向外暴露 tskMgr 的 newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
