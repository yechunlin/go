package loggo

import (
	"api/conf"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logFileName = time.Now().Format("2006-01-02") + ".log"

// 获取文件句柄
func OpenLogFile() *os.File {
	logFile := path.Join(conf.LOG_SAVE_DIR, logFileName)
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return file
}

// 日志追踪 info
func WriteLogStr(str string) {
	file := OpenLogFile()
	logger := logrus.New()
	logger.SetOutput(file)
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.Info(str)
}

// 自定义hook
type myHook struct {
}

func (h *myHook) Fire(e *logrus.Entry) error {
	//e.Data["aaa"] = "xxxxxxxx"
	return nil
}

func (h *myHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// 自定义日志中间件
func FormateLogger() gin.HandlerFunc {
	file := OpenLogFile()
	logger := logrus.New()
	logger.SetOutput(file)
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//logger.AddHook(&myHook{})
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		//代理
		agent := c.Request.UserAgent()

		//formData数据
		postForm := c.Request.PostForm.Encode()

		// 日志格式
		//logger.WithFields(logrus.Fields{
		//	"statusCode":  statusCode,
		//	"latencyTime": latencyTime,
		//	"clientIP":    clientIP,
		//	"reqMethod":   reqMethod,
		//	"reqUri":      reqUri,
		//}).Info()
		logger.Infof("[client_IP: %15s] [method: %s] [httpCode: %3d]  [reqUri: %s] [latencyTime: %13v] [agent: %s] [postForm: %s]",
			clientIP,
			reqMethod,
			statusCode,
			reqUri,
			latencyTime,
			agent,
			postForm,
		)
	}
}
