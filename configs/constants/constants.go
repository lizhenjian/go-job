package constants

import (
	"unsafe"
)

var Sleep = 2                                      // 队列没消息时，暂停秒数
var ProcessNameStatic = "goJobTopicQueueStatic:"   // 设置静态进程名, 方便管理, 默认值 goJobTopicQueueStatic
var ProcessNameDynamic = "goJobTopicQueueDynamic:" // 设置动态进程名, 方便管理, 默认值 goJobTopicQueueDynamic

var LogConf = struct {
	Dir                string `yaml:"dir"`
	ApplicationLogName string `yaml:"application_log_name"`
	ProcessLogName     string `yaml:"application_log_name"`
	Level              int    `yaml:"level"`
	MaxSize            int    `yaml:"max_size"`
	MaxBackups         int    `yaml:"max_backups"`
	MaxAge             int    `yaml:"max_age"`
	LocalTime          bool   `yaml:"local_time"`
}{
	Dir:                "./logs",
	ApplicationLogName: "application.log", //默认log存储名字
	ProcessLogName:     "process.log",     //默认log存储名字
	MaxSize:            100,               // 日志文件最大 size(MB)，缺省 100MB。
	MaxBackups:         5,                 // 最大过期日志保留的个数。
	MaxAge:             30,                // 保留过期文件的最大时间间隔，单位是天。
	LocalTime:          true,              // 是否使用本地时间来命名备份的日志。
}

var MessageFunc = ""      //消息通知函数
var Token = ""            //消息oken
var MaxPopNum = 100       //最大弹出数量
var TimerProcessCheck = 2 //进程检查时间间隔
var TimerListCheck = 5    //队列长度检查

type Topic struct {
	PkgName                string //包名
	Name                   string //方法名
	WorkerMinNum           int64  //最小工作进程数
	WorkerMaxNum           int64  //最大工作进程数
	WorkerMaxPendingLength int64  //最大队列长度
}

var ImportedFunctions unsafe.Pointer = nil

var Topics = []Topic{
	{PkgName: "exportJob", Name: "HandleExportJob", WorkerMinNum: 2, WorkerMaxNum: 5, WorkerMaxPendingLength: 50},
	{PkgName: "importJob", Name: "HandleImportJob", WorkerMinNum: 2, WorkerMaxNum: 5, WorkerMaxPendingLength: 50},
}
