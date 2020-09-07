package main

import (
	"go-web-frame/pkg/db"
	"go-web-frame/pkg/web"
)

func init() {

}

// 入口文件，main函数
func main() {
	// gin.SetMode(setting.ServerSetting.RunMode)

	web.StartWebServer()
	destroy()
}

func destroy() {
	db.Destroy() // 关闭数据库连接
}
