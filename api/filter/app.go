package filter

import "fabu.dev/api/service"

type App struct {
	service *service.App
}

func NewApp() *App {
	return &App{
		service: service.NewApp(),
	}
}
