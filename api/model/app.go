package model

import (
	"errors"

	"fabu.dev/api/pkg/api"

	"github.com/jinzhu/gorm"
)

type App struct {
	DetailColumns []string
	BaseModel
}

type AppInfo struct {
	Id uint64 `json:"id"`
}

func NewApp() *App {
	app := &App{
		DetailColumns: []string{"id"},
	}

	app.SetTableName("app")

	return app
}

// 判断团队是否有APP
func (m *App) HasByTeamId(teamId uint64) (bool, *api.Error) {
	appInfo := &AppInfo{}
	err := m.Db().Select(m.DetailColumns).Where("team_id = ?", teamId).First(appInfo).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, m.ProcessError(err)
}

func (m *App) GetAppListByTeamId(teamId uint64) ([]*AppInfo, *api.Error) {
	appList := make([]*AppInfo, 0, 8)
	err := m.Db().Select(m.DetailColumns).Where("team_id = ?", teamId).Find(&appList).Error

	return appList, m.ProcessError(err)
}
