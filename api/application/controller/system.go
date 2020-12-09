package controller

import (
	"net/http"

	"fabu.dev/api/application/filter"

	"fabu.dev/api/pkg/api"

	"github.com/gin-gonic/gin"
)

type System struct {
	paramFilter *filter.System
}

func NewSystem() *System {
	return &System{
		paramFilter: filter.NewSystem(),
	}
}

// @Tags 系统管理
// @Summary 健康检查API
// @Description 描述信息
// @Success 200 {string} string    "ok"
// @Router /system/health [get]
func (ctl *System) Health(c *gin.Context) {

	api.SetResponse(c, http.StatusOK, 1, "")
}

// @Tags 系统管理
// @Summary 系统配置 API
// @Description 系统配置
// @Success 200 {string} string    "ok"
// @Router /system/setting [POST]
func (ctl *System) Setting(c *gin.Context) {

	api.SetResponse(c, http.StatusOK, 1, "")
}
