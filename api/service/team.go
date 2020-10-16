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

func (s *Team) CreateTeam(params *request.TeamCreateParams, operator *model.Operator) (*model.TeamInfo, *api.Error) {

	obj := model.NewTeam()
	teamInfo := &model.TeamInfo{
		Owner:     operator.Id,
		Name:      params.Name,
		Status:    constant.StatusEnable,
		CreatedBy: operator.Account,
	}

	teamInfo, err := obj.Add(teamInfo)

	return teamInfo, err
}
