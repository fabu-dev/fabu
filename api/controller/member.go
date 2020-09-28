package controller

import (
	"fabu.dev/api/service"
	"net/http"

	"fabu.dev/api/filter"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"

	"github.com/gin-gonic/gin"
)

type Member struct {
}

func NewMember() *Member {
	return &Member{}
}


// @Tags 用户管理
// @Summary 获取单个用户信息API
// @Description 获取单个用户信息
// @Success 200 {string} string    "ok"
// @Router /user/info/1 [get]
func (m *Member) Info(c *gin.Context) {
	params, err := filter.View(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	// 调用service对应的方法
	member,err := service.GetMemberInfo(params.Id)
	if err!=nil{
		api.SetResponse(c, http.StatusOK, code.ERROR_DATABASE, err.Error())
		return
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
