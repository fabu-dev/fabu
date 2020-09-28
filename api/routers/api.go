package routers

import "fabu.dev/api/controller"

func CreateApiRouter() {

	// 会员API路由配置
	v1Member := Router.Group("/v1/user")
	{
		v1Member.GET("/info/:id", controller.NewMember().Info)           // 会员详情
		v1Member.GET("/group/:id", controller.NewMember().GroupView)     // 获取会员的团队信息
		v1Member.PUT("/password", controller.NewMember().UpdatePassword) // 会员详情
	}

	v1Auth := Router.Group("/v1/auth")
	{

		v1Auth.POST("/login", controller.NewAuth().Login)            // 登录
		v1Auth.POST("/2step-code", controller.NewAuth().Login)            // 登录
		v1Auth.POST("/logout", controller.NewAuth().Logout)           // 登出
		v1Auth.POST("/register", controller.NewAuth().Register)      // 注册
		v1Auth.GET("/forget", controller.NewAuth().Forget)     // 忘记密码
	}

	// 系统API路由配置
	system := Router.Group("/system")
	{
		system.GET("/health", controller.NewSystem().Health) // 监控检查
	}
}
