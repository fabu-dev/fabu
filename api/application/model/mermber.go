package model

import (
	"errors"

	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"fabu.dev/api/pkg/api/request"
	"github.com/jinzhu/gorm"
)

type Member struct {
	DetailColumns []string
	BaseModel
}

type MemberInfo struct {
	Id          uint64 `json:"id" gorm:"primary_key"`
	Mobile      string `json:"mobile"`
	Account     string `json:"account"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar" gorm:"default:''"` // 指定默认值
	Status      uint8  `json:"status"`
	AccessToken string `json:"access_token" gorm:"-"`
	Token       string `json:"token"`
}

func NewMember() *Member {
	member := &Member{
		DetailColumns: []string{"id", "mobile", "account", "name", "email", "password", "status", "token"},
	}

	member.SetTableName("member")

	return member
}

// 添加会员
func (m *Member) Add(member *MemberInfo) (*MemberInfo, *api.Error) {
	err := m.Db().Create(member).Error

	return member, m.ProcessError(err)
}

// 根据ID获取一组会员
func (m *Member) GetListById(memberIdList []uint64) ([]*MemberInfo, *api.Error) {
	if len(memberIdList) == 0 {
		return nil, nil
	}

	memberList := make([]*MemberInfo, len(memberIdList))
	err := m.Db().Select(m.DetailColumns).Where("id in (?)", memberIdList).Find(&memberList).Error

	return memberList, m.ProcessError(err)
}

// 根据登录账号获取会员信息
func (m *Member) GetDetailByAccount(params *request.LoginParams) (*MemberInfo, *api.Error) {
	member := &MemberInfo{}
	err := m.Db().Select(m.DetailColumns).Where("account = ? and password = ?", params.Account, params.Password).First(member).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, api.NewError(code.ErrorMemberNotExists, code.GetMessage(code.ErrorMemberNotExists))
	}

	return member, m.ProcessError(err)
}

// 根据token获取会员账号信息
func (m *Member) GetDetailByToken(token string) (*MemberInfo, *api.Error) {
	member := &MemberInfo{}
	err := m.Db().Select(m.DetailColumns).Where("token = ?", token).First(member).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, api.NewError(code.ErrorMemberNotExists, code.GetMessage(code.ErrorMemberNotExists))
	}

	return member, m.ProcessError(err)
}

// 根据email获取会员账号信息
func (m *Member) GetDetailByEmail(email string) (*MemberInfo, *api.Error) {
	member := &MemberInfo{}
	err := m.Db().Select(m.DetailColumns).Where("email = ?", email).First(member).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, api.NewError(code.ErrorMemberNotExists, code.GetMessage(code.ErrorMemberNotExists))
	}

	return member, m.ProcessError(err)
}

// 根据ID获取会员账号信息
func (m *Member) GetDetailByID(id uint64) (*MemberInfo, *api.Error) {
	member := &MemberInfo{}
	err := m.Db().Select(m.DetailColumns).Where("id = ?", id).First(member).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, api.NewError(code.ErrorMemberNotExists, code.GetMessage(code.ErrorMemberNotExists))
	}

	return member, m.ProcessError(err)
}
