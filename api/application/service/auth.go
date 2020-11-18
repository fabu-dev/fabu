package service

import (
	"time"

	"fabu.dev/api/pkg/api/code"

	"github.com/dgrijalva/jwt-go"

	"fabu.dev/api/application/model"
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

// 登录
func (s *Auth) Login(params *request.LoginParams) (*model.MemberInfo, *api.Error) {
	objMember := model.NewMember()

	member, err := objMember.GetDetailByAccount(params)
	if err != nil {
		return nil, err
	}

	ojbJwt := utils.NewJWT()
	claims := utils.CustomClaims{
		ID:     member.Id,
		Name:   member.Name,
		Mobile: member.Mobile,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "gelu",                          //签名的发行者
		},
	}

	var errJwt error
	member.AccessToken, errJwt = ojbJwt.CreateToken(claims)
	if errJwt != nil {
		return nil, api.NewError(code.ErrorSyntax, code.GetMessage(code.ErrorSyntax))
	}

	return member, err
}

// 注册
func (s *Auth) Register(params *request.RegisterParams) (*model.MemberInfo, *api.Error) {
	memberObj := model.NewMember()

	member := &model.MemberInfo{
		Account:  params.Account,
		Email:    params.Email,
		Name:     params.Name,
		Password: utils.GetMd5(params.Password),
		Status:   constant.StatusEnable,
		Token:    utils.GetRandom(20),
	}

	member, err := memberObj.Add(member)

	return member, err
}
