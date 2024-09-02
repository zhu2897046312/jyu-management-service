<template>
    <div>
      <h3>通知</h3>
      <ul>
        
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        notifications: [
          { type: '重要', content: '关于假期的通知' },
          { type: '提醒', content: '下周上交作业' },
          { type: '通知', content: '新学期课表已公布' }
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
  