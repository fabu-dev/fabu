package filter

import "fabu.dev/api/service"

type AppVersion struct {
	service *service.AppVersion
}

func NewAppVersion() *AppVersion {
	return &AppVersion{
		service: service.NewAppVersion(),
	}
}