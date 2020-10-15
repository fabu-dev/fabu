package model

type Member struct {
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
	return &Member{}
}

func (m *Member) GetTableName() string {
	return "member"
}

func (m *Member) Add(member *MemberInfo) (*MemberInfo, error) {
	err := m.Find().Table(m.GetTableName()).Create(member).Error

	if err != nil {
		return nil, err
	}

	return member, nil
}
