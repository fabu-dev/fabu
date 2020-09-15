package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 验证获取用户详情
func View(c *gin.Context) (*ViewParams, error) {
	params := &ViewParams{}

	if err := c.ShouldBindUri(params); err != nil {
		logrus.Error(err)

		return nil, err
	}

	return params, nil
}
