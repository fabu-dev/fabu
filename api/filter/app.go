package filter

import (
	"fabu.dev/api/controller/response"
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/api/global"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/service"
	"github.com/davecgh/go-spew/spew"
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
func (f *App) GetAppInfoByIdentifier(c *gin.Context) (*global.AppInfo, *api.Error) {
	params := &request.AppInfoParams{}
	if err := c.ShouldBindJSON(params); err != nil {
		logrus.Error("add member err : ", err)
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	appInfo, err := f.service.GetAppInfoByIdentifier(params.Identifier)

	return appInfo, err
}

// 保存上传文件信息
func (f *App) Save(c *gin.Context) (*global.AppInfo, *api.Error) {
	params := &request.SaveParams{}
	if err := c.ShouldBindJSON(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	operator := f.GetOperator(c)

	appInfo, err := f.service.Save(params, operator)

	return appInfo, err
}

// 验证获取APP详情
func (f *App) View(c *gin.Context) (*model.AppInfo, *api.Error) {
	params := &request.AppViewParams{}

	if err := c.ShouldBindUri(params); err != nil {
		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	app, err := f.service.GetInfoById(params.Id)
	spew.Dump(app)
	return app, err
}
