package model

type AppDownloadLog struct {
	DetailColumns []string
	BaseModel
}

func NewAppDownloadLog() *AppDownloadLog {
	AppDownloadLog := &AppDownloadLog{
		DetailColumns: []string{"id"},
	}

	AppDownloadLog.SetTableName("AppDownloadLog")

	return AppDownloadLog
}
