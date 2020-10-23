package model

type TeamStatistics struct {
	DetailColumns []string
	BaseModel
}

func NewTeamStatistics() *TeamStatistics {
	TeamStatistics := &TeamStatistics{
		DetailColumns: []string{"id"},
	}

	TeamStatistics.SetTableName("app_version")

	return TeamStatistics
}
