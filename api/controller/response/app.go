package response

import "fabu.dev/api/model"

type AppList struct {
	Count uint             `json:"count"`
	App   []*model.AppInfo `json:"app"`
}
