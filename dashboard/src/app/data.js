import request from '../assets/lib/request'
import { DICT_KEY, getCache, setCache } from '../assets/lib/cache'

export const fetchInitData = async () => {
  return await request({
    url: '/api/open/app/data',
    method: 'GET',
    data: {},
  })
}

export const getDictList = (group) => {
  const dictMap = getCache(DICT_KEY)
  if (!dictMap) {
    return null
  }
  return dictMap[group]
}

export const getDictValue = (group, label) => {
  const dictList = getDictList(group)
  if (!dictList) {
    return ''
  }
  for (const dict of dictList) {
    if (dict.label == label) {
      return dict.value
    }
  }
  return ''
}

export const getDictLabel = (group, value) => {
  const dictList = getDictList(group)
  if (!dictList) {
    return ''
  }
  for (const dict of dictList) {
    if (dict.value == value) {
      return dict.label
    }
  }
  return ''
}

export const getDictOptions = () => {
  const dict = getCache(DICT_KEY)
  if (!dict) {
    return []
  }
  const options = {}
  Object.keys(dict).forEach((key) => {
    const items = dict[key]
    const opts = []
    for (const item of items) {
      opts.push({
        label: item.label,
        value: item.value,
      })
    }
    options[key] = opts
  })
  return options
}
