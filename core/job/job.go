package job

import (
	"fmt"
	"go-jobs/app/jobs/exportJob"
	"go-jobs/configs/constants"
	"go-jobs/helpers/util"
	"go-jobs/service/service_redis"
	"time"
)

type MyStruct struct {
}

/**
 * @Author: lizhenjian
 * @LastEditors: lizhenjian
 * @Description: 任务运行方法，通过反射调用业务方法
 * @param {string} PkgName
 * @param {string} topicName
 */
func Run(PkgName string, topicName string) {
	fmt.Println(topicName + " is running")
	//TODO动态注册业务方法
	util.RegisterFunction(PkgName, topicName, exportJob.HandleExportJob)
	i := 0
	for {
		//单进程最大执行任务次数
		i++
		if i >= constants.MaxPopNum {
			fmt.Println("MaxPopNum:", constants.MaxPopNum)
			break
		}
		//拉取Redis队列消息
		rdb := service_redis.NewClient()
		jobParams, err := rdb.RPop(topicName).Result()
		if err != nil {
			fmt.Println("Pop data error:", err, topicName)
			//队列没有消息停留一秒
			time.Sleep(time.Duration(constants.Sleep) * time.Second)
		}
		if jobParams != "" {
			fmt.Println("Popped data:", jobParams, topicName)
			//通过包名和方法名动态调用业务方法
			funcName := topicName
			args := []interface{}{jobParams}
			results, err := util.Eval(PkgName, funcName, args...)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(results)
		}
	}
	//进程退出
	fmt.Println(topicName + " is exit")
}
