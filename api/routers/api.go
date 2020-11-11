package routers

import (
	"fabu.dev/api/controller"
	"fabu.dev/api/pkg/api/middleware"
)

func CreateApiRouter() {

	// 会员API路由配置
	v1Member := Router.Group("/v1/user").Use(middleware.VerifyToken())
	{
		v1Member.GET("/info/:id", controller.NewMember().Info)           // 会员详情
		v1Member.PUT("/password", controller.NewMember().UpdatePassword) // 修改密码
		v1Member.POST("/token", controller.NewMember().CreateToken)      // 生成TOKEN
	}

	// 团队API路由配置
	v1Team := Router.Group("/v1/team").Use(middleware.VerifyToken())
	{
		v1Team.GET("/", controller.NewTeam().GetList)                   // 会员团队列表
		v1Team.POST("/create", controller.NewTeam().Create)             // 创建团队
		v1Team.PUT("/edit", controller.NewTeam().Edit)                  // 编辑团队
		v1Team.GET("/info/:id", controller.NewTeam().View)              // 团队 详情
		v1Team.POST("/member/add", controller.NewTeam().AddMember)      // 邀请成员
		v1Team.DELETE("/member/del", controller.NewTeam().DeleteMember) // 移除成员
		v1Team.DELETE("/del", controller.NewTeam().Delete)              // 解散团队
		v1Team.GET("/member/:id", controller.NewTeam().GetMember)       // 获取团队成员信息
		v1Team.GET("/log", controller.NewTeam().GetLog)                 // 团队日志信息
	}

	// 登录认证API路由配置
	v1Auth := Router.Group("/v1/auth")
	{
		v1Auth.POST("/login", controller.NewAuth().Login)       // 登录
		v1Auth.POST("/2step-code", controller.NewAuth().Login)  // 登录
		v1Auth.POST("/logout", controller.NewAuth().Logout)     // 登出
		v1Auth.POST("/register", controller.NewAuth().Register) // 注册
		v1Auth.GET("/forget", controller.NewAuth().Forget)      // 忘记密码
	}

	// APP管理
	v1App := Router.Group("/v1/app").Use(middleware.VerifyToken())
	{
		v1App.GET("/", controller.NewApp().GetList)                     // APP 列表
		v1App.POST("/upload", controller.NewApp().Upload)               // APP 上传文件
		v1App.POST("/create", controller.NewApp().Save)                 // APP 保存上传信息
		v1App.GET("/info/:id", controller.NewApp().View)                // APP 详情
		v1App.PUT("/edit", controller.NewApp().Edit)                    // APP 编辑
		v1App.DELETE("/delete", controller.NewApp().Delete)             // APP 删除
		v1App.POST("/combine", controller.NewApp().Combine)             // APP 合并
		v1App.GET("/stat", controller.NewApp().GetStat)                 // APP 统计
		v1App.GET("/log", controller.NewApp().GetLog)                   // APP 统计
		v1App.POST("/base", controller.NewApp().GetAppInfoByIdentifier) // APP 统计

	}

	// APP版本管理
	v1AppVersion := Router.Group("/v1/app/version").Use(middleware.VerifyToken())
	{
		v1AppVersion.GET("/", controller.NewAppVersion().GetList)                    // APP 版本列表
		v1AppVersion.DELETE("/delete", controller.NewAppVersion().Delete)            // APP 版本删除
		v1AppVersion.GET("/download", controller.NewAppVersion().Download)           // APP 版本下载
		v1AppVersion.GET("/info/:id", controller.NewAppVersion().View)               // APP 版本详情
		v1AppVersion.POST("/publish", controller.NewAppVersion().Publish)            // APP 版本发布
		v1AppVersion.POST("/cancel", controller.NewAppVersion().Cancel)              // APP 版本取消发布
		v1AppVersion.POST("/edit", controller.NewAppVersion().Edit)                  // APP 版本编辑
		v1AppVersion.GET("/download/log", controller.NewAppVersion().GetDownloadLog) // APP 版本下载记录
	}

	// 系统API路由配置
	system := Router.Group("/system").Use(middleware.VerifyToken())
	{
		system.GET("/health", controller.NewSystem().Health)    // 监控检查
		system.POST("/setting", controller.NewSystem().Setting) // 监控检查
	}
}
