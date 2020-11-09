package filter

import (
	"fabu.dev/api/controller/response"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/api/global"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type App struct {
	BaseFilter
	service *service.App
}

func NewApp() *App {
	return &App{
		service: service.NewApp(),
	}
}

// 获取一个用户的团队列表
func (f *App) GetList(c *gin.Context) (*response.AppList, *api.Error) {
	params := &request.TeamIndexParams{}
	if err := c.ShouldBind(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	logrus.Info("getList params", params)
	result, err := f.service.GetListByTeamId(params)

	return result, err
}

// 上传文件
func (f *App) Upload(c *gin.Context) *api.Error {
	params := &request.UploadParams{}
	if err := c.ShouldBind(params); err != nil {
		return api.NewError(code.ErrorRequest, err.Error())
	}

	operator := f.GetOperator(c)

	err := f.service.Upload(params, operator)

	return err
}

// 解析上传好的app文件
func (f *App) GetAppInfo(c *gin.Context) (*global.AppInfo, *api.Error) {
	params := &request.AppInfoParams{}
	if err := c.ShouldBindJSON(params); err != nil {
		logrus.Error("add member err : ", err)
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	operator := f.GetOperator(c)
	appInfo, err := f.service.GetAppInfo(params, operator)

	return appInfo, err
}

// 上传文件
func (f *App) Save(c *gin.Context) (*global.AppInfo, *api.Error) {
	params := &request.SaveParams{}
	if err := c.ShouldBindJSON(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	operator := f.GetOperator(c)

	appInfo, err := f.service.Save(params, operator)

	return appInfo, err
}
