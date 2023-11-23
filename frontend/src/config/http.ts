import axios from 'axios'

const http = axios.create({
  baseURL: location.origin.replace('7702', '7701') + '/api/'
})

window.$http = http
