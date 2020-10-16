package service

import (
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/pkg/constant"
)

type Team struct {
}

func NewTeam() *Team {
	return &Team{}
}

func (s *Team) GetListByMember(memberId uint64) ([]*model.TeamInfo, *api.Error) {
	// 先获取会员所有的团队
	objTeamMember := model.NewTeamMember()
	teamIdSlice, err := objTeamMember.GetTeamId(memberId)
	if err != nil {
		return nil, err
	}

	// 获取团队信息
	objTeam := model.NewTeam()
	teamSlice, err := objTeam.GetListById(teamIdSlice)
	if err != nil {
		return nil, err
	}

	err = s.ApplyMember(teamIdSlice, &teamSlice)

	return teamSlice, err
}

// 应用团队的成员信息
func (s *Team) ApplyMember(teamIdSlice []uint64, teamSlice *[]*model.TeamInfo) *api.Error {
	objTeamMember := model.NewTeamMember()
	teamMemberSlice, err := objTeamMember.GetListByTeamId(teamIdSlice)
	if err != nil {
		return err
	}

	for _, teamMember := range teamMemberSlice {
		for _, team := range *teamSlice {
			if team.Id == teamMember.TeamId {
				team.Member = append(team.Member, teamMember)
			}
		}
	}

	return nil
}

// 创建团队
func (s *Team) Create(params *request.TeamCreateParams, operator *model.Operator) (*model.TeamInfo, *api.Error) {
	teamInfo := &model.TeamInfo{
		Owner:     operator.Id,
		Name:      params.Name,
		Status:    constant.StatusEnable,
		CreatedBy: operator.Account,
	}
	if err := s.CreateTeam(teamInfo); err != nil {
		return nil, err
	}

	teamMemberInfo := &model.TeamMemberInfo{
		TeamId:    teamInfo.Id,
		MemberId:  operator.Id,
		Role:      1,
		CreatedBy: operator.Account,
	}
	err := s.AddMember(teamMemberInfo)

	return teamInfo, err
}

// 保存数据到team表
func (s *Team) CreateTeam(teamInfo *model.TeamInfo) *api.Error {
	objTeam := model.NewTeam()

	return objTeam.Add(teamInfo)
}

// 保存数据到team_member表（将创建团队的人，加入到这个团队）
func (s *Team) AddMember(teamMemberInfo *model.TeamMemberInfo) *api.Error {
	objTeamMember := model.NewTeamMember()

	return objTeamMember.Add(teamMemberInfo)
}
