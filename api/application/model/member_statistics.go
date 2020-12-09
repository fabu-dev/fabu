package model

type MemberStatistics struct {
	DetailColumns []string
	BaseModel
}

func NewMemberStatistics() *MemberStatistics {
	MemberStatistics := &MemberStatistics{
		DetailColumns: []string{"id"},
	}

	MemberStatistics.SetTableName("app_version")

	return MemberStatistics
}
