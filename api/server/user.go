package server

import (
	"api/conf"
	"api/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context, param map[string]any) {
	token := fmt.Sprintf("%f", param["user"]) + fmt.Sprintf("%f", param["password"])
	param["token"] = util.MD5(token)
	util.ReturnData(c, conf.API_SERVER_SUCCESS, "you are logged in", param)
}
