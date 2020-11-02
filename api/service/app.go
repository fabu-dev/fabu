package service

import (
	"encoding/json"
	"time"

	"fabu.dev/api/pkg/api/global"

	"github.com/sirupsen/logrus"

	"fabu.dev/api/pkg/constant"
	"fabu.dev/api/pkg/db"

	"fabu.dev/api/pkg/api/code"

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

// 将文件保存到channel中
func (s *App) Upload(params *request.UploadParams, operator *model.Operator) *api.Error {
	// 将数据写入到channel中
	FileChannel <- &UploadInfo{
		Params:   params,
		Operator: operator,
	}

	// 当最后一个分片到达服务器后，等待文件合并完成
	if params.ChunkNumber == params.ChunkTotal {
		for {
			if ok, _ := db.Redis.HExists(constant.AppFileInfo, params.Identifier).Result(); ok {
				break
			}
			time.Sleep(time.Millisecond * 200)
		}
	}

	return nil
}

// 获取APP信息
func (s *App) GetAppInfo(params *request.AppInfoParams, operator *model.Operator) (*global.AppInfo, *api.Error) {
	apkString, err := db.Redis.HGet(constant.AppFileInfo, params.Identifier).Bytes()
	if err != nil {
		logrus.Error("read redis err：", err)
		return nil, api.NewError(code.ErrorAppFileParserFail, code.GetMessage(code.ErrorAppFileParserFail))
	}

	apk := &global.AppInfo{}
	if err := json.Unmarshal(apkString, apk); err != nil {
		logrus.Error("json err", err)
		return nil, api.NewError(code.ErrorAppFileParserFail, code.GetMessage(code.ErrorAppFileParserFail))
	}

	return apk, nil
}
