package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

/**
日志库 需求分析
1、支持向不同的地方输出日志
2、日志分级别
	1、Debug
	2、Trace
	2、Info
	3、Warning
	4、Error
	5、Fatal
3、日志要支持开关控制
4、日志要有时间、行号、文件名、日志级别、日志信息
5、日志文件要切割
	1、按照文件大小切割
		1、每次记录日志之前都判断当前写的这个文件的大小
	2、按照日期切割
		1、记录上次切割的小时，按照小时切割文件
		2、在写日志之前检查一下当前时间的小时数和切割的小时数是不是一致
*/

// Logger 构造方法
func NewConsoleLog(l string) ConsoleLogger {
	level, err := ParseLogLevel(l)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{Level: level}
}

func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(DEBUG) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := c.getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15-04-05"), GetLogLevelStr(lv), funcName, fileName, lineNo, msg)
	}
}

// Debug
func (c ConsoleLogger) Debug(fromat string, a ...interface{}) {
	c.log(DEBUG, fromat, a...)
}

// Trace
func (c ConsoleLogger) Trace(fromat string, a ...interface{}) {
	c.log(TRACE, fromat, a...)
}

// Info
func (c ConsoleLogger) Info(fromat string, a ...interface{}) {
	c.log(INFO, fromat, a...)
}

// Warning
func (c ConsoleLogger) Warning(fromat string, a ...interface{}) {
	c.log(WARNING, fromat, a...)
}

// Error
func (c ConsoleLogger) Error(fromat string, a ...interface{}) {
	c.log(ERROR, fromat, a...)
}

// Fatal
func (c ConsoleLogger) Fatal(fromat string, a ...interface{}) {
	c.log(FATAL, fromat, a...)
}

func (c ConsoleLogger) getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
