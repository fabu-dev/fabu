package response

import "fabu.dev/api/application/model"

type TeamList struct {
	CountTeam        int               `json:"count_team"`
	CountApp         uint              `json:"count_app"`
	CountAppDownload uint              `json:"count_app_download"`
	Team             []*model.TeamInfo `json:"team"`
}

type TeamMember struct {
	Role       uint8                   `json:"role"`
	MemberList []*model.TeamMemberInfo `json:"member_list"`
}
