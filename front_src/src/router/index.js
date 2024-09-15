// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import Index from '@/views/Index.vue';  
import Login from '@/views/Login.vue'; 
import EnrollCourse from '@/views/EnrollCourse.vue'; 
import PrintfCourse from '@/views/PrintfCourse.vue'; 
import UserInformation from '@/views/UserInformation.vue';
import PrintGrades from '@/views/PrintGrades.vue'
import GradeInfo from '../components/UserInformation/GradeInfo.vue';  
const routes = [
  {
    path: '/',
    redirect: '/login',  // 访问 '/' 时重定向到 '/login'
  },
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
  {
    path: '/enrollCourse',
    name: 'EnrollCourse',
    component: EnrollCourse,
  },
  {
    path: '/printfCourse',
    name: 'PrintfCourse',
    component: PrintfCourse,
  },
  {
    path: '/userInformation',
    name: 'UserInformation',
    component: UserInformation,
  },
  {
    path: '/gradeInformation',
    name: 'gradeInformation',
    component: GradeInfo,
  },
  {
    path: '/printGrades',
    name: 'PrintGrades',
    component: PrintGrades,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
