package controller

import (
	"go-web-frame/pkg/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type System struct {
}

func NewSystem() *System {
	return &System{}
}

// @Tags 系统管理
// @Summary 测试接口
// @Description 描述信息
// @Success 200 {string} string    "ok"
// @Router /system/health [get]
func (s *System) Health(c *gin.Context) {

	api.SetResponse(c, http.StatusOK, 1, "")
}
