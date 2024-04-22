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