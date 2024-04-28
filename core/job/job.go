package job

import (
	"go-jobs/configs/constants"
	"go-jobs/configs/register"
	"go-jobs/helpers/loger"
	"go-jobs/helpers/util"
	"go-jobs/service/service_redis"
	"time"
)

type MyStruct struct {
}

var LogApp = loger.InitLoggerApplication()
var LogError = loger.InitLoggerError()

/**
 * @Author: lizhenjian
 * @LastEditors: lizhenjian
 * @Description: 任务运行方法，通过反射调用业务方法
 * @param {string} PkgName
 * @param {string} topicName
 */
func Run(PkgName string, topicName string) {
	LogApp.Infoln(topicName + " is running")
	register.RegisterFunc()
	i := 0
	for {
		//单进程最大执行任务次数
		i++
		if i >= constants.MaxPopNum {
			LogApp.Infoln("MaxPopNum:", constants.MaxPopNum)
			break
		}
		//拉取Redis队列消息
		rdb := service_redis.NewClient()
		jobParams, err := rdb.RPop(topicName).Result()
		if err != nil {
			LogApp.Infoln("Pop data error:", err, topicName)
			//队列没有消息停留一秒
			time.Sleep(time.Duration(constants.Sleep) * time.Second)
		}
		if jobParams != "" {
			LogApp.Infoln("Popped data:", jobParams, topicName)
			//通过包名和方法名动态调用业务方法
			funcName := topicName
			args := []interface{}{jobParams}
			results, err := util.Eval(PkgName, funcName, args...)
			if err != nil {
				LogError.Errorln("Error:", err)
				return
			}
			LogApp.Infoln(topicName, results)
		}
	}
	//进程退出
	LogApp.Infoln(topicName + " is exit")
}
