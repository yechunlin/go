package model

import (
	"api/conf"
	"context"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbInstance *gorm.DB

func MysqllInit() {
	//写入文件
	//file := loggo.OpenLogFile()
	//输出控制台
	writer := os.Stdout

	newLogger := logger.New(
		log.New(writer, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	db, err := gorm.Open(conf.DB_SELECT[conf.DB_DRIVER], &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true, //关闭默认事务处理
		PrepareStmt:            true, //缓存预编译
	})
	if err != nil {
		panic("failed to connect database")
	}
	DbInstance = db
	db.Config.Logger.Info(context.Background(), "数据库初始化成功")
}
