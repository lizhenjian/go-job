package constants

import "unsafe"

var LogPath = ""
var LogSaveFileApp = "application.log" //默认log存储名字
var LogSaveFileWorker = "crontab.log"  // 进程启动相关log存储名字
var PidPath = ""                       // 默认pid存储路径
var Sleep = 2                          // 队列没消息时，暂停秒数
var ProcessName = "goJobTopicQueue"    // 设置进程名, 方便管理, 默认值 goJobTopicQueue

var MessageFunc = ""      //消息通知函数
var Token = ""            //消息oken
var MaxPopNum = 100       //最大弹出数量
var TimerProcessCheck = 2 //进程检查时间间隔
var TimerListCheck = 5    //队列长度检查

type Topic struct {
	PkgName                string
	Name                   string
	WorkerMinNum           int64
	WorkerMaxNum           int64
	WorkerMaxPendingLength int64
}

var ImportedFunctions unsafe.Pointer = nil

var Topics = []Topic{
	{PkgName: "exportJob", Name: "HandleExportJob", WorkerMinNum: 1, WorkerMaxNum: 5, WorkerMaxPendingLength: 50},
	{PkgName: "importJob", Name: "HandleImportJob", WorkerMinNum: 1, WorkerMaxNum: 5, WorkerMaxPendingLength: 50},
}
