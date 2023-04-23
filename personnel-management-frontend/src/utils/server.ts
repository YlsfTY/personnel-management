// import axios,{AxiosError} from "axios"
// import router from "@/router"
import axios, { AxiosRequestConfig } from "axios"

const Server = axios.create({
  timeout: 30000,
  baseURL: 'http://127.0.0.1:8000/api',
})

// 请求拦截
Server.interceptors.request.use(config => {
  if (localStorage.getItem("token")) {
    config.headers.Authorization = `Bearer ${localStorage.getItem("token")}`
  }
  return config
}, error => {
  return
})

// 响应拦截
Server.interceptors.response.use(response => {
  const code = response.status
  if (code >= 200 && code < 300) {
    return response.data
  }
}, (e) => {
  throw handleResErr(e)
})

// post请求
export const post = (config: AxiosRequestConfig<any>) => {
  config.headers = config.headers || {}
  config.headers["Content-Type"] = "application/x-www-form-urlencoded;charset=utf-8"
  return Server({
    ...config,
    method: "post",
    data: config.data,
  })
}

// get请求
export const get = (config: AxiosRequestConfig<any>) => {
  config.headers = config.headers || {}
  config.headers["Content-Type"] = "application/x-www-form-urlencoded;charset=utf-8"
  return Server({
    ...config,
    method: "get",
    data: config.data,
  })
}

export const postJson = (config: AxiosRequestConfig<any>) => {
  config.headers = config.headers || {}
  config.headers["Content-Type"] = "application/json;charset=utf-8"
  return Server({
    ...config,
    method: "post",
    data: config.data,
  })
}

const handleResErr = (err: any) => {
  // console.log('aaaaaa'+err);
  if (err.response) {
    const status = err.response.status
    const msg = err.response.data.msg
    switch (status) {
      case 500:
        switch (msg) {
          case "Failed to get personnel":
            // window.$message.warning("获取人员信息失败")
            return Promise.resolve()
            break;
          default:
            window.$message.warning("服务器异常");
            break;
        }
        break;
      case 400:
        switch (msg) {
          case "Account resolution failed":
            window.$message.warning("账户名解析失败")
            break;
          case "Password resolution failed":
            window.$message.warning("密码解析失败")
            break;
          case "Failed to parse JSON":
          case "Failed to validate JSON":
            window.$message.warning("格式解析失败")
            break;
          default:
            window.$message.error("请求参数有误，请检查输入")
            break;
        }
        break;
      case 401:
        switch (msg) {
          case "Unauthorized":
            // window.$message.warning("未授权")
            break;
          case "Token Parsing Failed":
            // window.$message.warning("Token解析失败")
            // router.push("/")
            break;
          default:
            window.$message.error("请求出错，请稍后重试")
          break;
        }
        break;
      case 404:
        switch(msg){
          case "Image not found":
            // window.$message.warning("图片未找到")
          break;
          default:
            window.$message.error("请求出错，请稍后重试")
          break;
        }
        break;
      case 409:
        switch (msg) {
          case "The user already exists.":
            window.$message.warning("用户已存在")
            break;
          case "The user does not exist.":
            window.$message.warning("用户不存在")
            break;
          default:
            window.$message.error("请求冲突，请稍后重试")
            break;
        }
        break;
      default:
        window.$message.error("请求出错，请稍后重试")
        break;
    }
  } else if (err.request) {
    window.$message.error('服务器未连接')
  } else {
    window.$message.error('未知错误')
  }
  // return Promise.resolve() 
  return Promise.reject(err)
}