package request

type TeamCreateParams struct {
	Name string `json:"name" binding:"required"`
}

type TeamMemberListParams struct {
	Id uint64 `uri:"id" binding:"required,numeric"`
}
