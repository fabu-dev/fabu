package model

type Member struct {
	Id uint64 `json:"id"`
	Mobile string `json:"mobile"`
	Account string `json:"account"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Avatar string `json:"avatar"`
	Status uint8 `json:"status"`
	RoleId string `json:"role_id"`
	Lang string `json:"lang"`
	Token string `json:"token"`
	BaseModel
}

func NewMember() *Member{

	return &Member{

	}
}




