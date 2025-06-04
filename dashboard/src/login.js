'use strict'

import './amis'
import './assets/css/login.css'
import {  getEnv, getToken, setToken, pages } from './app'
;(() => {
  if (getToken()) {
    window.location = pages.index
  }
  const env = getEnv()
  const amisJSON = {
    type: 'page',
    data: {
      env,
    },
    body: {
      type: 'flex',
      alignItems: 'center',
      style: {
        height: '100vh',
        backgroundColor: '#eaecef',
      },
      items: [
        {
          type: 'form',
          style: {
            width: '480px',
          },
          title: `路辰后台系统登录`,
          redirect: pages.index,
          mode: 'horizontal',
          horizontal: {
            left: 'col-sm-2',
            right: 'col-sm-10',
            offset: 'col-sm-offset-2',
          },
          api: {
            method: 'post',
            url: `${env.LoginAPI}`,
          },
          onFinished: (values) => {
            if (values.token) {
              setToken(values.token)
              window.location = pages.index
            }
            return false // 这样可以禁掉 amis 后续的默认行为
          },
          body: [
            {
              label: '用户名',
              type: 'input-text',
              name: 'username',
              required: true,
            },
            {
              label: '密码',
              type: 'input-password',
              name: 'password',
              required: true,
            },
          ],
          actions: [
            {
              type: 'submit',
              label: '登录',
              level: 'primary',
            },
          ],
        },
      ],
    },
  }
  let amisScoped = createAmis('#root', amisJSON, {})
})()
