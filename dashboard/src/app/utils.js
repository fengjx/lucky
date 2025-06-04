export const urlJoin = (...paths) => {
  // 过滤掉空值
  const filtered = paths.filter(Boolean)
  if (filtered.length === 0) return ''

  // 处理协议部分（如 http:// 或 https://）
  let first = filtered[0]
  let protocol = ''
  const protocolMatch = first.match(/^(https?:)\/\//)
  if (protocolMatch) {
    protocol = protocolMatch[1] + '//'
    first = first.slice(protocol.length)
  }

  // 拼接路径，去除多余的 /
  const joined = [first, ...filtered.slice(1)]
    .map((part, idx) => {
      if (idx === 0) {
        // 首段去除末尾 /
        return part.replace(/\/+$/, '')
      }
      // 其他段去除首尾 /
      return part.replace(/^\/+|\/+$/g, '')
    })
    .filter(Boolean)
    .join('/')

  return protocol + joined
}
