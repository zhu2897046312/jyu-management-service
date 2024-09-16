// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import Index from '@/views/Index.vue';  
import Admin from '@/views/Administrator.vue';
import CourseManagement from '@/components/Admin/CourseManagement.vue';
import BasicInfoManagement from '@/components/Admin/BasicInfoManagement.vue';
import StudentStatusManagement from '@/components/Admin/StudentStatusManagement.vue';
import ContactManagement from '@/components/Admin/ContactManagement.vue';

import Teacher from '@/views/TeacherIndex.vue';
import Login from '@/views/Login.vue'; 
import EnrollCourse from '@/views/EnrollCourse.vue'; 
import PrintfCourse from '@/views/PrintfCourse.vue'; 
import UserInformation from '@/views/UserInformation.vue';
import PrintGrades from '@/views/PrintGrades.vue';
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
    path: '/admin',
    name: 'Admin',
    component: Admin,
    children: [
      {
        path: 'course-management',
        name: 'CourseManagement',
        component: CourseManagement
      },
      {
        path: 'basicInfo-management',
        name: 'BasicInfoManagement',
        component: BasicInfoManagement,
      },
      {
        path: 'student-status-management',
        name: 'StudentStatusManagement',
        component: StudentStatusManagement,
      },
      {
        path: 'contact-management',
        name: 'ContactManagement',
        component: ContactManagement,
      },
    ]
  },
  {
    path: '/teacher',
    name: 'Teacher',
    component: Teacher,
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
