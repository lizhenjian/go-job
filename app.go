package main

import (
	index "go-jobs/app/controller/index"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//定义路由的GET方法及响应处理函数
	r.GET("/add_job", index.AddTestJob)
	r.Run() //默认在本地8080端口启动服务
}
