package util

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
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
func GetRandStr(l int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	obj := []byte(str)
	obj_len := len(obj)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(obj_len)))
	for i := 0; i < l; i++ {
		result = append(result, obj[rand.Intn(obj_len)])
	}
	return string(result)
}

// 随机号码
func GetRandMobileNumber() string {
	number := []byte("0123456789")
	sortNumber := []byte("356789")
	mobile := []byte("1")
	rand.Seed(time.Now().UnixNano())
	mobile = append(mobile, sortNumber[rand.Intn(len(sortNumber))])
	for i := 0; i < 9; i++ {
		mobile = append(mobile, number[rand.Intn(10)])
	}
	return string(mobile)
}

// 获取两个整数之间的随机数
func GetRandTwo(min int, max int) int {
	off := max - min
	rand.Seed(time.Now().UnixNano())
	tmp := rand.Intn(off)
	return tmp + min
}

// 类型转int
func GetInterfaceToInt(t1 interface{}) int {
	var t2 int
	switch t1.(type) {
	case uint:
		t2 = int(t1.(uint))
		break
	case int8:
		t2 = int(t1.(int8))
		break
	case uint8:
		t2 = int(t1.(uint8))
		break
	case int16:
		t2 = int(t1.(int16))
		break
	case uint16:
		t2 = int(t1.(uint16))
		break
	case int32:
		t2 = int(t1.(int32))
		break
	case uint32:
		t2 = int(t1.(uint32))
		break
	case int64:
		t2 = int(t1.(int64))
		break
	case uint64:
		t2 = int(t1.(uint64))
		break
	case float32:
		t2 = int(t1.(float32))
		break
	case float64:
		t2 = int(t1.(float64))
		break
	case string:
		t2, _ = strconv.Atoi(t1.(string))
		if t2 == 0 && len(t1.(string)) > 0 {
			f, _ := strconv.ParseFloat(t1.(string), 64)
			t2 = int(f)
		}
		break
	case nil:
		t2 = 0
		break
	case json.Number:
		t3, _ := t1.(json.Number).Int64()
		t2 = int(t3)
		break
	default:
		t2 = t1.(int)
		break
	}
	return t2
}
