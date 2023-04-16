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
      component:()=>import('@/pages/Home.vue'),
      redirect:'/home/message',
      children:[
        {
          name:'message',
          path:'message',
          component:()=>import('@/components/homeMessage.vue')
        },
        {
          name:'personnel',
          path:'personnel',
          component:()=>import('@/components/Personnel.vue')
        },
        {
          name:'formList',
          path:'formList',
          component:()=>import('@/components/formList.vue')
        }
      ]
    },
    {
      name:'test',
      path:'/test',
      component:()=>import('@/pages/Test.vue')
      // component:()=>import('@/pages/Alert.vue')
    }
  ]
})
