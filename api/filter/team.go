package filter

import (
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Team struct {
	service *service.Team
}

func NewTeam() *Team {
	return &Team{
		service: service.NewTeam(),
	}
}

// 获取一个用户的团队列表
func (f *Team) GetList(c *gin.Context) ([]*model.TeamInfo, *api.Error) {
	memberId := c.GetInt64("member_id")

	teamSlice, err := f.service.GetListByMember(uint64(memberId))

	return teamSlice, err
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

// 创建团队
func (f *Team) Create(c *gin.Context) (*model.TeamInfo, *api.Error) {
	params := &request.TeamCreateParams{}

	if err := c.ShouldBindJSON(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	operator := &model.Operator{
		Id:      c.GetInt64("member_id"),
		Account: c.GetString("account"),
	}

	// 调用service对应的方法
	teamInfo, err := f.service.Create(params, operator)

	return teamInfo, err
}

// 编辑团队
func (f *Team) Edit(c *gin.Context) (*model.TeamInfo, *api.Error) {
	params := &request.TeamEditParams{}

	if err := c.ShouldBindJSON(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	operator := &model.Operator{
		Id:      c.GetInt64("member_id"),
		Account: c.GetString("account"),
	}

	// 调用service对应的方法
	teamInfo, err := f.service.Edit(params, operator)

	return teamInfo, err
}
