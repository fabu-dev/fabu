package service

import (
	"fabu.dev/api/model"
)

func GetMemberInfo(memberId uint64) (*model.Member,error){
	member := model.NewMember()

	err := member.Find().Where("id = ?", memberId).First(member).Error
	if err != nil {
		return nil,err
	}

	return member, nil
}