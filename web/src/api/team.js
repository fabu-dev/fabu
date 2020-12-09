import request from '@/utils/request'

const teamApi = {
  Create: '/v1/team/create',
  Edit: '/v1/team/edit',
  Index: '/v1/team/',
  Info: '/v1/team/info/',
  Del: '/v1/team/del',
  Member: '/v1/team/member/',
  MemberDel: '/v1/team/member/del',
  MemberAdd: '/v1/team/member/add'

}

// 创建团队
export function create (parameter) {
  return request({
    url: teamApi.Create,
    method: 'post',
    data: parameter
  })
}

// 获取用户的团队列表
export function index (parameter) {
  return request({
    url: teamApi.Index,
    method: 'get',
    data: parameter
  })
}

// 获取团队的成员信息
export function getMember (id) {
  return request({
    url: teamApi.Member + id,
    method: 'get',
    data: ''
  })
}

// 获取团队的成员信息
export function edit (parameter) {
  return request({
    url: teamApi.Edit,
    method: 'put',
    data: parameter
  })
}

// 获取团队的成员信息
export function exit (parameter) {
  return request({
    url: teamApi.MemberDel,
    method: 'delete',
    data: parameter
  })
}

// 解散团队
export function del (parameter) {
  return request({
    url: teamApi.Del,
    method: 'delete',
    data: parameter
  })
}

// 邀请团队成员
export function addMember (parameter) {
  return request({
    url: teamApi.MemberAdd,
    method: 'post',
    data: parameter
  })
}

// 获取一个应用的详细信息
export function getTeamInfo (id) {
  return request({
    url: teamApi.Info + id,
    method: 'Get',
    timeout: 0
  })
}
