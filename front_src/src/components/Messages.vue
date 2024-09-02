<template>
    <el-card>
      <h3>消息</h3>
      <ul>
        
      </ul>
    </el-card>
  </template>
  
  <script>
  export default {
    data() {
      return {
        messages: [
          { time: '2024-05-29 09:14:49', content: '老师发布了新的作业通知' },
          { time: '2024-04-06 13:27:36', content: '作业截止时间即将到来' }
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
    font-size: 14px;
  }
  </style>
  