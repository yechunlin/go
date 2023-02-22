package main

import (
	"api/conf"
	"api/jobs"
	"api/model"
	"api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	//设置模式
	gin.SetMode(conf.APP_ENV)
	//初始化数据库
	model.ModelInit()
	//初始化路由，启动服务
	route.RouteInit()
	//初始化jobs
	jobs.InitJobs()
}
