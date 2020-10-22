package model

import (
	"errors"

	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/utils"
	"github.com/jinzhu/gorm"
)

type TeamMember struct {
	DetailColumns []string
	BaseModel
}

type TeamMemberInfo struct {
	Id            uint64         `json:"id" gorm:"primary_key"`
	TeamId        uint64         `json:"team_id"`
	MemberId      uint64         `json:"member_id"`
	MemberName    string         `json:"member_name" gorm:"-"`
	MemberAccount string         `json:"member_account" gorm:"-"`
	Role          uint8          `json:"role"`
	RoleName      string         `json:"role_name"  gorm:"-"`
	CreatedBy     string         `json:"created_by"`
	CreatedAt     utils.JSONTime `json:"created_at" gorm:"-"` // 插入时忽略该字段
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
	err := m.Db().Where("member_id = ?", memberId).Pluck("team_id", &teamIdSlice).Error

	return teamIdSlice, m.ProcessError(err)
}

// 获取一个团队的成员信息
func (m *TeamMember) GetListByTeamId(teamId uint64) ([]*TeamMemberInfo, *api.Error) {
	if teamId < 1 {
		return nil, api.NewError(code.ErrorRequest, "请求参数异常，ID不可以为空")
	}

	teamMemberSlice := make([]*TeamMemberInfo, 0, 16)

	err := m.Db().Select(m.DetailColumns).Where("team_id = ?", teamId).Find(&teamMemberSlice).Error

	return teamMemberSlice, m.ProcessError(err)
}

// 添加团队成员
func (m *TeamMember) Add(teamMemberInfo *TeamMemberInfo) *api.Error {
	err := m.Db().Create(teamMemberInfo).Error

	return m.ProcessError(err)
}

// 删除团队成员
func (m *TeamMember) Delete(id uint64) *api.Error {
	err := m.Db().Where("id = ?", id).Delete(&TeamMemberInfo{}).Error

	return m.ProcessError(err)
}

// 判断会员是否在团队里
func (m *TeamMember) IsInTeam(teamId, memberId uint64) (bool, *api.Error) {
	teamMemberInfo := &TeamMemberInfo{}
	err := m.Db().Select(m.DetailColumns).Where("team_id = ? and member_id = ?", teamId, memberId).First(teamMemberInfo).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, m.ProcessError(err)
}
