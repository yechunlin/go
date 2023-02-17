package conf

const (
	API_SERVER_SUCCESS = 0     //服务正常
	API_SERVER_ERROR   = 10000 //服务异常错误码
	API_ACCOUNT_ERROR  = 10001 //账号密码登录异常
	API_TOKEN_EMPTY    = 10002 //token为空或未传
)
