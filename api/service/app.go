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
	apk, err := s.GetAppInfoByIdentifier(params.Identifier)

	return apk, err
}

// 保存上传信息
func (s *App) Save(params *request.SaveParams, operator *model.Operator) (*global.AppInfo, *api.Error) {
	apk, err := s.GetAppInfoByIdentifier(params.Identifier)
	if err != nil {
		return nil, err
	}

	appInfo, err := s.SaveApp(apk, params, operator)
	if err != nil {
		return nil, err
	}

	// 判断app版本是否是第一次上传
	_, err = s.SaveAppVersion(apk, params, operator, appInfo.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// 保存app信息
func (s *App) SaveApp(apk *global.AppInfo, params *request.SaveParams, operator *model.Operator) (*model.AppInfo, *api.Error) {
	// 判断app是否是第一次上传
	ObjApp := model.NewApp()
	appInfo, err := ObjApp.GetInfoByBundleId(apk.BundleId)
	if err != nil {
		return nil, err
	}

	// 如果不是第一次上传，则修改版本号
	if appInfo != nil {
		appInfo.CurrentVersion = apk.Version
		appInfo.Icon = apk.Icon
		appInfo.CreatedBy = operator.Account

		err = ObjApp.Edit(appInfo)
		return appInfo, err
	}

	// 如果是第一次上传，则直接添加
	appInfo = &model.AppInfo{
		Name:           apk.Name,
		TeamId:         params.TeamId,
		Platform:       apk.Platform,
		Icon:           apk.Icon,
		ShortUrl:       "",
		BundleId:       apk.BundleId,
		CurrentVersion: apk.Version,
		Identifier:     "",
		Status:         constant.StatusEnable,
		CreatedBy:      operator.Account,
	}
	err = ObjApp.Add(appInfo)

	return appInfo, err
}

// 保存app信息
func (s *App) SaveAppVersion(apk *global.AppInfo, params *request.SaveParams, operator *model.Operator, appId uint64) (*model.AppVersionInfo, *api.Error) {
	// 判断app是否是第一次上传
	ObjAppVersion := model.NewAppVersion()
	appVersionInfo, err := ObjAppVersion.GetInfoByCode(appId, apk.Version)
	if err != nil {
		return nil, err
	}

	// 如果不是第一次上传，则修改版本号
	if appVersionInfo != nil {
		appVersionInfo.Description = params.Description
		appVersionInfo.Size = apk.Size
		appVersionInfo.Hash = apk.Identifier
		appVersionInfo.Path = ""
		appVersionInfo.IsPublish = constant.IsTrue
		appVersionInfo.Status = constant.StatusEnable
		appVersionInfo.UpdatedBy = operator.Account

		err = ObjAppVersion.Edit(appVersionInfo)
		return appVersionInfo, err
	}

	// 如果是第一次上传，则直接添加
	appVersionInfo = &model.AppVersionInfo{
		AppId:       appId,
		Tag:         "",
		Code:        apk.Version,
		Description: params.Description,
		Size:        apk.Size,
		Hash:        apk.Identifier,
		Path:        "",
		IsPublish:   constant.IsTrue,
		Status:      constant.StatusEnable,
		CreatedBy:   operator.Account,
	}
	err = ObjAppVersion.Add(appVersionInfo)

	return appVersionInfo, err
}

// 从redis里面获取缓存的文件信息
func (s App) GetAppInfoByIdentifier(identifier string) (*global.AppInfo, *api.Error) {
	apkString, err := db.Redis.HGet(constant.AppFileInfo, identifier).Bytes()
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
