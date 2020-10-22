package request

type TeamCreateParams struct {
	Name string `json:"name" binding:"required"`
}

type TeamEditParams struct {
	Id   uint64 `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type TeamMemberListParams struct {
	Id uint64 `uri:"id" binding:"required,numeric"`
}

type TeamDeleteParams struct {
	Id uint64 `uri:"id" binding:"required,numeric"`
}

type TeamMemberDeleteParams struct {
	Id uint64 `json:"id" binding:"required"`
}

type TeamMemberAddParams struct {
	Id    uint64 `json:"id" binding:"required"`
	Name  string `json:"name"`
	Email string `json:"email" binding:"required,email"`
	Role  uint8  `json:"role" binding:"required"`
}
