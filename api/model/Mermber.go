package model

type Member struct {
	Id uint64 `json:"id"`
	Mobile string `json:"mobile"`
	Name string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Avatar string `json:"avatar"`
	Status uint8 `json:"status"`
	Telephone string `json:"telephone"`
	RoleId string `json:"role_id"`
	Lang string `json:"lang"`
	Token string `json:"token"`
}