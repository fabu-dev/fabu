package service

import (
	"github.com/sirupsen/logrus"

	"fabu.dev/api/pkg/api/code"

	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/pkg/parser"
)

type UploadInfo struct {
	Params   *request.UploadParams
	Operator *model.Operator
}

var FileChannel = make(chan *UploadInfo, 1)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (s *App) Upload(params *request.UploadParams, operator *model.Operator) *api.Error {
	// 将数据写入到channel中
	FileChannel <- &UploadInfo{
		Params:   params,
		Operator: operator,
	}

	return nil
}

func (s *App) GetAppInfo(params *request.AppInfoParams, operator *model.Operator) (*parser.AppInfo, *api.Error) {
	apk, err := parser.NewAppParser("一米市集_3.12.0.5.ipa")
	logrus.Info(params.Filename)
	logrus.Info(apk, err)
	if err != nil {
		return nil, api.NewError(code.ErrorAppFileParserFail, code.GetMessage(code.ErrorAppFileParserFail))
	}

	return apk, nil
}
