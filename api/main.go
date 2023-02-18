package main

import (
	"api/conf"
	"api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	//设置模式
	gin.SetMode(conf.APP_ENV)
	//初始化路由，启动服务
	route.RouteInit()
}
