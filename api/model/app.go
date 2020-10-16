package model

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
