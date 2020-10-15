package filter

import (
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	service *service.Auth
}

func NewAuth() *Auth {
	return &Auth{
		service: service.NewAuth(),
	}
}

// 验证登录的请求参数
func (f *Auth) Login(c *gin.Context) (*model.MemberInfo, error) {
	params := &request.LoginParams{}

	if err := c.ShouldBindJSON(params); err != nil {
		logrus.Error(err)

		return nil, err
	}

	// 调用service对应的方法
	logrus.Info(params)
	member := &model.MemberInfo{
		Id:       1,
		Account:  "gelu",
		UserName: "gelu",
		Password: "111111",
		Status:   1,
		Token:    "4291d7da9005377ec9aec4a71ea837f",
	}

	return member, nil
}

// 验证出的请求参数
func (f *Auth) Logout(c *gin.Context) (*request.LogoutParams, error) {
	params := &request.LogoutParams{}

	if err := c.ShouldBind(params); err != nil {
		logrus.Error(err)

		return nil, err
	}

	return params, nil
}

// 验证注册信息
func (f *Auth) Register(c *gin.Context) (*model.MemberInfo, *api.Error) {
	params := &request.RegisterParams{}

	if err := c.ShouldBind(params); err != nil {
		logrus.Error(err)

		return nil, api.NewError(code.ERROR_REQUEST, err.Error())
	}

	member, err := f.service.Register(params)

	return member, err
}

// 验证获取用户详情
func (f *Auth) Forget(c *gin.Context) (*request.ForgetParams, error) {
	params := &request.ForgetParams{}

	if err := c.ShouldBindUri(params); err != nil {
		logrus.Error(err)

		return nil, err
	}

	return params, nil
}
