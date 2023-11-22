import axios from 'axios'

const http = axios.create({
  baseURL: 'http://localhost:7701/api/'
})

window.$http = http
