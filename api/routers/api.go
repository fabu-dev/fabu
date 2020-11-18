package routers

import (
	"net/http"

	controller2 "fabu.dev/api/application/controller"

	"fabu.dev/api/application/middleware"
)

func CreateApiRouter() {

	// 会员API路由配置
	v1Member := Router.Group("/v1/user").Use(middleware.VerifyAuth())
	{
		v1Member.GET("/info/:id", controller2.NewMember().Info)           // 会员详情
		v1Member.PUT("/password", controller2.NewMember().UpdatePassword) // 修改密码
		v1Member.POST("/token", controller2.NewMember().CreateToken)      // 生成TOKEN
	}

	// 团队API路由配置
	v1Team := Router.Group("/v1/team").Use(middleware.VerifyAuth())
	{
		v1Team.GET("/", controller2.NewTeam().GetList)                   // 会员团队列表
		v1Team.POST("/create", controller2.NewTeam().Create)             // 创建团队
		v1Team.PUT("/edit", controller2.NewTeam().Edit)                  // 编辑团队
		v1Team.GET("/info/:id", controller2.NewTeam().View)              // 团队 详情
		v1Team.POST("/member/add", controller2.NewTeam().AddMember)      // 邀请成员
		v1Team.DELETE("/member/del", controller2.NewTeam().DeleteMember) // 移除成员
		v1Team.DELETE("/del", controller2.NewTeam().Delete)              // 解散团队
		v1Team.GET("/member/:id", controller2.NewTeam().GetMember)       // 获取团队成员信息
		v1Team.GET("/log", controller2.NewTeam().GetLog)                 // 团队日志信息
	}

	// 登录认证API路由配置
	v1Auth := Router.Group("/v1/auth")
	{
		v1Auth.POST("/login", controller2.NewAuth().Login)       // 登录
		v1Auth.POST("/2step-code", controller2.NewAuth().Login)  // 登录
		v1Auth.POST("/logout", controller2.NewAuth().Logout)     // 登出
		v1Auth.POST("/register", controller2.NewAuth().Register) // 注册
		v1Auth.GET("/forget", controller2.NewAuth().Forget)      // 忘记密码
	}

	// APP管理
	v1App := Router.Group("/v1/app").Use(middleware.VerifyAuth())
	{
		v1App.GET("/", controller2.NewApp().GetList)                     // APP 列表
		v1App.POST("/upload", controller2.NewApp().Upload)               // APP 上传文件
		v1App.POST("/create", controller2.NewApp().Save)                 // APP 保存上传信息
		v1App.GET("/info/:id", controller2.NewApp().View)                // APP 详情
		v1App.PUT("/edit", controller2.NewApp().Edit)                    // APP 编辑
		v1App.DELETE("/delete", controller2.NewApp().Delete)             // APP 删除
		v1App.POST("/combine", controller2.NewApp().Combine)             // APP 合并
		v1App.GET("/stat", controller2.NewApp().GetStat)                 // APP 统计
		v1App.GET("/log", controller2.NewApp().GetLog)                   // APP 统计
		v1App.POST("/base", controller2.NewApp().GetAppInfoByIdentifier) // APP 统计

	}

	// APP版本管理
	v1AppVersion := Router.Group("/v1/app/version").Use(middleware.VerifyAuth())
	{
		v1AppVersion.GET("/", controller2.NewAppVersion().GetList)           // APP 版本列表
		v1AppVersion.DELETE("/delete", controller2.NewAppVersion().Delete)   // APP 版本删除
		v1AppVersion.GET("/info/:id", controller2.NewAppVersion().View)      // APP 版本详情
		v1AppVersion.POST("/publish", controller2.NewAppVersion().Publish)   // APP 版本发布
		v1AppVersion.POST("/cancel", controller2.NewAppVersion().Cancel)     // APP 版本取消发布
		v1AppVersion.POST("/edit", controller2.NewAppVersion().Edit)         // APP 版本编辑
		v1AppVersion.GET("/log", controller2.NewAppVersion().GetDownloadLog) // APP 版本下载记录
	}

	// 系统API路由配置
	system := Router.Group("/system").Use(middleware.VerifyAuth())
	{
		system.GET("/health", controller2.NewSystem().Health)    // 监控检查
		system.POST("/setting", controller2.NewSystem().Setting) // 监控检查
	}

	Router.GET("/download/:code", controller2.NewAppVersion().Download) // APP 下载

	Router.StaticFS("/static/app", http.Dir("./static/app"))
}
