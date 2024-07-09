import { createRouter, createWebHistory } from 'vue-router'
import Pastries from '../views/PastriesView.vue'
import Orders from '../views/OrdersView.vue'
import Schedules from '../views/SchedulesView.vue'

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
      path: '/schedules',
      name: 'schedules',
      component: Schedules
    },
  ]
})

export default router
