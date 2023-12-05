import axios from 'axios'

export const getAxiosInstance = () => {
  const instance = axios.create({
    baseURL: location.origin.replace('7702', '7701') + '/api/',
    headers: {
      Authorization: localStorage.vortexnotes_passcode
        ? 'Bearer ' + localStorage.vortexnotes_passcode
        : undefined
    }
  })

  instance.interceptors.response.use(
    response => response,
    error => {
      if (error.response) {
        if (error.response.status === 401) {
          localStorage.removeItem('vortexnotes_passcode')
          localStorage.removeItem('vortexnotes_auth_scope')
          location.href = '/auth'
        }
      }
      return Promise.reject(error)
    }
  )

  return instance
}

window.$http = getAxiosInstance()
