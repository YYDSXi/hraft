package log

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

type logFileWriter struct {
	infoWriter  *os.File
	debugWriter *os.File
	errorWriter *os.File
	warnWriter  *os.File
	rootPath    string
	fileDate    string //日期
	entryLevel  string
}

const (
	infoLevel  = "info"
	debugLevel = "debug"
	errorLevel = "error"
	warnLevel  = "warn"
)

var (
	fileWriter = &logFileWriter{rootPath: "./scopeLog", entryLevel: infoLevel}
	stdout     = colorable.NewColorableStdout()
)

type LoggerFormat struct {
	logWriter *logFileWriter
}

func (s *LoggerFormat) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05.000")
	var file string
	var length int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		length = entry.Caller.Line
	}
	//d为1 表示高亮
	//red: (prefix=31,b=40,d=1)
	//green (f=32,b=40,d=1)
	//yellow (f=33,b=40,d=1)
	//cyan (f=34,b=40,d=1)
	var colorFomat string
	switch entry.Level {
	case log.DebugLevel:
		//cyan
		colorFomat = fmt.Sprintf("%c[%d;%d;%dm", 0x1B, 1, 1, 34)
		s.logWriter.entryLevel = debugLevel
	case log.InfoLevel:
		//green
		colorFomat = fmt.Sprintf("%c[%d;%d;%dm", 0x1B, 1, 1, 32)
		s.logWriter.entryLevel = infoLevel
	case log.ErrorLevel:
		//red [显示方式（1：高亮）、背景颜色(黑色:40)、颜色 ]
		colorFomat = fmt.Sprintf("%c[%d;%d;%dm", 0x1B, 1, 1, 31)
		s.logWriter.entryLevel = errorLevel
	case log.WarnLevel:
		//yellow
		colorFomat = fmt.Sprintf("%c[%d;%d;%dm", 0x1B, 1, 1, 33)
		s.logWriter.entryLevel = warnLevel
	}
	msg := fmt.Sprintf("%s[%s:] [%s:%d] %s \u001B[0m msg:%s \n", colorFomat, entry.Level.String(), file, length, timestamp, entry.Message)
	return []byte(msg), nil
}

func (p *logFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	//判断是否需要切换日期
	fileDate := time.Now().Format("20060102")
	if p.fileDate != fileDate {
		switch p.entryLevel {
		case infoLevel:
			err = os.MkdirAll(fmt.Sprintf("%s/%s", p.rootPath, infoLevel), os.ModePerm)
			if err != nil {
				fmt.Errorf("%s\n", err.Error())
				return 0, err
			}
			infofileName := fmt.Sprintf("%s/%s/%s-%s.logger", p.rootPath, infoLevel, infoLevel, fileDate)
			p.infoWriter, err = os.OpenFile(infofileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
			checkError(err)
			writer := colorable.NewColorable(fileWriter.infoWriter)
			n, err = writer.Write(data)
		case errorLevel:
			err = os.MkdirAll(fmt.Sprintf("%s/%s", p.rootPath, errorLevel), os.ModePerm)
			if err != nil {
				fmt.Errorf("%s\n", err.Error())
				return 0, err
			}
			errorfileName := fmt.Sprintf("%s/%s/%s-%s.logger", p.rootPath, errorLevel, errorLevel, fileDate)
			p.errorWriter, err = os.OpenFile(errorfileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
			checkError(err)
			writer := colorable.NewColorable(fileWriter.errorWriter)
			n, err = writer.Write(data)
		case debugLevel:
			err = os.MkdirAll(fmt.Sprintf("%s/%s", p.rootPath, debugLevel), os.ModePerm)
			if err != nil {
				fmt.Errorf("%s\n", err.Error())
				return 0, err
			}
			debugfileName := fmt.Sprintf("%s/%s/%s-%s.logger", p.rootPath, debugLevel, debugLevel, fileDate)
			p.debugWriter, err = os.OpenFile(debugfileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
			checkError(err)
			writer := colorable.NewColorable(fileWriter.debugWriter)
			n, err = writer.Write(data)
		case warnLevel:
			err = os.MkdirAll(fmt.Sprintf("%s/%s", p.rootPath, warnLevel), os.ModePerm)
			if err != nil {
				fmt.Errorf("%s\n", err.Error())
				return 0, err
			}
			warnfileName := fmt.Sprintf("%s/%s/%s-%s.logger", p.rootPath, warnLevel, warnLevel, fileDate)
			p.warnWriter, err = os.OpenFile(warnfileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
			checkError(err)
			writer := colorable.NewColorable(fileWriter.warnWriter)
			n, err = writer.Write(data)
		default:
			fmt.Errorf("unexpected logger level")
		}
	}

	return
}
func checkError(err error) {
	if err != nil {
		fmt.Errorf("%s\n", err)
		return
	}
}

//初始化日志
func Init() {
	rootPath := "./scopeLog"
	log.SetLevel(log.DebugLevel)
	//创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", rootPath, infoLevel), os.ModePerm)
	checkError(err)
	err = os.MkdirAll(fmt.Sprintf("%s/%s", rootPath, errorLevel), os.ModePerm)
	checkError(err)
	err = os.MkdirAll(fmt.Sprintf("%s/%s", rootPath, debugLevel), os.ModePerm)
	checkError(err)
	err = os.MkdirAll(fmt.Sprintf("%s/%s", rootPath, warnLevel), os.ModePerm)
	checkError(err)
	log.SetFormatter(&LoggerFormat{
		logWriter: fileWriter,
	})
	log.SetReportCaller(true)
	log.SetOutput(fileWriter)
}
