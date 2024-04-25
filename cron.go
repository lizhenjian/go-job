package main

import (
	"go-jobs/configs/constants"
	"go-jobs/core/job"
	"go-jobs/helpers/loger"
	"go-jobs/service/service_redis"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/docker/docker/pkg/reexec"
	"github.com/sirupsen/logrus"
)

func init() {
	//初始化日志
	loger.InitLogger()
	logrus.Infoln("init start, os.Args = %+v\n", os.Args)
	// 读取配置文件中所有的topic
	for _, topic := range constants.Topics {
		// 注册所有的topic
		reexec.Register("staticProcess:"+topic.Name, func() {
			job.Run(topic.PkgName, topic.Name)
		})
		reexec.Register("dynamicProcess:"+topic.Name, func() {
			job.Run(topic.PkgName, topic.Name)
		})
		if reexec.Init() {
			logrus.Infoln("reexec.Init Exit:" + topic.Name)
			os.Exit(0)
		}
	}
}

/**
 * @Author: lizhenjian
 * @LastEditors: lizhenjian
 * @Description: 任务启动入口
 */
func main() {
	logrus.Infoln("main start, os.Args = %+v\n", os.Args)
	//执行定时任务
	timer1 := time.NewTicker(time.Duration(constants.TimerProcessCheck) * time.Second)
	timer2 := time.NewTicker(time.Duration(constants.TimerListCheck) * time.Second)
	for {
		select {
		case <-timer1.C:
			go func() {
				timerProcessCheck()
			}()
		case <-timer2.C:
			go func() {
				timerListCheck()
			}()
		}
	}
}

// 定时拉起static进程
func timerProcessCheck() {
	logrus.Infoln("timerProcessCheck---------------------------->")
	for _, topic := range constants.Topics {
		staticProcessCount, _ := countProcessInstances("staticProcess:" + topic.Name)
		if staticProcessCount < topic.WorkerMinNum {
			var wg sync.WaitGroup
			for i := 0; i < int(topic.WorkerMinNum); i++ {
				wg.Add(1)
				go func(topic constants.Topic) {
					defer wg.Done()
					cmd := reexec.Command("staticProcess:" + topic.Name)
					cmd.Stdin = os.Stdin
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Start(); err != nil {
						log.Panicf("failed to run command: %s", err)
					}
					logrus.Infoln("timerProcessCheck 一次性拉起static进程", "childProcess:"+topic.Name, cmd.Process.Pid)
					err := cmd.Wait()
					if err != nil {
						log.Panicf("timerProcessCheck failed to wait command: %s", err)
					}
				}(topic)
			}
			wg.Wait()
		}
		logrus.Infoln("timerProcessCheck finished successfully:", topic)
	}
	logrus.Infoln("timerProcessCheck All commands finished.")
}

// 定时检查队列长度拉起动态进程
func timerListCheck() {
	logrus.Infoln("timerListCheck---------------------------->")
	rdb := service_redis.NewClient()
	//检查redis队列长度
	for _, topic := range constants.Topics {
		length, _ := rdb.LLen(topic.Name).Result()
		logrus.Infoln("长度:", topic.Name, length)
		if topic.WorkerMaxPendingLength < length {
			//检查已有进程个数
			processCount, _ := countProcessInstances(topic.Name)
			if processCount < topic.WorkerMaxNum {
				var wg sync.WaitGroup
				for i := 0; i < int(topic.WorkerMaxNum); i++ {
					wg.Add(1)
					go func(topic constants.Topic) {
						defer wg.Done()
						cmd := reexec.Command("dynamicProcess:" + topic.Name)
						cmd.Stdin = os.Stdin
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						if err := cmd.Start(); err != nil {
							log.Panicf("failed to run command: %s", err)
						}
						logrus.Infoln("timerListCheck 一次性拉起多个进程", "dynamicProcess:"+topic.Name, cmd.Process.Pid)
						err := cmd.Wait()
						if err != nil {
							log.Panicf("timerListCheck failed to wait command: %s", err)
						}
						logrus.Infoln("timerListCheck finished successfully:", topic)
					}(topic)
				}
				wg.Wait()
				logrus.Infoln("timerListCheck All commands finished.")
			}
		}
	}
}

func isProcessRunning(processName string) bool {
	cmd := exec.Command("pgrep", processName)
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	// If the output is not empty, then the process is running
	return len(strings.TrimSpace(string(output))) > 0
}

func countProcessInstances(processName string) (int64, error) {
	cmd := exec.Command("pgrep", "-f", processName)
	output, err := cmd.Output()

	if err != nil {
		return 0, err
	}
	// Count the number of lines in the output
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	//转为int64
	lenth := int64(len(lines))
	return lenth, nil
}
