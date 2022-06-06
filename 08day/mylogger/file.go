package mylogger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

// 向文件里面写日志相关代码

type FileLogger struct {
	Level        LogLevel
	filePath     string // 日志文件保存路径
	fileName     string // 日志文件保存的文件名
	maxFileSize  int64
	fileObj      *os.File
	errorFileObj *os.File
	logChan      chan *logMsg
}

type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

// 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := ParseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{Level: logLevel, filePath: fp, fileName: fn, maxFileSize: maxSize, logChan: make(chan *logMsg, 50000)}
	fl.fileInit() // 按照文件路径和文件名将文件打开
	return fl
}

func (f *FileLogger) fileInit() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return err
	}

	errorFileObj, err := os.OpenFile(fullFileName+".err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open error file failed, err:", err)
		return err
	}

	// 日志文件都已经打开
	f.fileObj = fileObj
	f.errorFileObj = errorFileObj
	// 开启5个后台的 goroutine 去写日志
	for i := 0; i < 5; i++ {
		go f.writeLogBackground()
	}
	return nil
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

func (f *FileLogger) checkSize(file *os.File) (bool, error) {
	// 获取文件对象的详细信息
	fileInfo, err := file.Stat()
	if err != nil {
		return false, err
	}
	// 如果传入的文件大小大于日志文件的最大值，就返回：true
	return fileInfo.Size() > f.maxFileSize, nil
}

func (f *FileLogger) splitFile(fp, fn string, lv LogLevel, fl *os.File) (*os.File, error) {
	// 需要切割文件
	// 1、关闭当前的日志文件
	fl.Close()
	// 2、备份一下 rename: xx.log => xx.log.bak29831982389
	newStr := time.Now().Format("20160102150405000")
	var logName string = path.Join(fp, fn) // 当前的文件的名称和路径
	if lv >= ERROR {
		logName += ".err.log"
	}
	newLogName := fmt.Sprintf("%s.bak%s.log", logName, newStr)
	os.Rename(logName, newLogName)
	// 3、打开一个新文件，复制给新的 fileObj
	fileObj, err := os.OpenFile(newLogName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return nil, err
	}
	// 将新打开的文件日志对象赋值给 f.fileObj
	return fileObj, nil
}

func (f *FileLogger) writeLogBackground() {
	for {
		select {
		case logTmp := <-f.logChan:
			lv := logTmp.level
			fileCheckSize, err := f.checkSize(f.fileObj)
			if err != nil {
				return
			}
			if fileCheckSize {
				newFile, err := f.splitFile(f.filePath, f.fileName, lv, f.fileObj)
				if err != nil {
					return
				}
				f.fileObj = newFile
			}

			// log info
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTmp.timestamp, GetLogLevelStr(lv), logTmp.funcName, logTmp.fileName, logTmp.line, logTmp.msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if lv >= ERROR {
				errFileCheckSize, err := f.checkSize(f.errorFileObj)
				if err != nil {
					return
				}
				if errFileCheckSize {
					newFile, err := f.splitFile(f.filePath, f.fileName, lv, f.errorFileObj)
					if err != nil {
						return
					}
					f.errorFileObj = newFile
				}
				// 如果要记录的日志大于等于ERROR的级别，还要记录到 err 文件中
				fmt.Fprintf(f.errorFileObj, logInfo)
			}
		default:
			time.Sleep(time.Millisecond * 500)
		}

	}
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		lmsg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := f.getInfo(3)
		// 把日志写入到通道中
		// 1. 造一个logMsg对象
		logTmp := &logMsg{
			level:     lv,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02 15-04-05"),
			line:      lineNo,
			msg:       lmsg,
		}
		select {

		case f.logChan <- logTmp:
		default:
			// 把日志丢掉保证不出现阻塞
		}
	}
}

// Debug
func (f *FileLogger) Debug(fromat string, a ...interface{}) {
	f.log(DEBUG, fromat, a...)
}

// Trace
func (f *FileLogger) Trace(fromat string, a ...interface{}) {
	f.log(TRACE, fromat, a...)
}

// Info
func (f *FileLogger) Info(fromat string, a ...interface{}) {
	f.log(INFO, fromat, a...)
}

// Warning
func (f *FileLogger) Warning(fromat string, a ...interface{}) {
	f.log(WARNING, fromat, a...)
}

// Error
func (f *FileLogger) Error(fromat string, a ...interface{}) {
	f.log(ERROR, fromat, a...)
}

// Fatal
func (f *FileLogger) Fatal(fromat string, a ...interface{}) {
	f.log(FATAL, fromat, a...)
}

func (f *FileLogger) getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errorFileObj.Close()
}
