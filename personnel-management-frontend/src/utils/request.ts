
import {post,get} from "./server.js"

export const userLogin = data=>{
  return post({
    url:'/user/login',
    data
  })
}

export const userRegister = data=>{
  return post({
    url:'/user/register',
    data
  })
}