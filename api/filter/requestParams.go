package filter

type LoginParams struct {
	UserName string `form:"user_name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterParams struct {
	UserName string `form:"user_name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type ViewParams struct {
	Id uint `uri:"id" binding:"required,numeric"`
}
