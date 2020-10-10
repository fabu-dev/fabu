package controller

import (
	"fabu.dev/api/service"
	"net/http"

	"fabu.dev/api/filter"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"

	"github.com/gin-gonic/gin"
)

type Team struct {
}

func NewTeam() *Team {
	return &Team{}
}


// @Tags 用户管理
// @Summary 获取单个用户信息API
// @Description 获取单个用户信息
// @Success 200 {string} string    "ok"
// @Router /user/info/1 [get]
func (m *Team) Create(c *gin.Context) {
	params, err := filter.Create(c)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	// 调用service对应的方法
	err = service.CreateTeam(params)
	if err!=nil{
		api.SetResponse(c, http.StatusOK, code.ERROR_DATABASE, err.Error())
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, nil)
}

