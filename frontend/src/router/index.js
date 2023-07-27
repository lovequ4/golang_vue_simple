import { createRouter,createWebHistory } from "vue-router";

import Home from '../views/Home.vue'
import Order from '../views/Order.vue'
import Payment from '../views/Payment.vue'

const routes = [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path:"/order/:orderId",
      name: 'order',
      component: Order,
    },
    {
      path:"/pay/:orderId",
      name:'pay',
      component: Payment,
    },
  ];


const router = createRouter({
    history: createWebHistory(),
    routes,
  });
  
export default router;