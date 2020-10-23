package controller

import (
	"net/http"

	"fabu.dev/api/pkg/api/response"

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
// @Summary 用户团队列表 API
// @Description 用户团队列表
// @Success 200 {string} string    "ok"
// @Router /v1/team/ [GET]
func (ctl *Team) GetList(c *gin.Context) {
	teamSlice, err := ctl.paramFilter.GetList(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ErrorRequest, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, 1, result)
}

// @Tags 团队管理
// @Summary 创建团队 API
// @Description 创建团队
// @Success 200 {string} string    "ok"
// @Router /v1/team/create [POST]
func (ctl *Team) Create(c *gin.Context) {
	teamInfo, err := ctl.paramFilter.Create(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, teamInfo)
}

// @Tags 团队管理
// @Summary 编辑团队 API
// @Description 编辑团队
// @Success 200 {string} string    "ok"
// @Router /v1/team/edit [PUT]
func (ctl *Team) Edit(c *gin.Context) {
	teamInfo, err := ctl.paramFilter.Edit(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, teamInfo)
}

// @Tags 团队管理
// @Summary 添加团队成员 API
// @Description 添加团队成员
// @Success 200 {string} string    "ok"
// @Router /v1/team/member/add [POST]
func (ctl *Team) AddMember(c *gin.Context) {
	err := ctl.paramFilter.AddMember(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags 团队管理
// @Summary 移除团队成员 API
// @Description 移除团队成员
// @Success 200 {string} string    "ok"
// @Router /v1/team/member/del [DELETE]
func (ctl *Team) DeleteMember(c *gin.Context) {
	err := ctl.paramFilter.DeleteMember(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags 团队管理
// @Summary 解散团队成员 API
// @Description 移除团队成员
// @Success 200 {string} string    "ok"
// @Router /v1/team/member/del [DELETE]
func (ctl *Team) Delete(c *gin.Context) {
	err := ctl.paramFilter.Delete(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags 团队管理
// @Summary 获取团队成员信息 API
// @Description 获取团队成员信息
// @Success 200 {string} string    "ok"
// @Router /v1/team/member/1 [GET]
func (ctl *Team) GetMember(c *gin.Context) {
	memberList, err := ctl.paramFilter.GetMemberList(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	var memberRole uint8
	for _, teamMember := range memberList {
		if uint64(c.GetInt64("member_id")) == teamMember.MemberId {
			memberRole = teamMember.Role
		}
	}

	var result = &response.TeamMember{
		Role:       memberRole,
		MemberList: memberList,
	}

	api.SetResponse(c, http.StatusOK, 1, result)
}

// @Tags 团队管理
// @Summary 获取团队操作日志 API
// @Description 获取团队操作日志
// @Success 200 {string} string    "ok"
// @Router /v1/team/log/1 [GET]
func (ctl *Team) GetLog(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}
