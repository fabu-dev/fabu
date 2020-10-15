package filter

import (
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Team struct {
	service *service.Team
}

func NewTeam() *Team {
	return &Team{
		service: service.NewTeam(),
	}
}

//
func (f *Team) Create(c *gin.Context) (*request.TeamCreateParams, error) {
	params := &request.TeamCreateParams{}

	if err := c.ShouldBindJSON(params); err != nil {
		logrus.Error(err)

		return nil, err
	}

	// 调用service对应的方法
	err := f.service.CreateTeam(params)
	if err!=nil{
		api.SetResponse(c, http.StatusOK, code.ERROR_DATABASE, err.Error())
		return nil,nil
	}

	return params, nil
}

