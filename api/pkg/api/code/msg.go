package code

var Message = map[int]string{
	Success: "OK",

	ErrorRequest:  "请求参数错误",
	ErrorPanic:    "系统报错",
	ErrorSign:     "权限验证失败",
	ErrorDatabase: "数据库异常",
	ErrorSyntax:   "动态语法错误",

	ErrorTeamHasApp:      "团队在维护APP，无法删除",
	ErrorTeamMemberExist: "该用户已经在团队中",

	ErrorAppFileParserFail: "解析失败",

	ErrorMemberNotExists: "会员不存在",
}

func GetMessage(code int) string {
	if message, ok := Message[code]; ok {
		return message
	}

	return "未知错误"
}
