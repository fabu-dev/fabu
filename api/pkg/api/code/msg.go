package code

var Message = map[int]string{
	Success: "OK",

	ErrorRequest:  "请求参数错误",
	ErrorPanic:    "系统报错",
	ErrorSign:     "签名错误",
	ErrorDatabase: "数据库异常",
	ErrorSyntax:   "动态语法错误",
}

func GetMessage(code int) string {
	if message, ok := Message[code]; ok {
		return message
	}

	return "未知错误"
}
