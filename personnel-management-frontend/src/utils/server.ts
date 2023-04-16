// import axios,{AxiosError} from "axios"
import axios, { AxiosRequestConfig } from "axios"

const Server = axios.create({
  timeout: 30000,
  baseURL: 'http://127.0.0.1:8000/api',
  headers: {
    "Content-Type": "application/x-www-form-urlencoded;charset=utf-8",
  }
})

// 请求拦截
Server.interceptors.request.use(config => {
  // window.$message.loading("连接中")
  return config
})

// 响应拦截
Server.interceptors.response.use(response => {
  const code = response.status
  if (code >= 200 && code < 300) {
    return response.data
  } else {
    // window.$message.error('服务器异常')
    throw new Error(`请求出错，状态码：${code}`)
  }
}, error => {
  const { response } = error
  if (response) {
    const { status, data } = response
    // window.$message.error('服务器异常')
    throw new Error(`请求出错，状态码：${status}，错误信息：${data.msg}`)
  } else {
    window.$message.error('服务器未连接')
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