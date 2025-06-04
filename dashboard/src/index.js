'use strict'

import './amis'
import './assets/css/index.css'
import { USER_INFO_KEY, setCache } from './assets/lib/cache'
import { getEnv, fetchMenu, fetchUserInfo, pages, urlJoin } from './app'

// 填充完整 schemaApi
const fillSchemaApi = (menus) => {
  const env = getEnv()
  menus.forEach((menu) => {
    if (menu.children) {
      fillSchemaApi(menu.children)
    }
    if (menu.schemaApi && !menu.schemaApi.startsWith('http')) {
      menu.schemaApi = urlJoin(env.BaseAPI, menu.schemaApi)
    }
  })
  return menus
}

const loadUserInfo = async () => {
  const userInfo = await fetchUserInfo()
  if (!userInfo) {
    return null
  }
  setCache(USER_INFO_KEY, userInfo)
  return userInfo
}

;(async () => {
  const env = getEnv()
  if (!getToken()) {
    window.location = pages.login
    return
  }
  const userInfo = await loadUserInfo()
  if (!userInfo) {
    window.location = pages.login
    return
  }

  // 从服务器拉取菜单
  const data = await fetchMenu()
  const menus = fillSchemaApi(data.pages || [])

  const app = {
    type: 'app',
    data: {
      userInfo,
      env,
    },
    brandName: 'lucky',
    logo: './logo.png',
    header: {
      type: 'grid',
      columns: [
        {
          md: 0,
          body: [],
        },
        {
          md: 9,
          body: [
            {
              type: 'flex',
              justify: 'flex-end',
              className: 'header-right',
              items: [
                {
                  type: 'dropdown-button',
                  label: '源码',
                  trigger: 'hover',
                  icon: 'fa fa-github',
                  buttons: [
                    {
                      type: 'button',
                      label: 'lucky',
                      actionType: 'url',
                      url: 'https://github.com/fengjx/lucky',
                    },
                    {
                      type: 'button',
                      label: 'lucky-web',
                      actionType: 'url',
                      url: 'https://github.com/fengjx/lucky-web',
                    },
                  ],
                },
                {
                  type: 'dropdown-button',
                  label: '${userInfo.nickname}',
                  trigger: 'hover',
                  icon: 'fa fa-user',
                  buttons: [
                    {
                      type: 'button',
                      label: '文档',
                    },
                    {
                      label: '退出登录',
                      type: 'button',
                      actionType: 'dialog',
                      dialog: {
                        title: '确认退出登录？',
                        onEvent: {
                          confirm: {
                            actions: [
                              {
                                actionType: 'custom',
                                script: 'logout()',
                              },
                            ],
                          },
                        },
                      },
                    },
                  ],
                },
              ],
            },
          ],
        },
      ],
    },
    pages: menus,
  }

  let amisInstance = createAmis('#root', app, {})

  appHistory.listen((state) => {
    amisInstance.updateProps({
      location: state.location || state,
    })
  })
})()
