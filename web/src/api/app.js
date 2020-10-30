import request from '@/utils/request'

const appApi = {
  Upload: '/app/upload',
  GetBase: '/app/base',
  Index: ''
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

// 上传APP
export function getBase (parameter) {
  console.log('params GetBase', parameter)
  return request({
    url: appApi.GetBase,
    method: 'POST',
    timeout: 0,
    data: parameter
  })
}
