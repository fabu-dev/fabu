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

// 获取会员的团队列表
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

	return teamSlice, err
}

// 获取单个团队的成员信息
func (s *Team) GetMemberList(teamId uint64) ([]*model.TeamMemberInfo, *api.Error) {

	objTeamMember := model.NewTeamMember()
	teamMemberList, err := objTeamMember.GetListByTeamId(teamId)

	if len(teamMemberList) == 0 {
		return teamMemberList, nil
	}

	err = s.ApplyMember(teamMemberList)

	return teamMemberList, err
}

// 给团队成员列表应用用户名
func (s *Team) ApplyMember(teamMemberList []*model.TeamMemberInfo) *api.Error {
	memberIdList := make([]uint64, len(teamMemberList))
	for _, teamMember := range teamMemberList {
		teamMember.RoleName = model.TeamRoleMap[teamMember.Role]
		memberIdList = append(memberIdList, teamMember.MemberId)
	}

	objMember := model.NewMember()
	memberList, err := objMember.GetListById(memberIdList)
	if err != nil {
		return err
	}

	for _, teamMember := range teamMemberList {
		for _, member := range memberList {
			if teamMember.MemberId == member.Id {
				teamMember.MemberName = member.UserName
				teamMember.MemberAccount = member.Account
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
		MemberId:  uint64(operator.Id),
		Role:      constant.TeamRoleCreator,
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
