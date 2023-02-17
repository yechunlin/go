package main

import (
	"api/conf"
	"api/middleware/logger"
	"api/route"
	"api/util"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"syscall"
)

func main() {
	//设置模式
	gin.SetMode(conf.APP_ENV)
	//初始化路由，启动服务
	logger.WriteLogStr("初始化路由,启动服务")
	server := endless.NewServer(":8081", route.RouteInit())
	server.BeforeBegin = func(add string) {
		util.FilePutContents("pid.txt", strconv.Itoa(syscall.Getpid()), os.O_CREATE|os.O_RDWR)
	}
	server.ListenAndServe()
}
