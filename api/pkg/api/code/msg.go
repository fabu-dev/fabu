package code

var Message = map[int]string{
	Success: "OK",

	ERROR_REQUEST_ERROR: "请求参数错误",
	ERROR_PANIC:         "系统报错",
	ERROR_SIGN:          "签名错误",
}

func GetMessage(code int) string {
	if message, ok := Message[code]; ok {
		return message
	}

	return "未知错误"
}
