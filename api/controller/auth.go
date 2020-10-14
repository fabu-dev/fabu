package controller

import (
	"fabu.dev/api/filter"
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

// @Tags 用户登录注册相关接口
// @Summary 登录 API
// @Description 登录
// @Success 200 {string} string    "ok"
// @Router /v1/auth/login [POST]
func (ctl *Auth) Login(c *gin.Context) {
	params, err := filter.Login(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	// 调用service对应的方法
	logrus.Info(params)
	member := model.Member{
		Id: 1,
		Account :"gelu",
		UserName: "gelu",
		Password: "111111",
		Avatar: "https://gw.alipayobjects.com/zos/rmsportal/jZUIxmJycoymBprLOUbT.png",
		Status: 1,
		RoleId: "admin",
		Lang: "zh-CN",
		Token: "4291d7da9005377ec9aec4a71ea837f",
	}
	api.SetResponse(c, http.StatusOK, code.Success, member)
}

// @Tags 用户登录注册相关接口
// @Summary 登出 API
// @Description 登出
// @Success 200 {string} string    "ok"
// @Router /v1/auth/logout [POST]
func (ctl *Auth) Logout(c *gin.Context) {
	params, err := filter.Logout(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	// 调用service对应的方法
	logrus.Info(params)

	api.SetResponse(c, http.StatusOK, code.Success, "")
}

// @Tags 用户登录注册相关接口
// @Summary 忘记密码 API
// @Description 忘记密码
// @Success 200 {string} string    "ok"
// @Router /v1/auth/forget [GET]
func (ctl *Auth) Forget(c *gin.Context) {
	params, err := filter.Forget(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	// 调用service对应的方法
	logrus.Info(params)

	api.SetResponse(c, http.StatusOK, code.Success, "")
}

// @Tags 用户登录注册相关接口
// @Summary 注册 API
// @Description 编辑团队
// @Success 200 {string} string    "ok"
// @Router /v1/auth/register [POST]
func (ctl *Auth) Register(c *gin.Context) {
	params, err := filter.Register(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	// 调用service对应的方法
	logrus.Info(params)

	api.SetResponse(c, http.StatusOK, code.Success, "")
}
