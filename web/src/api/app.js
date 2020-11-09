import request from '@/utils/request'

const appApi = {
  Upload: '/app/upload',
  Save: '/app/create',
  GetBase: '/app/base',
  Index: '/app/'
}

// 上传APP
export function upload (parameter) {
  console.log('params upload', parameter)
  return request({
    url: appApi.Upload,
    method: 'POST',
    timeout: 0,
    data: parameter
  })
}

// 保存APP信息
export function save (parameter) {
  console.log('params upload', parameter)
  return request({
    url: appApi.Save,
    method: 'POST',
    timeout: 0,
    data: parameter
  })
}

// 获取上次的app信息
export function getBase (parameter) {
  console.log('params GetBase', parameter)
  return request({
    url: appApi.GetBase,
    method: 'POST',
    timeout: 0,
    data: parameter
  })
}

// 保存APP信息 // 这里用params而不是 data
export function getList (parameter) {
  console.log('params getList', parameter)
  return request({
    url: appApi.Index,
    method: 'Get',
    timeout: 0,
    params: parameter
  })
}
