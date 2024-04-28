package loger

import (
	"fmt"
	"go-jobs/configs/constants"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// Level 日志级别。建议从服务配置读取。
var LogConf = constants.LogConf

// Init logrus logger.
func InitLoggerApplication() *logrus.Logger {
	// 设置日志格式。
	var LogApp = logrus.New()
	LogApp.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	LogApp.SetLevel(logrus.DebugLevel)
	LogApp.SetReportCaller(true) // 打印文件、行号和主调函数。
	// 实现日志滚动。
	logger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%v/%v", LogConf.Dir, LogConf.ApplicationLogName), // 日志输出文件路径。
		MaxSize:    LogConf.MaxSize,                                               // 日志文件最大 size(MB)，缺省 100MB。
		MaxBackups: LogConf.MaxBackups,                                            // 最大过期日志保留的个数。
		MaxAge:     LogConf.MaxAge,                                                // 保留过期文件的最大时间间隔，单位是天。
		LocalTime:  LogConf.LocalTime,                                             // 是否使用本地时间来命名备份的日志。
	}
	// 同时输出到标准输出与文件。
	LogApp.SetOutput(io.MultiWriter(logger, os.Stdout))
	return LogApp
}

func InitLoggerProcess() *logrus.Logger {
	// 设置日志格式。
	var LogProcess = logrus.New()
	LogProcess.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	LogProcess.SetLevel(logrus.DebugLevel)
	LogProcess.SetReportCaller(true) // 打印文件、行号和主调函数。
	// 实现日志滚动。
	logger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%v/%v", LogConf.Dir, LogConf.ProcessLogName), // 日志输出文件路径。
		MaxSize:    LogConf.MaxSize,                                           // 日志文件最大 size(MB)，缺省 100MB。
		MaxBackups: LogConf.MaxBackups,                                        // 最大过期日志保留的个数。
		MaxAge:     LogConf.MaxAge,                                            // 保留过期文件的最大时间间隔，单位是天。
		LocalTime:  LogConf.LocalTime,                                         // 是否使用本地时间来命名备份的日志。
	}
	// 同时输出到标准输出与文件。
	LogProcess.SetOutput(io.MultiWriter(logger, os.Stdout))
	return LogProcess
}

func InitLoggerError() *logrus.Logger {
	// 设置日志格式。
	var LoggerError = logrus.New()
	LoggerError.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	LoggerError.SetLevel(logrus.DebugLevel)
	LoggerError.SetReportCaller(true) // 打印文件、行号和主调函数。
	// 实现日志滚动。
	logger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%v/%v", LogConf.Dir, LogConf.ErrorLogName), // 日志输出文件路径。
		MaxSize:    LogConf.MaxSize,                                         // 日志文件最大 size(MB)，缺省 100MB。
		MaxBackups: LogConf.MaxBackups,                                      // 最大过期日志保留的个数。
		MaxAge:     LogConf.MaxAge,                                          // 保留过期文件的最大时间间隔，单位是天。
		LocalTime:  LogConf.LocalTime,                                       // 是否使用本地时间来命名备份的日志。
	}
	// 同时输出到标准输出与文件。
	LoggerError.SetOutput(io.MultiWriter(logger, os.Stdout))
	return LoggerError
}
