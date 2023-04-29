import { personnel } from "@/types/Personnel.js"
import { post, get, postJson, deleteReq } from "./server.js"
import { UserData } from "@/types/UserData"

export const userLogin = (data: UserData) => {
  return post({
    url: '/user/login',
    data
  }).then((res:any)=>{
    window.$message.success("登录成功")
    sessionStorage.setItem("userName",res.data.userName)
    localStorage.setItem("token",res.data.token)
  }).catch(()=>Promise.resolve())
}

export const userRegister = (data: UserData) => {
  return post({
    url: '/user/register',
    data
  }).then(()=>{
    window.$message.success("注册成功，请登录")
  }).catch(()=>Promise.resolve())
}

export const userInfo = () => {
  return get({
    url: '/user/info',
  })
}

export const getPersonnel = () =>{
  return get({
    url:'/personnel/getPersonnel',
  }).catch(()=>Promise.resolve())
}

export const createPersonnel = (data: personnel) => {
  return postJson({
    url: '/personnel/create',
    data
  }).catch(()=>Promise.resolve())
}

export const getImage = ()=>{
  return get({
    url:'/personnel/getImage',
  })
}

export const adminInfo = () => {
  return get({
    url: '/admin/isAdmin',
  })
}

export const a_getListData = ()=>{
  return get({
    url:'/admin/listData',
  })
}

export const a_getPersonnel = (userName:string)=>{
  return get({
    url:'/admin/personnelData',
    params:{
      userName
    }
  }).catch(()=>Promise.resolve())
}

export const a_createPersonnel = (data: personnel,userName:string) => {
  return postJson({
    url: '/admin/createPersonnel',
    params:{
      userName
    },
    data
  }).catch(()=>Promise.resolve())
}

export const a_getImage = (userName:string)=>{
  return get({
    url:'/admin/getImage',
    params:{
      userName
    }
  })
}

export const a_deletePersonnel = (userName:string)=>{
  return deleteReq({
    url:'/admin/deletePersonnel',
    params:{
      userName
    }
  })
}
