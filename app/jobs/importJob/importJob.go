package importJob

import (
	"fmt"
	"time"
)

// 测试导入任务方法
func HandleImportJob(data string) {
	fmt.Println("--------------HandleImportJob--------------")
	fmt.Println(data)
	fmt.Println("--------------HandleImportJob--------------")
	time.Sleep(1 * time.Second)
}
