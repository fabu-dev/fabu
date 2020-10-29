package main

import (
	"fabu.dev/api/job"
	"fabu.dev/api/pkg/db"
	"fabu.dev/api/pkg/web"
)

func init() {

}

// 入口文件，main函数
func main() {
	// gin.SetMode(setting.ServerSetting.RunMode)

	job.NewCommand().Process()

	web.StartWebServer()
	destroy()

}

func destroy() {
	db.Destroy() // 关闭数据库连接
}
