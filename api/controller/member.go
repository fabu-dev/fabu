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
// @Router /v1/user/info/1 [GET]
func (ctl *Member) Info(c *gin.Context) {
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

// @Tags 用户管理
// @Summary 获取会员的团队信息 API
// @Description 获取会员的团队信息
// @Success 200 {string} string    "ok"
// @Router /v1/user/group/1 [GET]
func (ctl *Member) GroupView(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
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
