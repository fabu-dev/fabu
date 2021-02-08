import request from '@/utils/request'

const appApi = {
  Upload: '/v1/app/upload',
  Save: '/v1/app/create',
  GetBase: '/v1/app/base',
  Index: '/v1/app/',
  Square: '/v1/app/square',
  Info: '/v1/app/info/',
  InfoByShort: '/v1/app/info?short_url=',
  Delete: '/v1/app/delete'
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

// 获取app列表 // 这里用params而不是 data
export function getList (parameter) {
  console.log('params getList', parameter)
  return request({
    url: appApi.Index,
    method: 'Get',
    timeout: 0,
    params: parameter
  })
}

export function getSquare () {
  return request({
    url: appApi.Square,
    method: 'Get',
    timeout: 0
  })
}

// 获取一个应用的详细信息
export function getInfo (id) {
  console.log('params getInfo', id)
  return request({
    url: appApi.Info + id,
    method: 'Get',
    timeout: 0
  })
}

// 获取一个应用的详细信息
export function getInfoByShort (shortUrl) {
  console.log('params getInfoByShort', shortUrl)
  return request({
    url: appApi.InfoByShort + shortUrl,
    method: 'Get',
    timeout: 0
  })
}

// 删除一个app的版本
export function deleteApp (parameter) {
  return request({
    url: appApi.Delete,
    method: 'Delete',
    timeout: 0,
    data: parameter
  })
}
