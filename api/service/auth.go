package service

import (
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/pkg/constant"
	"fabu.dev/api/pkg/utils"
	"github.com/sirupsen/logrus"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (s *Auth) Login(params *request.LoginParams) (*model.MemberInfo, *api.Error) {
	memberObj := model.NewMember()

	member, err := memberObj.GetDetailByAccount(params)
	if err != nil {
		return nil, api.NewError(code.ERROR_DATABASE, err.Error())
	}

	return member, nil
}

func (s *Auth) Register(params *request.RegisterParams) (*model.MemberInfo, *api.Error) {
	memberObj := model.NewMember()
	member := &model.MemberInfo{
		Account:  params.Account,
		Email:    params.Email,
		Password: params.Password,
		Status:   constant.StatusEnable,
		Token:    utils.GetRandom(20),
	}

	member, err := memberObj.Add(member)
	if err != nil {
		return nil, api.NewError(code.ERROR_DATABASE, err.Error())
	}

	logrus.Info("register success, member id is ", member.Id)

	return member, nil
}
