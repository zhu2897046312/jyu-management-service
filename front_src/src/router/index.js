// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import Index from '@/views/Index.vue';  
import Login from '@/views/Login.vue'; 

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/index',
    name: 'Index',
    component: Index,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
