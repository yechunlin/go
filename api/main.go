package main

import (
	"api/middleware/logger"
	"api/route"
	"api/util"
	"github.com/fvbock/endless"
	"os"
	"strconv"
	"syscall"
)

func main() {
	//初始化路由，启动服务
	logger.WriteLogStr("初始化路由,启动服务")
	server := endless.NewServer(":8081", route.RouteInit())
	server.BeforeBegin = func(add string) {
		util.FilePutContents("pid.txt", strconv.Itoa(syscall.Getpid()), os.O_CREATE|os.O_RDWR)
	}
	server.ListenAndServe()
}
