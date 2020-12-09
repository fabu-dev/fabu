package controller

import (
	"net/http"
	"os"
	"path"

	"github.com/sirupsen/logrus"

	"fabu.dev/api/application/filter"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"github.com/astaxie/beego/logs"
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
	err := ctl.paramFilter.Delete(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags APP版本管理
// @Summary APP版本下载 API
// @Description App版本下载
// @Success 200 {string} string    "ok"
// @Router /v1/app/version/download/:id [GET]
func (ctl *AppVersion) Download(c *gin.Context) {
	filePath, err := ctl.paramFilter.Download(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, err.Code, err.Message)
		return
	}

	logrus.Info("file path :", filePath)
	//打开文件
	fileTmp, errByOpenFile := os.Open(filePath)
	defer fileTmp.Close()

	//获取文件的名称
	fileName := path.Base(filePath)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	if errByOpenFile != nil {
		logs.Error("获取文件失败")
		c.Redirect(http.StatusFound, "/404")
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")

	c.File(filePath)
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
