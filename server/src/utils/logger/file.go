package logger

import (
	"fmt"
	"time"
)

var (
	// LogSavePath 日志存放路径
	LogSavePath = "logs"
	// LogSaveName 日志存放名称
	LogSaveName = ""
	// LogFileExt a
	LogFileExt = "log"
	// TimeFormat 日期格式化
	TimeFormat = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s/%s", prefixPath, suffixPath)
}
