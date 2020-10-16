package model

type AppVersion struct {
	DetailColumns []string
	BaseModel
}

func NewAppVersion() *AppVersion {
	AppVersion := &AppVersion{
		DetailColumns: []string{"id"},
	}

	AppVersion.SetTableName("app_version")

	return AppVersion
}
