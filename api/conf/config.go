package conf

// release 生产 | debug 开发 | test 测试

const APP_ENV = "debug"

// 常量配置
const (
	MYSQLADDRESS = "124.70.187.23" //数据库连接IP

	LOG_SAVE_DIR = "E:/go2/api/runtime/logs"
)

//MysqlConf
/**
数据库mysql
*/
type MysqlConf struct {
}
