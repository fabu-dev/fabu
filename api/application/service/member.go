package service

import (
	"fabu.dev/api/application/model"
	"fabu.dev/api/pkg/api"
)

type Member struct {
}

func NewMember() *Member {
	return &Member{}
}

// 获取会员详细信息
func (s *Member) GetMemberInfo(memberId uint64) (*model.MemberInfo, *api.Error) {
	objMember := model.NewMember()

	memberInfo, err := objMember.GetDetailByID(memberId)
	return memberInfo, err
}
