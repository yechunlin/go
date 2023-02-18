package route

import (
	"api/conf"
	"api/controller"
	"api/middleware/auth"
	"api/middleware/logger"
	recover2 "api/middleware/recover"
	"api/util"
	"os"
	"strconv"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func routeConf() *gin.Engine {
	r := gin.New()
	r.Use(recover2.MidError(), logger.FormateLogger(), auth.BaseApi())

	user := r.Group("/user")
	{
		user.POST("/login", controller.Login)
	}

	return r
}

// 启动服务
func RouteInit() {
	server := endless.NewServer(conf.SERVER_ADDRESS, routeConf())
	server.BeforeBegin = func(add string) {
		util.FilePutContents(conf.PID_FILE, strconv.Itoa(syscall.Getpid()), os.O_CREATE|os.O_RDWR)
	}
	server.ListenAndServe()
}
