package controller

import (
	"net/http"

	"fabu.dev/api/filter"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"

	"github.com/gin-gonic/gin"
)

type Team struct {
	paramFilter *filter.Team
}

func NewTeam() *Team {
	return &Team{
		paramFilter: filter.NewTeam(),
	}
}

// @Tags 团队管理
// @Summary 创建团队 API
// @Description 创建团队
// @Success 200 {string} string    "ok"
// @Router /v1/team/create [POST]
func (ctl *Team) Create(c *gin.Context) {
	_, err := ctl.paramFilter.Create(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ErrorRequest, "")
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, nil)
}

// @Tags 团队管理
// @Summary 编辑团队 API
// @Description 编辑团队
// @Success 200 {string} string    "ok"
// @Router /v1/team/edit [PUT]
func (ctl *Team) Edit(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags 团队管理
// @Summary 添加团队成员 API
// @Description 添加团队成员
// @Success 200 {string} string    "ok"
// @Router /v1/team/member/add [POST]
func (ctl *Team) AddMember(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags 团队管理
// @Summary 移除团队成员 API
// @Description 移除团队成员
// @Success 200 {string} string    "ok"
// @Router /v1/team/member/del [DELETE]
func (ctl *Team) DelMember(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags 团队管理
// @Summary 获取团队信息 API
// @Description 获取团队信息
// @Success 200 {string} string    "ok"
// @Router /v1/team/member/del [GET]
func (ctl *Team) View(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags 团队管理
// @Summary 获取团队操作日志 API
// @Description 获取团队操作日志
// @Success 200 {string} string    "ok"
// @Router /v1/team/log/1 [GET]
func (ctl *Team) GetLog(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}
