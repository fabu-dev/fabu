import request from '@/utils/request'

const appApi = {
  Upload: '/app/upload',
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
