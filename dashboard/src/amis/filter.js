import { getDictLabel, getDictValue } from '../app'
const amisLib = amisRequire('amis')

amisLib.registerFilter('getDictLabel', (group, value) => {
  return getDictLabel(group, value)
})

amisLib.registerFilter('getDictValue', (group, label) => {
  return getDictValue(group, label)
})
