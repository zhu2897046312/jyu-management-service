<template>
  <div class="grades-card">
    <button>成绩</button>
    <button>考试</button>
    <ul>
      <li v-for="(grade, index) in grades" :key="index">
        <a href="#">
          <span class="course">{{ grade.course }}</span>
          <span class="score">{{ grade.score }} 分</span>
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      grades: [
        {course: "操作系统", score: "87",}
      ]
    };
  },
  // mounted() {
  //   this.fetchGrades();
  // },
  // methods: {
  //   async fetchGrades() {
  //     try {
  //       const response = await axios.get('/api/grades');
  //       this.grades = response.data;
  //     } catch (error) {
  //       console.error('获取成绩数据失败:', error);
  //     }
  //   }
  // }
};
</script>

<style scoped>
.grades-card {
  padding: 10px;
  background-color: white;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  border-radius: 5px;
}

h3 {
  font-size: 16px;
  margin-bottom: 10px;
}

ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

li {
  display: flex;
  justify-content: space-between;
  padding: 5px 0;
  font-size: 14px;
  border-bottom: 1px solid #eaeaea;
}

.course {
  font-weight: bold;
}

.score {
  color: #409eff;
}
</style>
