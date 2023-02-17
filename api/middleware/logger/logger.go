package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

const logFileDir = "E:/go2/api/runtime/logs"

var logFileName = time.Now().Format("2006-01-02") + ".log"

func GetInstance() *logrus.Logger {
	logFile := path.Join(logFileDir, logFileName)
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND, os.ModeAppend)
	if err != nil {

	}
	logger := logrus.New()
	logger.SetOutput(file)
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
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

func FormateLogger() gin.HandlerFunc {
	logFile := path.Join(logFileDir, logFileName)
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND, os.ModeAppend)
	if err != nil {

	}
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
