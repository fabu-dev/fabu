package request

type LoginParams struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LogoutParams struct {
}

type RegisterParams struct {
	Account         string `json:"account" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

type ForgetParams struct {
}
