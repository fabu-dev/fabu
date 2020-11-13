package controller

import (
	"net/http"

	"fabu.dev/api/application/filter"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"

	"github.com/gin-gonic/gin"
)

type Member struct {
	paramFilter *filter.Member
}

func NewMember() *Member {
	return &Member{
		paramFilter: filter.NewMember(),
	}
}

// @Tags 用户管理
// @Summary 获取单个用户信息API
// @Description 获取单个用户信息
// @Success 200 {string} string    "ok"
// @Router /v1/user/info/1 [GET]
func (ctl *Member) Info(c *gin.Context) {
	member, err := ctl.paramFilter.View(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, member)
}

// @Tags 用户管理
// @Summary 修改密码 API
// @Description 修改密码
// @Success 200 {string} string    "ok"
// @Router /v1/user/password [PUT]
func (ctl *Member) UpdatePassword(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags 用户管理
// @Summary 生成TOKEN API
// @Description 生成TOKEN
// @Success 200 {string} string    "ok"
// @Router /v1/user/token [POST]
func (ctl *Member) CreateToken(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}
