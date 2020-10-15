package request

type ViewParams struct {
	Id uint64 `uri:"id" binding:"required,numeric"`
}
