import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/views/Login'
import Homepage from '@/views/Homepage'

Vue.use(Router)

const router = new Router({
  mode:'history',
  routes: [
    {
      path: '/',
      name: Login,
      component:Login
    },
    {
      path: '/Homepage',
      name: Homepage,
      component:Homepage
    }
  ]
})

router.beforeEach((to, from, next) => {
  if(to.path === '/'){
    return next();
  }
  const tokenStr = window.sessionStorage.getItem('token');
  if(!tokenStr) {
    return next('/');
  }
  next();
})

export default router
