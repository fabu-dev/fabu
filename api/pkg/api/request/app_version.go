package request

type AppVersionIndexParams struct {
	AppId uint64 `form:"app_id" binding:"required,numeric"`
}

type AppVersionDeleteParams struct {
	Id uint64 `uri:"id" binding:"required,numeric"`
}

type DownloadParams struct {
	Hash string `uri:"hash" binding:"required"`
}
