import request from '../assets/lib/request'
import { ACCESS_TOKEN, getCache, setCache } from '../assets/lib/cache'
import { pages } from '../app'

export const getToken = () => {
  const token = getCache(ACCESS_TOKEN)
  return token
}

export const setToken = (token) => {
  setCache(ACCESS_TOKEN, token)
}

export const logout = () => {
  setCache(ACCESS_TOKEN, null)
  window.location = pages.login
}

export const fetchUserInfo = async () => {
  return await request({
    url: '/admin/sys/user/info',
    data: {},
  })
}
