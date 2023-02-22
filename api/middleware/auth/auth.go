package auth

import (
	"api/conf"
	"api/extend/jwter"
	"api/middleware/loggo"
	"api/util"
	"strings"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件
func AuthApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
		loggo.WriteLogStr(token)
		if token == "" {
			util.ReturnDataErr(c, conf.API_TOKEN_EMPTY, "请先登录", nil, 403)
			c.Abort()
			return
		}
		claims, err := jwter.ParseToken(token)
		if err != nil {
			util.ReturnDataErr(c, conf.API_TOKEN_EMPTY, err.Error(), nil, 403)
			c.Abort()
			return
		}
		c.Set("userId", claims.UserId)
		c.Next()
	}
}
