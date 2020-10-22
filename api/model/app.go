package model

import (
	"errors"

	"fabu.dev/api/pkg/api/code"

	"fabu.dev/api/pkg/api"

	"github.com/jinzhu/gorm"
)

type App struct {
	DetailColumns []string
	BaseModel
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
	team := &struct {
		Id uint64
	}{}
	err := m.Db().Select(m.DetailColumns).Where("team_id = ?", teamId).First(team).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return true, api.NewError(code.ErrorDatabase, err.Error())
	}

	return true, nil
}
