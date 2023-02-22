package route

import (
	"api/conf"
	"api/controller"
	"api/middleware/auth"
	"api/middleware/loggo"
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
	r.Use(recover2.MidError(), loggo.FormateLogger())

	user := r.Group("/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}
	user.Use(auth.AuthApi())
	{
		user.POST("/getInfo", controller.GetInfo)
		user.POST("/userList", controller.GetUserList)
		user.POST("/updateUser", controller.UpdateUser)
	}
	return r
}

// 启动服务
func RouteInit() {
	//routeConf().Run()
	server := endless.NewServer(conf.SERVER_ADDRESS, routeConf())
	server.BeforeBegin = func(add string) {
		util.FilePutContents(conf.PID_FILE, strconv.Itoa(syscall.Getpid()), os.O_CREATE|os.O_RDWR)
	}
	loggo.WriteLogStr("服务启动")
	server.ListenAndServe()
}
