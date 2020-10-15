package filter

import "fabu.dev/api/service"

type System struct {
	service *service.System
}

func NewSystem() *System {
	return &System{
		service: service.NewSystem(),
	}
}