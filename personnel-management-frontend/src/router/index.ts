import { adminInfo, userInfo } from '@/utils/request'
import { NavigationGuardNext, RouteLocationNormalized, createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/index',
      meta: {
        requiresAuth: false
      }
    },
    {
      // 登录/欢迎页
      name: 'index',
      path: '/index',
      component: () => import('@/pages/Index.vue'),
      meta: {
        requiresAuth: false
      },
    },
    {
      // 主页
      name: 'home',
      path: '/home',
      component: () => import('@/pages/Home.vue'),
      // redirect:`home/personnel/${sessionStorage.getItem('userName')}`,
      redirect: to => {
        return {
          name: 'personnel',
          params: {
            userName: sessionStorage.getItem('userName')
          },
        }
      },
      children: [
        {
          name: 'personnel',
          path: 'personnel/:userName',
          props:true,
          component: () => import('@/components/Personnel.vue'),
          meta: {
            requiresAuth: true
          }
        },
        {
          name: 'formList',
          path: 'formList/:userName',
          props:true,
          component: () => import('@/components/formList.vue'),
          meta: {
            requiresAuth: true
          }
        },
        {
          name:'userList',
          path:'userList',
          component:()=>import('@/components/userList.vue'),
          meta:{
            requiresAuth:true,
            adminAuth:true
          }
        }
      ],
      meta: {
        requiresAuth: true
      }
    },
    {
      name: 'test',
      path: '/test',
      component: () => import('@/pages/Test.vue'),
      meta: {
        requiresAuth: false
      }
    }
  ]
})

const tokenGuard = (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
  let isAuthenticated = to.matched.some(item => item.meta.requiresAuth)
  let isAdminAuth = to.matched.some(item => item.meta.adminAuth)
  const token = localStorage.getItem('token')
  if (!isAuthenticated) {
    if (!token) {
      next()
      return
    }
    userInfo().then((res: any) => {
      sessionStorage.setItem("userName",res.data.user)
      next('/home')
    }).catch(() => {
      next()
      return Promise.resolve()
    })
    return
  }else{
    if (!token) {
      next('/')
      return
    }
    if(isAdminAuth){
      adminInfo().then((res:any)=>{
        sessionStorage.setItem("userName",res.data.user)
        next()
      }).catch(()=>{
        next('home')
        return Promise.resolve()
      })
      return
    }else{
      userInfo().then((res:any)=>{
        sessionStorage.setItem("userName",res.data.user)
        next()
      }).catch(()=>{
        next('/')
        return Promise.resolve()
      })
      return
    }
  }
}

router.beforeEach(tokenGuard)

export default router