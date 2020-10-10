package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//
func Create(c *gin.Context) (*TeamCreateParams, error) {
	params := &TeamCreateParams{}

	if err := c.ShouldBindJSON(params); err != nil {
		logrus.Error(err)

		return nil, err
	}

	return params, nil
}

