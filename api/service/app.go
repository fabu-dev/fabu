package service

import (
	"encoding/json"
	"time"

	"fabu.dev/api/controller/response"

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

// 获取会员的团队列表
func (s *App) GetListByTeamId(params *request.AppIndexParams) (*response.AppList, *api.Error) {
	// 先获取会员所有的团队
	objTeamMember := model.NewApp()
	appSlice, err := objTeamMember.GetAppListByTeamId(params.TeamId)
	if err != nil {
		return nil, err
	}

	result := &response.AppList{
		Count: 0,
		App:   appSlice,
	}

	return result, err
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

// 保存上传信息
// TODO 判断版本号，用最大的版本号作为最新版本号
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
	appInfo, err := ObjApp.GetInfoByBundleId(apk.BundleId, apk.Platform)
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

// 获取会员详细信息
func (s *App) GetInfoById(appId uint64) (*model.AppInfo, *api.Error) {
	objApp := model.NewApp()

	appInfo, err := objApp.GetInfoById(appId)

	s.ApplyPlatformName(appInfo)

	return appInfo, err
}

// 平台名
func (s *App) ApplyPlatformName(appInfo *model.AppInfo) {
	appInfo.PlatformName = constant.PlatformMap[appInfo.Platform]
}
