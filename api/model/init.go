package model

// 团队角色
var TeamRoleMap = make(map[uint8]string, 3)

func init() {
	initTeamRoleMap()
}

func initTeamRoleMap() {
	TeamRoleMap[1] = "团队成员"
	TeamRoleMap[2] = "管理员"
	TeamRoleMap[3] = "创建者"
}
