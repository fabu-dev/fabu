package service

import (
	"encoding/json"
	"errors"
	"fabu.dev/api/pkg/config"
	"fabu.dev/api/pkg/parser"
	"fabu.dev/api/pkg/short"
	"fabu.dev/api/pkg/utils"
	"image/png"
	"io"
	"os"
	"path"
	"time"

	"fabu.dev/api/application/controller/response"

	"fabu.dev/api/pkg/api/global"

	"github.com/sirupsen/logrus"

	"fabu.dev/api/pkg/constant"
	"fabu.dev/api/pkg/db"

	"fabu.dev/api/pkg/api/code"

	"fabu.dev/api/application/model"
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
	ObjApp := model.NewApp()
	appSlice, err := ObjApp.GetAppListByTeamId(params.TeamId)
	if err != nil {
		return nil, err
	}

	for _, app := range appSlice {
		s.ApplyPlatformName(app)
	}

	result := &response.AppList{
		Count: 0,
		App:   appSlice,
	}

	return result, err
}

func (s *App) GetPublicApps() (*response.AppList, *api.Error) {
	// 先获取会员所有的团队
	ObjApp := model.NewApp()
	appSlice, err := ObjApp.GetPublicApps()
	if err != nil {
		return nil, err
	}

	for _, app := range appSlice {
		s.ApplyPlatformName(app)
	}

	result := &response.AppList{
		Count: 0,
		App:   appSlice,
	}

	return result, err
}

