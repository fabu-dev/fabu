package service

import (
	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/request"
	"fabu.dev/api/pkg/constant"
	"fabu.dev/api/pkg/utils"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (s *Auth) Login(params *request.LoginParams) (*model.MemberInfo, *api.Error) {
	objMember := model.NewMember()

	member, err := objMember.GetDetailByAccount(params)

	return member, err
}

func (s *Auth) Register(params *request.RegisterParams) (*model.MemberInfo, *api.Error) {
	memberObj := model.NewMember()

	member := &model.MemberInfo{
		Account:  params.Account,
		Email:    params.Email,
		Password: utils.GetMd5(params.Password),
		Status:   constant.StatusEnable,
		Token:    utils.GetRandom(20),
	}

	member, err := memberObj.Add(member)

	return member, err
}
