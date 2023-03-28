import {createRouter,createWebHistory} from 'vue-router'

export default createRouter({
  history:createWebHistory(),
  routes: [
    {
      path:'/',
      redirect:'/index'
    },
    {
      // 登录/欢迎页
      name:'index',
      path:'/index',
      component:()=>import('@/pages/Index.vue')
    },
    {
      // 主页
      name:'home',
      path:'/home',
      component:()=>import('@/pages/Home.vue')
    },
    {
      name:'test',
      path:'/test',
      // component:()=>import('@/pages/Test.vue')
      component:()=>import('@/pages/Alert.vue')
    }
  ]
})

