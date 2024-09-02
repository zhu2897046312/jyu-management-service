<template>
    <el-card>
      <h3>成绩</h3>
      <ul>
        
      </ul>
    </el-card>
  </template>
  
  <script>
  export default {
    data() {
      return {
        grades: [
          { course: '操作系统', score: 85 },
          { course: '软件开发技术', score: 93 },
          { course: '计算机网络', score: 77 }
        ]
      };
    },
    computed: {
      formattedGrades() {
        return this.grades.map(grade => `${grade.course}: ${grade.score} 分`);
      }
    }
  };
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
  