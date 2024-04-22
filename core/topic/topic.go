package topic

import (
	"encoding/json"
	"fmt"
	"go-jobs/service/service_redis"
)

/**
 * @Author: lizhenjian
 * @LastEditors: lizhenjian
 * @Description: 任务消息推送方法
 * @param {string} jobName
 * @param {string} jobFunc
 * @param {string} jobData
 */
func Push(jobName string, jobFunc string, jobData string) (resp []interface{}, err error) {
	if jobName == "" {
		return nil, fmt.Errorf("jobName is empty")
	}
	if jobFunc == "" {
		return nil, fmt.Errorf("jobFunc is empty")
	}
	if jobData == "" {
		return nil, fmt.Errorf("jobData is empty")
	}
	//redis list push
	jobParams := map[string]string{
		"jobName": jobName,
		"jobFunc": jobFunc,
		"jobData": jobData,
	}
	jobParamsStr, _ := json.Marshal(jobParams)
	rdb := service_redis.NewClient()
	err = rdb.LPush(jobName, jobParamsStr).Err()
	return
}
