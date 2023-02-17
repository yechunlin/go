package auth

import (
	"api/conf"
	"api/util"
	"github.com/gin-gonic/gin"
)

func BaseApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			util.ReturnDataErr(c, conf.API_TOKEN_EMPTY, "请先登录", nil, 403)
			c.Abort()
		}
		c.Next()
	}
}
