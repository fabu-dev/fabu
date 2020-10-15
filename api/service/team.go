package service

import (
	"fabu.dev/api/pkg/api"
)

type Team struct {

}

func NewTeam() *Team {
	return &Team{}
}

func (s *Team)CreateTeam(params *api.TeamCreateParams) error{


	return nil
}
