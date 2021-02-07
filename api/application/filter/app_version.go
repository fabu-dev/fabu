package filter

import (
	"fabu.dev/api/application/controller/response"
	"fabu.dev/api/application/model"
	"fabu.dev/api/application/service"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/api/request"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AppVersion struct {
	BaseFilter
	service *service.AppVersion
}

func NewAppVersion() *AppVersion {
	return &AppVersion{
		service: service.NewAppVersion(),
	}
}

func (f *AppVersion) GetList(c *gin.Context) (*response.AppVersionList, *api.Error) {
	params := &request.AppVersionIndexParams{}
	if err := c.ShouldBind(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	logrus.Info("getList params", params)
	result, err := f.service.GetListByAppId(params)

	return result, err
}

func (f *AppVersion) Delete(c *gin.Context) *api.Error {
	params := &request.AppVersionDeleteParams{}

	if err := c.ShouldBindJSON(params); err != nil {
		return api.NewError(code.ErrorRequest, err.Error())
	}

	operator := f.GetOperator(c)

	// 调用service对应的方法
	err := f.service.Delete(params, operator)

	return err
}

func (f *AppVersion) GetInfoByHash(c *gin.Context) (*model.AppVersionInfo, *api.Error) {
	params := &request.AppVersionByHashParams{}
	if err := c.ShouldBindUri(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	appVersion, err := f.service.GetInfoByHash(params.Hash)
	if err != nil {
		return nil, err
	}
	return appVersion, nil
}
