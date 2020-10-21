package model

var TeamRoleMap = make(map[uint8]string, 3)

func init() {
	initTeamRole()

}

func initTeamRole() {
	TeamRoleMap[1] = "所有人"
	TeamRoleMap[2] = "管理员"
	TeamRoleMap[3] = "团队成员"
}
