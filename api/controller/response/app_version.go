package response

import "fabu.dev/api/model"

type AppVersionList struct {
	Count      uint                    `json:"count"`
	AppVersion []*model.AppVersionInfo `json:"app_version"`
}
