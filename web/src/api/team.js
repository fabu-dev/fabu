import request from '@/utils/request'

const teamApi = {
  Create: '/team/create',
  Index: '/team/',
  Member: '/team/member/'
}

// 创建团队
export function create (parameter) {
  console.log(process.env.VUE_APP_API_BASE_URL)
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
