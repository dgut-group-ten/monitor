package main

import (
	"github.com/robfig/cron/v3"
	"monitor/core/handler"
	"monitor/core/task"
	"monitor/core/util"
	"net/http"
)

func main() {
	// 打日志
	util.Log.Infof("Exec start.\n")
	util.Log.Infof("Welcome to log monitor.\n")

	// 定时任务
	c := cron.New()                           // 定义一个cron运行器
	c.AddFunc("*/5 * * * * *", task.Print5s)  // 定时5秒，每5秒执行print5
	c.AddFunc("*/15 * * * * *", task.Print5m) // 定时15秒，每5秒执行print5
	c.Start()                                 // 开始
	defer c.Stop()

	// 网页路由
	http.HandleFunc("/", handler.HelloHandler)

	err := http.ListenAndServe(":8004", nil)
	if err != nil {
		util.Log.Panicln("Failed to start server , err: " + err.Error())
	}
}
