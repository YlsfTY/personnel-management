import { personnel } from "@/types/Personnel.js"
import { post, get } from "./server.js"
import { UserData } from "@/types/UserData"

export const userLogin = (data: UserData) => {
  return post({
    url: '/user/login',
    data
  })
}

export const userRegister = (data: UserData) => {
  return post({
    url: '/user/register',
    data
  })
}

export const getPersonnel = (userName:string) =>{
  return get({
    url:'/personnel/information/'+userName,
  })
}