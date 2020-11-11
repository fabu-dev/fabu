package filter

import (
	"fabu.dev/api/controller/response"
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Team struct {
	BaseFilter
	service *service.Team
}

func NewTeam() *Team {
	return &Team{
		service: service.NewTeam(),
	}
}

// 获取一个用户的团队列表
func (f *Team) GetList(c *gin.Context) (*response.TeamList, *api.Error) {
	memberId := c.GetInt64("member_id")

	result, err := f.service.GetListByMember(uint64(memberId))

	return result, err
}

// 获取一个团队的用户信息
func (f *Team) GetMemberList(c *gin.Context) ([]*model.TeamMemberInfo, *api.Error) {
	params := &request.TeamMemberListParams{}
	if err := c.ShouldBindUri(params); err != nil {
		logrus.Error(err)

		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	teamMemberList, err := f.service.GetMemberList(params.Id)

	return teamMemberList, err
}

// 获取一个团队的用户信息
func (f *Team) DeleteMember(c *gin.Context) *api.Error {
	params := &request.TeamMemberDeleteParams{}
	if err := c.ShouldBindJSON(params); err != nil {
		logrus.Error(err)

		return api.NewError(code.ErrorRequest, err.Error())
	}

	err := f.service.DeleteMember(params.Id)
	return err
}

// 获取一个团队的用户信息
func (f *Team) AddMember(c *gin.Context) *api.Error {
	params := &request.TeamMemberAddParams{}
	if err := c.ShouldBindJSON(params); err != nil {
		logrus.Error("add member err : ", err)
		return api.NewError(code.ErrorRequest, err.Error())
	}

	operator := f.GetOperator(c)

	err := f.service.InviteMember(params, operator)
	return err
}

// 创建团队
func (f *Team) Create(c *gin.Context) (*model.TeamInfo, *api.Error) {
	params := &request.TeamCreateParams{}

	if err := c.ShouldBindJSON(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	operator := f.GetOperator(c)

	// 调用service对应的方法
	teamInfo, err := f.service.Create(params, operator)

	return teamInfo, err
}

// 解散团队
func (f *Team) Delete(c *gin.Context) *api.Error {
	params := &request.TeamDeleteParams{}

	if err := c.ShouldBindJSON(params); err != nil {
		return api.NewError(code.ErrorRequest, err.Error())
	}

	operator := f.GetOperator(c)

	// 调用service对应的方法
	err := f.service.Delete(params, operator)

	return err
}

// 编辑团队
func (f *Team) Edit(c *gin.Context) (*model.TeamInfo, *api.Error) {
	params := &request.TeamEditParams{}

	if err := c.ShouldBindJSON(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	operator := f.GetOperator(c)

	// 调用service对应的方法
	teamInfo, err := f.service.Edit(params, operator)

	return teamInfo, err
}

// 验证获取团队详情
func (f *Team) View(c *gin.Context) (*model.TeamInfo, *api.Error) {
	params := &request.TeamViewParams{}

	if err := c.ShouldBindUri(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	app, err := f.service.GetInfoById(params.Id)

	return app, err
}
