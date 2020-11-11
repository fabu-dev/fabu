package controller

import (
	"net/http"

	"fabu.dev/api/filter"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"github.com/gin-gonic/gin"
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

}

// @Tags APP版本管理
// @Summary APP版本下载 API
// @Description App版本下载
// @Success 200 {string} string    "ok"
// @Router /v1/app/version/download/:id [GET]
func (ctl *AppVersion) Download(c *gin.Context) {

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
