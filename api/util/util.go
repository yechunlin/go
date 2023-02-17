package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

//ReturnData
/**
api统一返回
*/
func ReturnData(c *gin.Context, code int16, msg any, data interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

//ReturnDataErr
/**
错误异常调用返回
*/
func ReturnDataErr(c *gin.Context, code int16, msg any, data interface{}, httpCode int) {
	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// MD5
/**
加密
*/
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
