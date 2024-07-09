import { createRouter, createWebHistory } from 'vue-router'
import Pastries from '../views/PastriesView.vue'
import Orders from '../views/OrdersView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/pastries'
    },
    {
      path: '/pastries',
      name: 'pastries',
      component: Pastries
    },
    {
      path: '/orders',
      name: 'orders',
      component: Orders
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    }
  ]
})

export default router
