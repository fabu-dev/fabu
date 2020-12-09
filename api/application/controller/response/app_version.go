package response

import "fabu.dev/api/application/model"

type AppVersionList struct {
	Count      uint                    `json:"count"`
	AppVersion []*model.AppVersionInfo `json:"app_version"`
}
