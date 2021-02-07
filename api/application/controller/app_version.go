package controller

import (
	"fabu.dev/api/application/filter"
	"fabu.dev/api/application/service"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"github.com/gin-gonic/gin"
	"net/http"
	"text/template"
)

type AppVersion struct {
	paramFilter *filter.AppVersion
}

func NewAppVersion() *AppVersion {
	return &AppVersion{
		paramFilter: filter.NewAppVersion(),
	}
}

// @Tags APP版本管理
// @Summary APP版本列表 API
// @Description App版本版本列表
// @Success 200 {string} string    "ok"
// @Router /v1/app/version/ [GET]
func (ctl *AppVersion) GetList(c *gin.Context) {
	result, err := ctl.paramFilter.GetList(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ErrorRequest, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, 1, result)
}

// @Tags APP版本管理
// @Summary APP版本删除 API
// @Description APP版本删除
// @Success 200 {string} string    "ok"
// @Router /v1/app/version/delete [GET]
func (ctl *AppVersion) Delete(c *gin.Context) {
	err := ctl.paramFilter.Delete(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags iOS plist
// @Summary APP版本下载 API
// @Description iOS plist
// @Success 200  string    "ok"
// @Router /plist/:hash.plist [GET]
func (ctl *AppVersion) Plist(c *gin.Context) {
	appVersionInfo, err := ctl.paramFilter.GetInfoByHash(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}
	appInfo := service.App{}
	app, err := appInfo.GetInfoById(appVersionInfo.AppId)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	url := "https://" + c.Request.Host + "/"
	app.Icon = url + app.Icon
	appVersionInfo.Path = url + appVersionInfo.Path

	c.Header("Content-Type", "application/force-download")
	c.Header("Content-Disposition", "attachment;fileName=manifest.plist")
	//c.HTML(http.StatusOK, "template.plist", gin.H{
	//	"appVersionInfo": appVersionInfo,
	//	"appInfo":        app,
	//})

	tmpl, parseErr := template.ParseFiles("./static/template/template.plist")
	if parseErr != nil {
		api.SetResponse(c, http.StatusOK, 10002, "parseErr "+parseErr.Error())
		return
	}
	exeErr := tmpl.Execute(c.Writer, gin.H{
		"appVersionInfo": appVersionInfo,
		"appInfo":        app,
	})
	if exeErr != nil {
		api.SetResponse(c, http.StatusOK, 10002, "exeErr")
		return
	}

	return
}

// @Tags APP版本管理
// @Summary APP版本详情 API
// @Description App版本
// @Success 200 {string} string    "ok"
// @Router /v1/app/version/info/1 [GET]
func (ctl *AppVersion) View(c *gin.Context) {

}

// @Tags APP版本管理
// @Summary APP版本发布 API
// @Description APP版本发布
// @Success 200 {string} string    "ok"
// @Router /v1/app/version/publish [POST]
func (ctl *AppVersion) Publish(c *gin.Context) {

}

// @Tags APP版本管理
// @Summary APP版本取消发布 API
// @Description APP版本取消发布
// @Success 200 {string} string    "ok"
// @Router /v1/app/version/cancel [POST]
func (ctl *AppVersion) Cancel(c *gin.Context) {

}

// @Tags APP版本管理
// @Summary APP版本编辑 API
// @Description APP版本编辑
// @Success 200 {string} string    "ok"
// @Router /v1/app/version/edit [POST]
func (ctl *AppVersion) Edit(c *gin.Context) {

}

// @Tags APP版本管理
// @Summary APP版本下载记录 API
// @Description APP版本下载记录
// @Success 200 {string} string    "ok"
// @Router /v1/app/version/download/log [GET]
func (ctl *AppVersion) GetDownloadLog(c *gin.Context) {

}
