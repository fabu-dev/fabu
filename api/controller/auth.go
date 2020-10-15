package controller

import (
	"net/http"

	"fabu.dev/api/filter"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	paramFilter *filter.Auth
}

func NewAuth() *Auth {
	return &Auth{
		paramFilter: filter.NewAuth(),
	}
}

// @Tags 用户登录注册相关接口
// @Summary 登录 API
// @Description 登录
// @Success 200 {string} string    "ok"
// @Router /v1/auth/login [POST]
func (ctl *Auth) Login(c *gin.Context) {
	member, err := ctl.paramFilter.Login(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST, "")
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, member)
}

// @Tags 用户登录注册相关接口
// @Summary 登出 API
// @Description 登出
// @Success 200 {string} string    "ok"
// @Router /v1/auth/logout [POST]
func (ctl *Auth) Logout(c *gin.Context) {
	params, err := ctl.paramFilter.Logout(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST, "")
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
	params, err := ctl.paramFilter.Forget(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST, "")
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
	params, err := ctl.paramFilter.Register(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	// 调用service对应的方法
	logrus.Info(params)

	api.SetResponse(c, http.StatusOK, code.Success, "")
}
