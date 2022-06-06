package taillog

import "learning.goland.com/studyGoland/11day/logagent/etcd"

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry []*etcd.LogEntry // 把当前日志收集项配置信息保存起来
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{logEntry: logEntryConf}
	for _, logEntry := range logEntryConf {
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}
