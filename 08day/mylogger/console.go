package mylogger

import (
	"errors"
	"strings"
)

type LogLevel uint16

type Logger interface {
	Debug(fromat string, a ...interface{})
	Trace(fromat string, a ...interface{})
	Info(fromat string, a ...interface{})
	Warning(fromat string, a ...interface{})
	Error(fromat string, a ...interface{})
	Fatal(fromat string, a ...interface{})
}

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// Logger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

func ParseLogLevel(s string) (LogLevel, error) {
	s = strings.ToUpper(s)
	switch s {
	case "DEBUG":
		return DEBUG, nil

	case "TRACE":
		return TRACE, nil

	case "INFO":
		return INFO, nil

	case "WARNING":
		return WARNING, nil

	case "ERROR":
		return ERROR, nil

	case "FATAL":
		return FATAL, nil

	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func GetLogLevelStr(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}
