import request from '@/utils/request'

const userApi = {
  Login: '/v1/auth/login',
  Logout: '/v1/auth/logout',
  ForgePassword: '/v1/auth/forge-password',
  Register: '/v1/auth/register',
  twoStepCode: '/v1/auth/2step-code',
  SendSms: '/v1/account/sms',
  SendSmsErr: '/v1/account/sms_err',
  UserInfo: '/v1/user/info',
  UserMenu: '/v1/user/nav'
}

/**
 * login func
 * parameter: {
 *     username: '',
 *     password: '',
 *     remember_me: true,
 *     captcha: '12345'
 * }
 * @param parameter
 * @returns {*}
 */
export function login (parameter) {
  console.log(userApi.Login)
  console.log(process.env.VUE_APP_API_BASE_URL)
  return request({
    url: userApi.Login,
    method: 'post',
    data: parameter
  })
}

/**
 * register func
 * parameter: {
 *     username: '',
 *     password: '',
 *     remember_me: true,
 *     captcha: '12345'
 * }
 * @param parameter
 * @returns {*}
 */
export function register (parameter) {
  console.log(userApi.Register)
  return request({
    url: userApi.Register,
    method: 'post',
    data: parameter
  })
}

export function getSmsCaptcha (parameter) {
  return request({
    url: userApi.SendSms,
    method: 'post',
    data: parameter
  })
}

export function getInfo (id) {
  return request({
    url: userApi.UserInfo + '/' + id,
    method: 'get',
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}

export function getCurrentUserNav () {
  return request({
    url: userApi.UserMenu,
    method: 'get'
  })
}

export function logout () {
  return request({
    url: userApi.Logout,
    method: 'post',
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}

/**
 * get user 2step code open?
 * @param parameter {*}
 */
export function get2step (parameter) {
  return request({
    url: userApi.twoStepCode,
    method: 'post',
    data: parameter
  })
}
