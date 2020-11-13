package service

import (
	"fabu.dev/api/application/controller/response"
	"fabu.dev/api/application/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/pkg/constant"
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

// 删除App版本
func (s *AppVersion) Delete(params *request.AppVersionDeleteParams, operator *model.Operator) *api.Error {
	teamInfo := &model.AppVersionInfo{
		Id:        params.Id,
		Status:    constant.StatusDisable,
		UpdatedBy: operator.Account,
	}

	// todo 更新app的一些统计信息，版本信息

	return s.DeleteAppVersion(teamInfo)
}

// 逻辑删除team表的记录
func (s *AppVersion) DeleteAppVersion(appVersionInfo *model.AppVersionInfo) *api.Error {
	objAppVersion := model.NewAppVersion()

	if err := objAppVersion.Edit(appVersionInfo); err != nil {
		return err
	}

	return nil
}
