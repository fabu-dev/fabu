package filter

import (
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Filter func(c *gin.Context)

type Member struct {
	service *service.Member
}

func NewMember() *Member {
	return &Member{
		service: service.NewMember(),
	}
}

// 验证获取用户详情
func (f *Member) View(c *gin.Context) (*model.Member, error) {
	params := &request.ViewParams{}

	if err := c.ShouldBindUri(params); err != nil {
		logrus.Error(err)

		return nil, err
	}

	// 调用service对应的方法
	member, err := f.service.GetMemberInfo(params.Id)
	if err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_DATABASE, err.Error())
		return nil, err
	}

	return member, nil
}
