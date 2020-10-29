package utils

// 处理panic，将报错信息，发生到企业微信报警
func HandlePanic() {
	if panicInfo := recover(); panicInfo != nil {
		panicInfo, ok := panicInfo.(string)
		if !ok {
			panicInfo = "出错了！"
		}

		//将错误信息加入企业微信报警
		message := make(map[string]interface{})
		message["message"] = panicInfo

		// TODO 发送邮件
	}
}
