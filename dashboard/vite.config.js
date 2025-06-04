import { defineConfig } from 'vite'
import path from 'path'

export default defineConfig({
  server: {
    port: 3001,
  },
  base: process.env.NODE_ENV === 'production' ? '/dashboard' : '',
  build: {
    outDir: '../static/dashboard',
    rollupOptions: {
      input: {
        index: path.resolve(__dirname, 'index.html'),
        login: path.resolve(__dirname, 'login.html'),
      },
    },
  },
})
