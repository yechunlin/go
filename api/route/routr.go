package route

import (
	"api/controller"
	"api/middleware/auth"
	"api/middleware/logger"
	recover2 "api/middleware/recover"
	"api/util"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"syscall"
	"time"
)

func RouteInit() {
	r := gin.New()
	r.Use(recover2.MidError(), logger.FormateLogger(), auth.BaseApi())

	user := r.Group("/user")
	{
		user.POST("/login", controller.Login)
	}
	server := endless.NewServer(":8081", r)
	server.ReadHeaderTimeout = 3 * 60 * time.Second
	server.WriteTimeout = 3 * 60 * time.Second
	server.MaxHeaderBytes = 1 << 20
	server.BeforeBegin = func(add string) {
		util.FilePutContents("pid.txt", strconv.Itoa(syscall.Getpid()), os.O_CREATE|os.O_RDWR)
	}
	server.ListenAndServe()
}
