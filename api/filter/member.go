package filter

import (
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/service"
	"github.com/gin-gonic/gin"
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
func (f *Member) View(c *gin.Context) (*model.MemberInfo, *api.Error) {
	params := &request.ViewParams{}

	if err := c.ShouldBindUri(params); err != nil {

		return nil, api.NewError(code.ErrorRequest, err.Error())
	}

	// 调用service对应的方法
	member, err := f.service.GetMemberInfo(params.Id)

	return member, err
}
