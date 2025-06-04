export const getEnv = () => {
  const env = {
    MODE: import.meta.env.MODE,
    BaseAPI: import.meta.env.VITE_BASE_API,
    LoginAPI: import.meta.env.VITE_LOGIN_API,
  }
  return env
}
