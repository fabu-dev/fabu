package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 验证登录的请求参数
func Login(c *gin.Context) (*LoginParams, error) {
	params := &LoginParams{}

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
func View(c *gin.Context) (*ViewParams, error) {
	params := &ViewParams{}

	if err := c.ShouldBindUri(params); err != nil {
		logrus.Error(err)

		return nil, err
	}

	return params, nil
}
