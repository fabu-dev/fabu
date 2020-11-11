package request

type AppVersionIndexParams struct {
	AppId uint64 `form:"app_id" binding:"required,numeric"`
}
