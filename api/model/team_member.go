package model

import (
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/utils"
)

type TeamMember struct {
	DetailColumns []string
	BaseModel
}

type TeamMemberInfo struct {
	Id        uint64         `json:"id" gorm:"primary_key"`
	TeamId    uint64         `json:"team_id"`
	MemberId  int64          `json:"member_id"`
	Role      uint8          `json:"role"`
	CreatedBy string         `json:"created_by"`
	CreatedAt utils.JSONTime `json:"created_at" gorm:"-"` // 插入时忽略该字段
}

func NewTeamMember() *TeamMember {
	TeamMember := &TeamMember{
		DetailColumns: []string{"id", "team_id", "member_id", "role", "created_by", "created_at"},
	}

	TeamMember.SetTableName("team_member")

	return TeamMember
}

// 获取会员的团队ID
func (m *TeamMember) GetTeamId(memberId uint64) ([]uint64, *api.Error) {
	teamIdSlice := make([]uint64, 0, 8)
	if err := m.Db().Where("member_id = ?", memberId).Pluck("team_id", &teamIdSlice).Error; err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	return teamIdSlice, nil
}

func (m *TeamMember) GetListByTeamId(teamId []uint64) ([]*TeamMemberInfo, *api.Error) {
	if len(teamId) == 0 {
		return nil, nil
	}

	teamMemberSlice := make([]*TeamMemberInfo, 0, len(teamId))

	err := m.Db().Select(m.DetailColumns).Where("team_id = (?)", teamId).Find(&teamMemberSlice).Error
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	return teamMemberSlice, nil
}

func (m *TeamMember) Add(teamMemberInfo *TeamMemberInfo) *api.Error {
	err := m.Db().Create(teamMemberInfo).Error
	if err != nil {
		return api.NewError(code.ErrorDatabase, err.Error())
	}

	return nil
}
