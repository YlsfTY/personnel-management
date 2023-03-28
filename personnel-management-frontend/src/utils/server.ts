// import axios,{AxiosError} from "axios"
import axios, { AxiosRequestConfig } from "axios"
import {useMessage} from 'naive-ui'

const Server = axios.create({
  timeout: 30000,
  // baseURL: '/api/',
  baseURL: 'http://127.0.0.1:8080/api',
  headers: {
    "Content-Type": "application/x-www-form-urlencoded;charset=utf-8",
  }
})

// 请求拦截
Server.interceptors.request.use(config => {
  useMessage().loading('连接中')
  return config
})

// 响应拦截
Server.interceptors.response.use(response => {
  const code = response.status
  if (code >= 200 && code < 300){
    return response.data
  } else {
    throw new Error(`请求出错，状态码：${code}`)
  }
}, error => {
  const { response } = error
  if (response) {
    const { status, data } = response
    throw new Error(`请求出错，状态码：${status}，错误信息：${data}`)
  } else {
    throw new Error('请求出错，请检查网络连接')
  }
})

// post请求
export const post = (config: AxiosRequestConfig<any>) => {
  return Server({
    ...config,
    method: "post",
    data: config.data,
  })
}

// get请求
export const get = (config: AxiosRequestConfig<any>) => {
  return Server({
    ...config,
    method: "get",
    data: config.data,
  })
}