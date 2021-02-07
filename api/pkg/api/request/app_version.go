package request

type AppVersionIndexParams struct {
	AppId uint64 `form:"app_id" binding:"required,numeric"`
}

type AppVersionDeleteParams struct {
	Id uint64 `uri:"id" binding:"required,numeric"`
}

type AppVersionByHashParams struct {
	Hash string `uri:"hash" binding:"required"`
}
