import request from '@/utils/request'

const teamApi = {
  Create: '/v1/team/create',
}


export function create (parameter) {
  console.log(teamApi.Create)
  console.log(process.env.VUE_APP_API_BASE_URL)
  return request({
    url: teamApi.Create,
    method: 'post',
    data: parameter
  })
}
