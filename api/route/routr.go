package route

import (
	"api/controller"
	"api/middleware/auth"
	"api/middleware/logger"
	recover2 "api/middleware/recover"
	"github.com/gin-gonic/gin"
)

func RouteInit() *gin.Engine {
	r := gin.New()
	r.Use(recover2.MidError(), logger.FormateLogger(), auth.BaseApi())

	user := r.Group("/user")
	{
		user.POST("/login", controller.Login)
	}
	return r
}
