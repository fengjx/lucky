import axios from 'axios'

import { getToken, setToken, getEnv } from '../../../app'
import consts from '../../../consts'

const service = axios.create({
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000, // request timeout
  headers: {
    'Content-Type': 'application/json',
  },
})

// request interceptor
service.interceptors.request.use(
  (config) => {
    const env = getEnv()
    config.baseURL = env.BaseAPI
    config.headers[consts.ADMIN_TOKEN] = getToken() || ''
    return config
  },
  (e) => {
    console.log(e) // for debug
    return Promise.reject(e)
  }
)

// response interceptor
service.interceptors.response.use(
  (response) => {
    const headers = response.headers
    if (headers && headers[consts.ADMIN_REFRESH_TOKEN]) {
      const token = headers[consts.ADMIN_REFRESH_TOKEN]
      console.log('refresh token', token)
      setToken(token)
    }
    return response.data
  },
  (e) => {
    if (e.response?.status == 401 && window.location.pathname != '/login') {
      setToken('')
      window.location = '/login'
      return
    }
    console.log('net err', e) // for debug
    return Promise.reject(e)
  }
)

export default service
