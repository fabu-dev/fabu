package routers

import "go-web-frame/controller"

func CreateApiRouter() {

	// 会员API路由配置
	v1Member := Router.Group("/v1/member")
	{
		v1Member.GET("/info/:id", controller.NewMember().View)           // 会员详情
		v1Member.POST("/login", controller.NewMember().Login)            // 登录
		v1Member.POST("/register", controller.NewMember().Register)      // 注册
		v1Member.GET("/group/:id", controller.NewMember().GroupView)     // 获取会员的团队信息
		v1Member.PUT("/password", controller.NewMember().UpdatePassword) // 会员详情
	}

	// 系统API路由配置
	system := Router.Group("/system")
	{
		system.GET("/health", controller.NewSystem().Health) // 监控检查
	}
}
