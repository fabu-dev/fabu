package filter

import (
	"fabu.dev/api/controller/response"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AppVersion struct {
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
