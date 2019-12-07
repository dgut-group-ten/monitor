package main

import (
	"monitor/core/conf"
	"monitor/core/handler"
	"monitor/core/util"
	"net/http"
	"os"
)

func main() {
	// 打日志
	logFd, err := os.OpenFile(conf.LogFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		util.Log.Out = logFd
		defer logFd.Close()
	}
	util.Log.Infof("Exec start.\n")
	util.Log.Infof("Welcome to log monitor.\n")

	// 网页
	http.HandleFunc("/", handler.HelloHandler)

	err = http.ListenAndServe(":8004", nil)
	if err != nil {
		util.Log.Panicln("Failed to start server , err: " + err.Error())
	}
}
