package index

import (
	"encoding/json"
	"fmt"
	"go-jobs/core/topic"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func AddTestJob(c *gin.Context) {
	//循环添加10个测试任务
	for i := 0; i < 100; i++ {
		jobData := map[string]interface{}{
			"file": "fileName",
			"id":   i,
		}
		jobDataStr, _ := json.Marshal(jobData)
		topic.Push("HandleExportJob", "HandleExportJob", string(jobDataStr))
	}
	fmt.Println("AddTestJob")
	c.JSON(200, gin.H{
		"message": "AddTestJob",
	})
}
