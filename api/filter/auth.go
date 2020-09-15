package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 验证登录的请求参数
func Login(c *gin.Context) (*LoginParams, error) {
	params := &LoginParams{}

	if err := c.ShouldBindJSON(params); err != nil {
		logrus.Error(err)

		return nil, err
	}

	return params, nil
}

// 验证出的请求参数
func Logout(c *gin.Context) (*LogoutParams, error) {
	params := &LogoutParams{}

	if err := c.ShouldBind(params); err != nil {
		logrus.Error(err)

		return nil, err
	}

	return params, nil
}

// 验证注册信息
func Register(c *gin.Context) (*RegisterParams, error) {
	params := &RegisterParams{}

	if err := c.ShouldBind(params); err != nil {
		logrus.Error(err)

		return nil, err
	}

	return params, nil
}

// 验证获取用户详情
func Forget(c *gin.Context) (*ForgetParams, error) {
	params := &ForgetParams{}

	if err := c.ShouldBindUri(params); err != nil {
		logrus.Error(err)

		return nil, err
	}

	return params, nil
}
