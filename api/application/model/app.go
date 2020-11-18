package model

import (
	"errors"

	"fabu.dev/api/pkg/constant"

	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/utils"

	"github.com/jinzhu/gorm"
)

type App struct {
	DetailColumns []string
	BaseModel
}

type AppInfo struct {
	Id             uint64         `json:"id" gorm:"primary_key"`
	Name           string         `json:"name"`
	TeamId         uint64         `json:"team_id"`
	Platform       uint8          `json:"platform"`
	PlatformName   string         `json:"platform_name" gorm:"-"`
	Icon           string         `json:"icon"`
	ShortUrl       string         `json:"short_url"`
	BundleId       string         `json:"bundle_id"`
	CurrentVersion string         `json:"current_version"`
	Identifier     string         `json:"identifier"`
	Status         uint8          `json:"status"`
	CreatedBy      string         `json:"created_by"`
	CreatedAt      utils.JSONTime `json:"created_at" gorm:"-"` // 插入时忽略该字段
	UpdatedBy      string         `json:"updated_by"`
	UpdatedAt      string         `json:"updated_at" gorm:"-"` // 插入时忽略该字段
}

func NewApp() *App {
	app := &App{
		DetailColumns: []string{"id", "name", "team_id", "platform", "icon", "short_url", "bundle_id", "current_version", "identifier", "status", "updated_at", "updated_by", "created_at", "created_by"},
	}

	app.SetTableName("app")

	return app
}

// 判断团队是否有APP
func (m *App) HasByTeamId(teamId uint64) (bool, *api.Error) {
	appInfo := &AppInfo{}
	err := m.Db().Select(m.DetailColumns).Where("team_id = ? and status = ?", teamId, constant.StatusEnable).First(appInfo).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, m.ProcessError(err)
}

// 查询团队的app列表
func (m *App) GetAppListByTeamId(teamId uint64) ([]*AppInfo, *api.Error) {
	appList := make([]*AppInfo, 0, 8)
	err := m.Db().Select(m.DetailColumns).Where("team_id = ? and status = ?", teamId, constant.StatusEnable).Find(&appList).Error

	return appList, m.ProcessError(err)
}

// 添加app信息
func (m *App) Add(app *AppInfo) *api.Error {
	err := m.Db().Create(app).Error

	return m.ProcessError(err)
}

// 编辑app信息
func (m *App) Edit(app *AppInfo) *api.Error {
	err := m.Db().Where("id = ?", app.Id).Updates(app).Error

	return m.ProcessError(err)
}

// 根据包名和平台获取app详细信息
func (m *App) GetInfoByBundleId(bundleId string, platform uint8) (*AppInfo, *api.Error) {
	appInfo := &AppInfo{}
	err := m.Db().Select(m.DetailColumns).Where("bundle_id = ? and platform = ? and status = ?", bundleId, platform, constant.StatusEnable).First(appInfo).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return appInfo, m.ProcessError(err)
}

// 根据id获取app详细信息
func (m *App) GetInfoById(id uint64) (*AppInfo, *api.Error) {
	appInfo := &AppInfo{}
	err := m.Db().Select(m.DetailColumns).Where("id = ? and status = ?", id, constant.StatusEnable).First(appInfo).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return appInfo, m.ProcessError(err)
}

// 编辑app信息
func (m *App) Delete(app *AppInfo) *api.Error {
	err := m.Db().Where("id = ?", app.Id).Delete(app).Error

	return m.ProcessError(err)
}
