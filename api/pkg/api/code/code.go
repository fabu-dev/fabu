package code

const (
	Success = 1

	// 通用的错误提示
	ErrorRequest  = 10001 // 请求参数错误
	ErrorPanic    = 10002 // 系统报错
	ErrorSign     = 10003 // 签名错误
	ErrorDatabase = 10004 // 数据库错误
	ErrorSyntax   = 10005 // 动态语法错误
)
