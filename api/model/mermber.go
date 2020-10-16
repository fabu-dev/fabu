package model

import "fabu.dev/api/pkg/api/request"

type Member struct {
	DetailColumns []string
	BaseModel
}

type MemberInfo struct {
	Id       uint64 `json:"id" gorm:"primary_key"`
	Mobile   string `json:"mobile"`
	Account  string `json:"account"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Avatar   string `json:"avatar"`
	Status uint8  `json:"status"`
	Token  string `json:"token"`
}

func NewMember() *Member {
	member := &Member{
		DetailColumns: []string{"id", "mobile", "account", "user_name", "email", "password", "status", "token"},
	}

	member.SetTableName("member")

	return member
}

func (m *Member) Add(member *MemberInfo) (*MemberInfo, error) {
	err := m.Db().Create(member).Error

	return member, err
}

func (m *Member) GetDetailByAccount(params *request.LoginParams) (*MemberInfo, error) {
	member := &MemberInfo{}
	err := m.Db().Select(m.DetailColumns).
		Where("account = ? and password = ?", params.Account, params.Password).
		First(member).Error

	return member, err
}

func (m *Member) GetDetailByToken(token string) (*MemberInfo, error) {
	member := &MemberInfo{}
	err := m.Db().Select(m.DetailColumns).
		Where("token = ?", token).
		First(member).Error

	return member, err
}
