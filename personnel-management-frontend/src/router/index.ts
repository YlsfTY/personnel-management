import { userInfo } from '@/utils/request'
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
      redirect: '/home/personnel',
      children: [
        // {
          // name: 'message',
          // path: 'message',
          // component: () => import('@/components/homeMessage.vue'),
          // meta: {
            // requiresAuth: true
          // }
        // },
        {
          name: 'personnel',
          path: 'personnel',
          component: () => import('@/components/Personnel.vue'),
          meta: {
            requiresAuth: true
          }
        },
        {
          name: 'formList',
          path: 'formList',
          component: () => import('@/components/formList.vue'),
          meta: {
            requiresAuth: true
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
      // component:()=>import('@/pages/Alert.vue')
    }
  ]
})

const tokenGuard = (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
  // const isAuthenticated = to.meta.requiresAuth
  let isAuthenticated = to.matched.some(item => item.meta.requiresAuth)
  // console.log(`to.path: ${to.path}, isAuthenticated: ${isAuthenticated}`) 
  const token = localStorage.getItem('token')
  if (!isAuthenticated) {
    console.log(isAuthenticated);
    if (!token) {
      next()
      return
    }
    userInfo().then((res: any) => {
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
    userInfo().then((res: any) => {
      next()
    }).catch(() => {
      next('/')
      return Promise.resolve()
    })
    return
  }
}


router.beforeEach(tokenGuard)

export default router