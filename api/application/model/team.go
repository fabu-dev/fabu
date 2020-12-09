package model

import (
	"errors"

	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/constant"
	"fabu.dev/api/pkg/utils"
	"github.com/jinzhu/gorm"
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
	Id        uint64         `json:"id" gorm:"primary_key"`
	Owner     int64          `json:"owner"`
	Name      string         `json:"name"`
	Status    uint8          `json:"status"`
	CreatedBy string         `json:"created_by"`
	CreatedAt utils.JSONTime `json:"created_at" gorm:"-"` // 插入时忽略该字段
	UpdatedBy string         `json:"updated_by"`
	UpdatedAt string         `json:"updated_at" gorm:"-"` // 插入时忽略该字段
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

	err := m.Db().Select(m.DetailColumns).Where("id in (?) and status = ?", teamId, constant.StatusEnable).Find(&teamSlice).Error

	return teamSlice, m.ProcessError(err)
}

// 创建团队
func (m *Team) Add(team *TeamInfo) *api.Error {
	err := m.Db().Create(team).Error

	return m.ProcessError(err)
}

// 编辑团队
func (m *Team) Edit(team *TeamInfo) *api.Error {
	err := m.Db().Where("id = ?", team.Id).Updates(team).Error

	return m.ProcessError(err)
}

// 根据id获取app详细信息
func (m *Team) GetInfoById(id uint64) (*TeamInfo, *api.Error) {
	appInfo := &TeamInfo{}
	err := m.Db().Select(m.DetailColumns).Where("id = ?", id).First(appInfo).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return appInfo, m.ProcessError(err)
}
