package service

import (
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/request"
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
