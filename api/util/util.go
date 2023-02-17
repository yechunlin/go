package util

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"os"
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

//FilePutContents
/**
写入文件
*/
func FilePutContents(file string, content string, flag int) {
	f, _ := os.OpenFile(file, flag, 0777)
	defer f.Close()
	writer := bufio.NewWriter(f)
	writer.WriteString(content)
	writer.Flush()
}
