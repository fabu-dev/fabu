package controller

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"

	"fabu.dev/api/pkg/api/code"

	"fabu.dev/api/application/filter"
	"fabu.dev/api/pkg/api"
	"github.com/gin-gonic/gin"
)

type App struct {
	paramFilter *filter.App
}

func NewApp() *App {
	return &App{
		paramFilter: filter.NewApp(),
	}
}

// @Tags APP管理
// @Summary APP列表 API
// @Description APP列表
// @Success 200 {string} string    "ok"
// @Router /v1/app/ [GET]
func (ctl *App) GetList(c *gin.Context) {
	result, err := ctl.paramFilter.GetList(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ErrorRequest, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, 1, result)
}

// @Tags APP管理
// @Summary App上传 API
// @Description App上传
// @Success 200 {string} string    "ok"
// @Router /v1/app/upload [POST]
func (ctl *App) Upload(c *gin.Context) {
	err := ctl.paramFilter.Upload(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags APP管理
// @Summary App上传 API
// @Description App上传
// @Success 200 {string} string    "ok"
// @Router /v1/app/upload [POST]
func (ctl *App) Save(c *gin.Context) {
	appInfo, err := ctl.paramFilter.Save(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, 1, appInfo)
}

// @Tags APP管理
// @Summary App详情 API
// @Description App详情
// @Success 200 {string} string    "ok"
// @Router /v1/app/info [GET]
func (ctl *App) GetAppInfoByIdentifier(c *gin.Context) {
	appInfo, err := ctl.paramFilter.GetAppInfoByIdentifier(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, 1, appInfo)
}

// @Tags APP管理
// @Summary App详情 API
// @Description App详情
// @Success 200 {string} string    "ok"
// @Router /v1/app/info/1 [GET]
func (ctl *App) View(c *gin.Context) {
	teamInfo, err := ctl.paramFilter.View(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	spew.Dump(teamInfo)

	api.SetResponse(c, http.StatusOK, code.Success, teamInfo)
}

// @Tags APP管理
// @Summary App编辑 API
// @Description App编辑
// @Success 200 {string} string    "ok"
// @Router /v1/app/edit [PUT]
func (ctl *App) Edit(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags APP管理
// @Summary App删除 API
// @Description App删除
// @Success 200 {string} string    "ok"
// @Router /v1/app/delete [DELETE]
func (ctl *App) Delete(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags APP管理
// @Summary App合并 API
// @Description App合并
// @Success 200 {string} string    "ok"
// @Router /v1/app/combine [POST]
func (ctl *App) Combine(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags APP管理
// @Summary App统计 API
// @Description App统计
// @Success 200 {string} string    "ok"
// @Router /v1/app/stat [GET]
func (ctl *App) GetStat(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags APP管理
// @Summary 获取App操作日志 API
// @Description 获取App操作日志
// @Success 200 {string} string    "ok"
// @Router /v1/app/log [GET]
func (ctl *App) GetLog(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}
