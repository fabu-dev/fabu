package request

type TeamCreateParams struct {
	Name string `json:"name" binding:"required"`
}
