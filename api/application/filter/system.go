package filter

import "fabu.dev/api/application/service"

type System struct {
	BaseFilter
	service *service.System
}

func NewSystem() *System {
	return &System{
		service: service.NewSystem(),
	}
}
