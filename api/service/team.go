package service

import "fabu.dev/api/pkg/api/request"

type Team struct {

}

func NewTeam() *Team {
	return &Team{}
}

func (s *Team)CreateTeam(params *request.TeamCreateParams) error{


	return nil
}
