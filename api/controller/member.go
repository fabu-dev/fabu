package controller

import (
	"fabu.dev/api/model"
	"net/http"

	"fabu.dev/api/filter"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"

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

// 用户团队API
func (m *Member) GroupView(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

func (m *Member) UpdatePassword(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}
