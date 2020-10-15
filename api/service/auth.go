package service

import (
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/pkg/constant"
	"github.com/sirupsen/logrus"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (s *Auth) Register(params *request.RegisterParams) (*model.MemberInfo, *api.Error) {
	memberObj := model.NewMember()
	member := &model.MemberInfo{
		Account:  params.Account,
		Email:    params.Email,
		Password: params.Password,
		Status:   constant.StatusEnable,
	}

	member, err := memberObj.Add(member)
	if err != nil {
		return nil, api.NewError(code.ERROR_DATABASE, err.Error())
	}

	logrus.Info("register success, member id is ", member.Id)

	return member, nil
}
