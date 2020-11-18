import request from '@/utils/request'

const appVersionApi = {
  Index: '/v1/app/version/',
  Delete: '/v1/app/version/delete'
}

// 获取一个app的版本列表
export function getList (parameter) {
  console.log('params getList', parameter)
  return request({
    url: appVersionApi.Index,
    method: 'Get',
    timeout: 0,
    params: parameter
  })
}

// 删除一个app的版本
export function deleteVersion (parameter) {
  console.log('params getList', parameter)
  return request({
    url: appVersionApi.Delete,
    method: 'Delete',
    timeout: 0,
    data: parameter
  })
}
