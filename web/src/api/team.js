import request from '@/utils/request'

const teamApi = {
  Create: '/team/create',
  Index: '/team/'
}

// 创建团队
export function create (parameter) {
  console.log(teamApi.Create)
  console.log(process.env.VUE_APP_API_BASE_URL)
  return request({
    url: teamApi.Create,
    method: 'post',
    data: parameter
  })
}

// 获取用户的团队列表
export function index (parameter) {
  console.log(teamApi.Index)
  return request({
    url: teamApi.Index,
    method: 'get',
    data: parameter
  })
}
