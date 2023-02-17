package controller

import (
	"api/conf"
	"api/server"
	"api/util"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var param = map[string]any{
		"user":     c.DefaultPostForm("user", ""),
		"password": c.DefaultPostForm("password", ""),
	}
	if param["user"] == "" || param["password"] == "" {
		util.ReturnData(c, conf.API_ACCOUNT_ERROR, "账号有误", nil)
		return
	}
	server.Login(c, param)
}
