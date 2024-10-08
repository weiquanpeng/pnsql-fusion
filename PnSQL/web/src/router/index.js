import {createRouter, createWebHashHistory, useRoute} from 'vue-router'
import NProgress from 'nprogress'
import {useCounterStore} from "@/stores/user.js";
import 'nprogress/nprogress.css'
// import { menuTreeData } from '@/mock/data.js'
import businessRoutes from '@/router/businessRoutes';
NProgress.configure({showSpinner: false})

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login.vue')
  },
  {
    path: '/:catchAll(.*)',
    meta: {closeTab: true},
    component: () => import('@/views/404.vue')
  },
  {
    path: '/index',
    name: 'index',
    component: () => import('@/views/index.vue'),
    redirect: '/index/dashboard',
    children: businessRoutes // 将业务路由作为 children 加入
  },
]


const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach(async (to, from, next) => {
  NProgress.start()
  const useStore = useCounterStore()
  if (!useStore.token && to.name !== 'login') {
    next({name: 'login'})
  } else if (useStore.token && to.name === 'login') {
    next({name: 'dashboard'})
  } else {
    next()
  }
})


router.afterEach(()=>{
  NProgress.done()
})

export default router
