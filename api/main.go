package main

import (
	"api/middleware/logger"
	"api/route"
	"github.com/fvbock/endless"
	"syscall"
	"time"
)

func main() {
	//初始化路由，启动服务
	Route := route.RouteInit()
	server := endless.NewServer(":8081", Route)
	server.ReadHeaderTimeout = 3 * 60 * time.Second
	server.WriteTimeout = 3 * 60 * time.Second
	server.MaxHeaderBytes = 1 << 20
	server.BeforeBegin = func(add string) {
		logger.GetInstance().Infof("pid is:%3d", syscall.Getpid())
	}
	server.ListenAndServe()
}
