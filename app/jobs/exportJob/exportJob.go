package exportJob

import (
	"fmt"
	"time"
)

// 测试导出任务方法
func HandleExportJob(data string) {
	fmt.Println("--------------HandleExportJob--------------")
	fmt.Println(data)
	fmt.Println("--------------HandleExportJob--------------")
	time.Sleep(1 * time.Second)
}