// 将文件保存到channel中
func (s *App) Upload(params *request.UploadParams, operator *model.Operator) *api.Error {
	// 判断文件 Hash 是否已存在
	appVersionInfo, _ := model.NewAppVersion().GetInfoByHash(params.Identifier)
	if appVersionInfo != nil {
		return api.NewError(code.ErrorAppReUpload, code.GetMessage(code.ErrorAppReUpload))
	}

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

func (s *App) UploadByAPI(params *request.UploadByAPIParams) *api.Error {
	// Token 校验
	memberModel := model.NewMember()
	member, err := memberModel.GetDetailByAPIToken(params.Token)
	if err != nil {
		return err
	}

	// 获取默认 TeamId
	var teamId uint64
	teamModel := model.NewTeamMember()
	teamIds, err := teamModel.GetTeamIdForUpload(member.Id)
	if err != nil {
		return err
	}
	for _, value := range teamIds {
		if value == params.TeamId {
			teamId = value
		}
	}
	if teamId == 0 && params.TeamId == 0 {
		teamId = teamIds[0]
	}
	if teamId == 0 {
		return api.NewError(10001, "团队信息错误")
	}

	// 环境默认为测试
	if params.Env == 0 {
		params.Env = 2
	}

	// 文件处理
	src, nErr := params.File.Open()
	if nErr != nil {
		return api.NewError(10002, nErr.Error())
	}
	defer src.Close()

	fileMD5, nErr := utils.CalcFileMD5(params.File)
	if nErr != nil {
		return api.NewError(10002, nErr.Error())
	}

	// 判断文件 Hash 是否已存在
	appVersionInfo, err := model.NewAppVersion().GetInfoByHash(fileMD5)
	if appVersionInfo != nil {
		return api.NewError(code.ErrorAppReUpload, code.GetMessage(code.ErrorAppReUpload))
	}

	// 存储文件
	filename := config.Conf.System.AppSaveRootPath + fileMD5 + path.Ext(params.File.Filename)
	out, nErr := os.Create(filename)
	if nErr != nil {
		return api.NewError(10002, nErr.Error())
	}
	defer out.Close()
	_, nErr = io.Copy(out, src)
	if nErr != nil {
		return api.NewError(10002, nErr.Error())
	}

	// 解析 APP 文件，并进行数据库处理
	saveParams := &request.SaveParams{Env: params.Env, Identifier: fileMD5, Description: params.Desc, TeamId: teamId}
	operator := &model.Operator{Id: int64(member.Id), Account: member.Account}
	nErr = s.DealApp(filename, fileMD5, saveParams, operator)
	if nErr != nil {
		return api.NewError(10002, nErr.Error())
	}

	return nil
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
	appInfo, err := ObjApp.GetInfoByBundleId(apk.BundleId, apk.Platform, params.Env)
	if err != nil {
		return nil, err
	}

	// 如果不是第一次上传，则修改版本号
	if appInfo != nil {
		appInfo.Name = apk.Name
		appInfo.CurrentVersion = apk.Version
		appInfo.CurrentBuild = apk.Build
		appInfo.Icon = apk.Icon
		appInfo.UpdatedBy = operator.Account
		appInfo.Status = constant.StatusEnable

		err = ObjApp.Edit(appInfo)
		return appInfo, err
	}

	// 如果是第一次上传，则直接添加
	appInfo = &model.AppInfo{
		Name:           apk.Name,
		TeamId:         params.TeamId,
		Platform:       apk.Platform,
		Env:            params.Env,
		Icon:           apk.Icon,
		ShortUrl:       apk.ShortKey,
		BundleId:       apk.BundleId,
		CurrentVersion: apk.Version,
		CurrentBuild:   apk.Build,
		Status:         constant.StatusEnable,
		CreatedBy:      operator.Account,
	}
	err = ObjApp.Add(appInfo)

	return appInfo, err
}

// 保存app信息
func (s *App) SaveAppVersion(apk *global.AppInfo, params *request.SaveParams, operator *model.Operator, appId uint64) (*model.AppVersionInfo, *api.Error) {

	ObjAppVersion := model.NewAppVersion()

	//// 判断app是否是第一次上传
	//appVersionInfo, err := ObjAppVersion.GetInfoByVersion(appId, apk.Version, apk.Build)
	//if err != nil {
	//	return nil, err
	//}
	//// 如果不是第一次上传，则修改版本号
	//if appVersionInfo != nil {
	//	appVersionInfo.Description = params.Description
	//	appVersionInfo.Size = apk.Size
	//	appVersionInfo.Hash = apk.Identifier
	//	appVersionInfo.Path = apk.Path
	//	appVersionInfo.IsPublish = constant.IsTrue
	//	appVersionInfo.Status = constant.StatusEnable
	//	appVersionInfo.UpdatedAt = utils.GetCurrentDateTime()
	//	appVersionInfo.UpdatedBy = operator.Account
	//
	//	err = ObjAppVersion.Edit(appVersionInfo)
	//	return appVersionInfo, err
	//}

	// 如果是第一次上传，则直接添加
	appVersionInfo := &model.AppVersionInfo{
		AppId:       appId,
		Build:       apk.Build,
		Version:     apk.Version,
		Description: params.Description,
		Size:        apk.Size,
		Hash:        apk.Identifier,
		Path:        apk.Path,
		IsPublish:   constant.IsTrue,
		Status:      constant.StatusEnable,
		CreatedBy:   operator.Account,
	}
	err := ObjAppVersion.Add(appVersionInfo)

	return appVersionInfo, err
}

// 从Redis里面获取缓存的文件信息
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

// 获取App详细信息
func (s *App) GetInfoById(appId uint64) (*model.AppInfo, *api.Error) {
	objApp := model.NewApp()

	appInfo, err := objApp.GetInfoById(appId)

	s.ApplyPlatformName(appInfo)

	return appInfo, err
}

// 获取App详细信息
func (s *App) GetInfoByShortUrl(shortUrl string) (*model.AppInfo, *api.Error) {
	objApp := model.NewApp()

	appInfo, err := objApp.GetInfoByShortUrl(shortUrl)
	if err != nil {
		return appInfo, err
	}

	s.ApplyPlatformName(appInfo)
	return appInfo, err
}

// 删除App
func (s *App) Delete(params *request.AppDeleteParams, operator *model.Operator) *api.Error {
	appInfo := &model.AppInfo{
		Id:        params.Id,
		Status:    constant.StatusDisable,
		UpdatedBy: operator.Account,
	}

	// todo 更新app的一些统计信息，版本信息

	return s.DeleteApp(appInfo)
}

// 逻辑删除App表的记录
func (s *App) DeleteApp(appInfo *model.AppInfo) *api.Error {
	objApp := model.NewApp()

	if err := objApp.Delete(appInfo); err != nil {
		return err
	}

	return nil
}

// 平台名
func (s *App) ApplyPlatformName(appInfo *model.AppInfo) {
	appInfo.PlatformName = constant.PlatformMap[appInfo.Platform]
}

// APP 文件上传后的处理
func (s *App) DealApp(filename, identifier string, params *request.SaveParams, operator *model.Operator) error {
	apk, err := parser.NewAppParser(filename)
	if err != nil {
		return err
	}

	iconName := config.Conf.System.AppSaveRootPath + identifier + ".png"
	out, err := os.Create(iconName)
	if err != nil {
		return err
	}
	defer out.Close()

	err = png.Encode(out, apk.Icon)
	if err != nil {
		return err
	}

	shortKey, err := short.NewPool().GetShortKey()
	if err != nil {
		logrus.Error("shortKey err：", err)
	}

	app := &global.AppInfo{
		Name:       apk.Name,
		BundleId:   apk.BundleId,
		Version:    apk.Version,
		Build:      apk.Build,
		Icon:       iconName,
		Size:       apk.Size,
		Identifier: identifier,
		Platform:   apk.Platform,
		ShortKey:   shortKey,
		Path:       filename,
	}

	appInfo, aErr := s.SaveApp(app, params, operator)
	if aErr != nil {
		return errors.New(aErr.Message)
	}

	_, aErr = s.SaveAppVersion(app, params, operator, appInfo.Id)
	if aErr != nil {
		return errors.New(aErr.Message)
	}

	return nil
}
