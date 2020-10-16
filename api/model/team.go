package model

import (
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
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
	Id        uint64 `json:"id" gorm:"primary_key"`
	Owner     int64  `json:"owner"`
	Name      string `json:"name"`
	Status    uint8  `json:"status"`
	CreatedBy string `json:"created_by"`
}

func NewTeam() *Team {
	Team := &Team{
		DetailColumns: []string{"id"},
	}

	Team.SetTableName("team")

	return Team
}

func (m *Team) Add(team *TeamInfo) (*TeamInfo, *api.Error) {
	err := m.Db().Create(team).Error
	if err != nil {
		return nil, api.NewError(code.ErrorDatabase, err.Error())
	}
	return team, nil
}
