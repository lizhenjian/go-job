# go-job

高性能任务队列系统 go-job

<br/>

- 启动任务进程：

```
go run cron.go
```

- 启动测试WEB服务：

```
go run app.go
```

- 向任务推送测试任务：

```
curl '127.0.0.1:8080/add_job'
```

- 项目结构

```
.
├── README.md
├── app
│   ├── controller
│   │   └── index
│   │       └── index.go      //测试web接口
│   └── jobs                  //任务都写写到这个目录下
│       ├── exportJob
│       │   └── exportJob.go
│       └── importJob
│           └── importJob.go
├── app.go
├── config.yaml              //配置文件
├── configs
│   ├── constants
│   │   └── constants.go     //任务参数设置
│   └── register
│       └── register.go      //任务注册配置文件
├── core
│   ├── job
│   │   └── job.go           //任务运行方法
│   └── topic
│       └── topic.go         //任务消息推送方法
├── cron.go
├── go.mod
├── go.sum
├── helpers
│   ├── loger
│   │   └── loger.go
│   └── util
│       └── util.go
├── logs
│   ├── application.log
│   ├── error.log
│   └── process.log
├── recive.go
├── send.go
└── service
    └── service_redis
        └── service_redis.go
```
运行截图：
![image](https://github.com/lizhenjian/go-job/assets/8486305/5f6b42b3-7ad4-45e7-b366-277c1f8781eb)


体系架构图
![image](https://github.com/lizhenjian/go-job/assets/8486305/b12119ed-aa2a-44c0-a901-f3c4b490f550)
