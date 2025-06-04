import request from '../assets/lib/request'

export const fetchMenu = async () => {
  return await request({
    url: '/admin/menus',
    method: 'GET',
    data: {},
  })
}

export * from './data'
export * from './user'
export * from './env'
export * from './pages'
export * from './utils'
