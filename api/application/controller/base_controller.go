package controller

import (
	"fabu.dev/api/application/filter"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Success()
	Failure()
}

// 控制器基类，抽象出一些通用的方法放在基类
type BaseController struct {
	C *gin.Context
}

func (ctl *BaseController) Success() {

	ctl.Response()
}

func (ctl *BaseController) Failure() {

	ctl.Response()
}

func (ctl *BaseController) Response() {

}

func (ctl *BaseController) Filter(c *gin.Context, filterHandel filter.Filter) {
	filterHandel(c)
}

func (ctl *BaseController) SetGinContext(c *gin.Context) {
	ctl.C = c
}
