import request from '@/utils/request'

const appVersionApi = {
  Index: 'app/version/'
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
