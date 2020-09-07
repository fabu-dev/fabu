package controller

import (
	"go-web-frame/filter"
	"go-web-frame/pkg/api"
	"go-web-frame/pkg/api/code"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type Member struct {
}

func NewMember() *Member {
	return &Member{}
}

// 用户资料API
func (m *Member) View(c *gin.Context) {
	params, err := filter.View(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	// 调用service对应的方法
	logrus.Info(params)

	api.SetResponse(c, http.StatusOK, code.Success, "")
}

// 用户登录API
func (m *Member) Login(c *gin.Context) {
	params, err := filter.Login(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	// 调用service对应的方法
	logrus.Info(params)

	api.SetResponse(c, http.StatusOK, code.Success, "")
}

// 用户注册API
func (m *Member) Register(c *gin.Context) {
	params, err := filter.Register(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	// 调用service对应的方法
	logrus.Info(params)

	api.SetResponse(c, http.StatusOK, code.Success, "")
}

// 用户团队API
func (m *Member) GroupView(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

func (m *Member) UpdatePassword(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}
