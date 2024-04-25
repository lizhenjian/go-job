package loger

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// Level 日志级别。建议从服务配置读取。
var LogConf = struct {
	Dir     string `yaml:"dir"`
	Name    string `yaml:"name"`
	Level   string `yaml:"level"`
	MaxSize int    `yaml:"max_size"`
}{
	Dir:     "./logs",
	Name:    "application.log",
	Level:   "trace",
	MaxSize: 1,
}

// Init logrus logger.
func InitLogger() error {
	// 设置日志格式。
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true) // 打印文件、行号和主调函数。
	// 实现日志滚动。
	logger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%v/%v", LogConf.Dir, LogConf.Name), // 日志输出文件路径。
		MaxSize:    LogConf.MaxSize,                                 // 日志文件最大 size(MB)，缺省 100MB。
		MaxBackups: 10,                                              // 最大过期日志保留的个数。
		MaxAge:     30,                                              // 保留过期文件的最大时间间隔，单位是天。
		LocalTime:  true,                                            // 是否使用本地时间来命名备份的日志。
	}
	// 同时输出到标准输出与文件。
	logrus.SetOutput(io.MultiWriter(logger, os.Stdout))
	return nil
}
