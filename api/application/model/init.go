package model

// 团队角色
var TeamRoleMap = make(map[uint8]string, 3)

func init() {
	initTeamRoleMap()
}

func initTeamRoleMap() {
	TeamRoleMap[1] = "普通成员" // 可查看
	TeamRoleMap[2] = "开发者"  // 可上传APP
	TeamRoleMap[3] = "管理员"  // 可管理团队
}
