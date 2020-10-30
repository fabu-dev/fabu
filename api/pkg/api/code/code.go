package code

const (
	Success = 1

	// 通用的错误提示
	ErrorRequest  = 10001 // 请求参数错误
	ErrorPanic    = 10002 // 系统报错
	ErrorSign     = 10003 // 签名错误
	ErrorDatabase = 10004 // 数据库错误
	ErrorSyntax   = 10005 // 动态语法错误

	// team模块
	ErrorTeamHasApp      = 20001 // 团队还在维护APP
	ErrorTeamMemberExist = 20002 // 用户已经在团队中

	// 会员模块
	ErrorMemberNotExists = 30001 // 会员不存在

	// app模块
	ErrorAppFileParserFail = 40001 // app解析失败
)
