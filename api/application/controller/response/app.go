package response

import "fabu.dev/api/application/model"

type AppList struct {
	Count uint             `json:"count"`
	App   []*model.AppInfo `json:"app"`
}
