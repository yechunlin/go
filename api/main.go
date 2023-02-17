package main

import (
	"api/middleware/logger"
	"api/route"
)

func main() {
	//初始化路由，启动服务
	logger.WriteLogStr("初始化路由,启动服务")
	route.RouteInit()
}
