package exportJob

import (
	"fmt"
	"time"
)

// 自定义任务方法都写在这个jobs目录下面
func HandleExportJob(data string) {
	fmt.Println("HandleExportJob___________________________")
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
