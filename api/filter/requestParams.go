package filter

type LoginParams struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LogoutParams struct {
}

type RegisterParams struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ForgetParams struct {
}

type ViewParams struct {
	Id uint `uri:"id" binding:"required,numeric"`
}
