export * from './keys'

const prefix = 'lucky.'

// 缓存时间包装
interface DataWrap {
  data: any
  expireAt: number
}

const buildKey = (key: string): any => {
  return prefix + '_' + key
}

const _get = (key: string): [any, boolean] => {
  const item = localStorage.getItem(buildKey(key))
  if (!item) {
    return [null, false]
  }
  const dw: DataWrap = JSON.parse(item)
  const now = Date.now()
  if (now >= dw.expireAt) {
    return [null, false]
  }
  return [dw.data, true]
}

export const getCache = (key: string, createFun: Function) => {
  const [res, ok] = _get(key)
  if (ok) {
    return res
  }
  if (!createFun) {
    return null
  }
  const [data, expire] = createFun()
  if (!data) {
    return null
  }
  setCache(key, data, expire)
  return data
}

export const setCache = (key: string, data: any, expire: number = 0) => {
  if (!expire || expire < 0) {
    expire = 1000 * 60 * 60 * 24 * 30
  }
  const dw: DataWrap = {
    data,
    expireAt: Date.now() + expire,
  }
  localStorage.setItem(buildKey(key), JSON.stringify(dw))
}

export const removeCache = (key: string) => {
  localStorage.removeItem(buildKey(key))
}
