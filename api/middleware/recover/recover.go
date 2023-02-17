package recover

import (
	"api/conf"
	"api/util"
	"github.com/gin-gonic/gin"
)

// 自定义中间件，异常错误处理
func MidError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				util.ReturnDataErr(c, conf.API_SERVER_ERROR, err, nil, 500)
				c.Abort()
			}
		}()
		c.Next()
	}
}
