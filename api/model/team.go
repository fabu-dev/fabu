package model

import (
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/utils"
)

type Team struct {
	DetailColumns []string
	BaseModel
}

type Operator struct {
	Id      int64  `json:"id"`
	Account string `json:"account"`
}

type TeamInfo struct {
	Id        uint64            `json:"id" gorm:"primary_key"`
	Owner     int64             `json:"owner"`
	Name      string            `json:"name"`
	Status    uint8             `json:"status"`
	CreatedBy string            `json:"created_by"`
	CreatedAt utils.JSONTime    `json:"created_at" gorm:"-"` // 插入时忽略该字段
	Member    []*TeamMemberInfo `json:"member" gorm:"-"`
}

func NewTeam() *Team {
	Team := &Team{
		DetailColumns: []string{"id", "owner", "name", "status", "created_by", "created_at"},
	}

	Team.SetTableName("team")

	return Team
}

// 根据ID获取团队列表
func (m *Team) GetListById(teamId []uint64) ([]*TeamInfo, *api.Error) {
	if len(teamId) == 0 {
		return nil, nil
	}

	teamSlice := make([]*TeamInfo, 0, len(teamId))

	err := m.Db().Select(m.DetailColumns).Where("id = (?)", teamId).Find(&teamSlice).Error
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}

	return teamSlice, nil
}

func (m *Team) Add(team *TeamInfo) *api.Error {
	err := m.Db().Create(team).Error
	if err != nil {
		return api.NewError(code.ErrorDatabase, err.Error())
	}

	return nil
}
