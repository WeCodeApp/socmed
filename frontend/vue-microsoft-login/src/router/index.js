import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import HomeView from '../views/HomeView.vue'
import AuthCallback from '../views/AuthCallback.vue'

const routes = [
  { path: '/home', name: 'Home', component: HomeView },
  { path: '/login', name: 'Login', component: LoginView },
  { path: '/auth/callback', name: 'AuthCallback', component: AuthCallback }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// router.beforeEach((to, from, next) => {
//   const isAuth = !!localStorage.getItem('auth_token')
//   if (to.name !== 'Login' && !isAuth) {
//     next({ name: 'Login' })
//   } else {
//     next()
//   }
// })

export default router
