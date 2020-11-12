// 日志pkg, 系统日志基础库
// author: ericlang
package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// log等级
const (
	LOG_TRACE = 0
	LOG_DEBUG = 1
	LOG_INFO  = 2
	LOG_WARN  = 3
	LOG_ERROR = 4
)

// log文件创建方式
const (
	LOG_ALTER_HOUR  = 0
	LOG_ALTER_DAY   = 1
	LOG_ALTER_MONTH = 2
)

func init() {
	fmt.Println("yyyyyyyyy")
}

type LogEx struct {
	fileDir    string      // log输出目录
	filePrefix string      // log文件名前缀
	cTime      time.Time   // 文件创建时间, 省去每次读File的stat的逻辑
	logger     *log.Logger // 日志执行器
	level      int         // 当前输出日志等级
	alterMode  int         // 日志文件滚动方式
	alterTimer *time.Timer // 文件滚动计时器
}

func (logger *LogEx) createLogFile() error {
	// 拼接文件名
	var fileName string
	switch logger.alterMode {
	case LOG_ALTER_HOUR:
		fileName = fmt.Sprintf("%s/%s_%d_%02d_%02d_%02d.log", logger.fileDir, logger.filePrefix, logger.cTime.Year(), logger.cTime.Month(), logger.cTime.Day(), logger.cTime.Hour())
		time.AfterFunc(logger.cTime.Round(time.Hour).Sub(logger.cTime), alterLogFile)
	case LOG_ALTER_DAY:
		fileName = fmt.Sprintf("%s/%s_%d_%02d_%02d.log", logger.fileDir, logger.filePrefix, logger.cTime.Year(), logger.cTime.Month(), logger.cTime.Day())
		time.AfterFunc(logger.cTime.Round(24*time.Hour).Sub(logger.cTime), alterLogFile)
	case LOG_ALTER_MONTH:
		fileName = fmt.Sprintf("%s/%s_%d_%02d.log", logger.fileDir, logger.filePrefix, logger.cTime.Year(), logger.cTime.Month())
		time.AfterFunc(time.Date(logger.cTime.Year(), logger.cTime.Month(), 1, 0, 0, 0, 0, time.UTC).Sub(logger.cTime), alterLogFile)
	}

	outFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModeDir)
	if err != nil {
		fmt.Println("init log file failed, err:", err.Error())
		return err
	}
	// logger初始化一次
	once.Do(func() {
		logger.logger = log.New(outFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	})
	logger.logger.SetOutput(outFile)
	return nil
}

// 全局logger
var LEVEL_PREFIX_STR = []string{"[TRACE]", "[DEBUG]", "[INFO]", "[WARN]", "[ERROR]"}
var gLogEx LogEx
var once sync.Once

// 轮换log文件
func alterLogFile() {
	gLogEx.createLogFile()
}

// log初始化
func InitConf(dir string, filePrefix string, level int, alterMode int) error {
	gLogEx = LogEx{}
	gLogEx.fileDir = dir
	gLogEx.filePrefix = filePrefix
	gLogEx.level = level
	gLogEx.alterMode = alterMode
	gLogEx.cTime = time.Now()

	return gLogEx.createLogFile()
}

// 向log日志文件输出日志
func Println(level int, format string, v ...interface{}) {
	if level < gLogEx.level || gLogEx.logger == nil {
		return // 低等级log不输出
	}
	gLogEx.logger.Printf(LEVEL_PREFIX_STR[level]+format+"\r\n", v...)
}
