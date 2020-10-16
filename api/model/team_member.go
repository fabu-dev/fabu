package model

type TeamMember struct {
	DetailColumns []string
	BaseModel
}

func NewTeamMember() *TeamMember {
	TeamMember := &TeamMember{
		DetailColumns: []string{"id"},
	}

	TeamMember.SetTableName("team_member")

	return TeamMember
}
