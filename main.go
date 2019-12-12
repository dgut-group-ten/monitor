package main

import (
	"github.com/robfig/cron/v3"
	"monitor/core/handler"
	"monitor/core/task"
	"net/http"
)

func main() {

	// 定时任务
	c := cron.New()
	_, _ = c.AddFunc("*/30 * * * *", task.UpdateUserOperation)
	_, _ = c.AddFunc("*/50 * * * *", task.UpdatePVUV)
	c.Start()
	defer c.Stop()

	// 网页路由
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/vc/resource/", handler.VisitCountHandler)
	http.HandleFunc("/uo/history/", handler.HistoryHandler)

	err := http.ListenAndServe(":8004", nil)
	if err != nil {
		panic("Failed to start server , err: " + err.Error())
	}
}
