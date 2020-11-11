package service

import (
	"fabu.dev/api/controller/response"
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/request"
)

type AppVersion struct {
}

func NewAppVersion() *AppVersion {
	return &AppVersion{}
}

// 获取会员的团队列表
func (s *AppVersion) GetListByAppId(params *request.AppVersionIndexParams) (*response.AppVersionList, *api.Error) {
	// 先获取会员所有的团队
	objTeamMember := model.NewAppVersion()
	appSlice, err := objTeamMember.GetListByAppId(params.AppId)
	if err != nil {
		return nil, err
	}

	result := &response.AppVersionList{
		Count:      0,
		AppVersion: appSlice,
	}

	return result, err
}
