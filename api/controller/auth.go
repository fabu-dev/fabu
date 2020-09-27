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

func (a *Auth) Login(c *gin.Context) {
	params, err := filter.Login(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	// 调用service对应的方法
	logrus.Info(params)
	member := model.Member{
		Id: 1,
		Name :"gelu",
		UserName: "gelu",
		Password: "111111",
		Avatar: "https://gw.alipayobjects.com/zos/rmsportal/jZUIxmJycoymBprLOUbT.png",
		Status: 1,
		Telephone: "15000502790",
		RoleId: "admin",
		Lang: "zh-CN",
		Token: "4291d7da9005377ec9aec4a71ea837f",
	}
	api.SetResponse(c, http.StatusOK, code.Success, member)
}


func (a *Auth) Logout(c *gin.Context) {
	params, err := filter.Logout(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	// 调用service对应的方法
	logrus.Info(params)

	api.SetResponse(c, http.StatusOK, code.Success, "")
}

func (a *Auth) Forget(c *gin.Context) {
	params, err := filter.Forget(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	// 调用service对应的方法
	logrus.Info(params)

	api.SetResponse(c, http.StatusOK, code.Success, "")
}

func (a *Auth) Register(c *gin.Context) {
	params, err := filter.Register(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	// 调用service对应的方法
	logrus.Info(params)

	api.SetResponse(c, http.StatusOK, code.Success, "")
}
