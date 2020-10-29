package filter

import (
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
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

func (f *App) Upload(c *gin.Context) *api.Error {
	params := &request.UploadParams{}
	if err := c.ShouldBind(params); err != nil {
		logrus.Error("add member err : ", err)
		return api.NewError(code.ErrorRequest, err.Error())
	}

	//file, errs := c.FormFile("file")
	//if errs != nil {
	//	return api.NewError(code.ErrorRequest, fmt.Sprintf("get form err: %s", errs.Error()))
	//}

	//params.File = file

	operator := f.GetOperator(c)

	err := f.service.Upload(params, operator)

	return err
}
