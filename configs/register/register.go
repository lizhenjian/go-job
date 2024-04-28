package register

import (
	"go-jobs/app/jobs/exportJob"
	"go-jobs/app/jobs/importJob"
	"go-jobs/helpers/util"
)

/**
 * @Author: lizhenjian
 * @LastEditors: lizhenjian
 * @Description: 注册任务方法配置
 */
func RegisterFunc() {
	util.RegisterFunction("exportJob", "HandleExportJob", exportJob.HandleExportJob)
	util.RegisterFunction("importJob", "HandleImportJob", importJob.HandleImportJob)
}
