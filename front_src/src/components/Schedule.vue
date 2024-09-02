<template>
    <el-card>
      <h3>课程</h3>
      <ul>
        
      </ul>
    </el-card>
  </template>
  
  <script>
  export default {
    data() {
      return {
        courses: [
          { weekday: '周一', time: '7-8节', name: '软件开发技术', room: '北区607' },
          { weekday: '周二', time: '3-4节', name: '计算机网络', room: '北区307' },
          { weekday: '周五', time: '5-6节', name: '操作系统', room: '北区303' }
        ]
      }
    }
  }
  </script>
  
  <style scoped>
  ul {
    list-style: none;
    padding: 0;
  }
  
  li {
    margin-bottom: 10px;
  }
  </style>
  