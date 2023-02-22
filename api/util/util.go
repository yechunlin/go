package util

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
	"strconv"
	"time"

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

//GetNowDatetime
/**
获取当前日期与时间
*/
func GetNowDatetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 字符串转int8
func StrToInt8(str string) int {
	intVal, _ := strconv.Atoi(str)
	return intVal
}

// 获取随机字符串
func GetRandStr(len int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	obj := []string{str}
	var result string
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(len)))
	for i := 0; i < len; i++ {
		result += obj[rand.Intn(len)]
	}
	return result
}

// 随机号码
func GetRandMobileNumber() string {
	number := []string{"0123456789"}
	sortNumber := []string{"356789"}
	mobile := "1"
	rand.Seed(time.Now().UnixNano())
	mobile += sortNumber[rand.Intn(len(sortNumber))]
	for i := 0; i < 9; i++ {
		mobile += number[rand.Intn(10)]
	}
	return mobile
}
