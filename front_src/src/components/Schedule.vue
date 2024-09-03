<template>
  <div class="courses-card">
    <h3>课程</h3>
    <ul>
      <li v-for="(course, index) in courses" :key="index">
        <span class="weekday">{{ course.weekday }}</span>
        <span class="time">{{ course.time }}</span>
        <span class="name">{{ course.name }}</span>
        <span class="room">{{ course.room }}</span>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      courses: []
    };
  },
  mounted() {
    this.fetchCourses();
  },
  methods: {
    async fetchCourses() {
      try {
        const response = await axios.get('/api/courses');
        this.courses = response.data;
      } catch (error) {
        console.error('获取课程数据失败:', error);
      }
    }
  }
};
</script>

<style scoped>
.courses-card {
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

.weekday {
  flex-basis: 10%;
  font-weight: bold;
}

.time {
  flex-basis: 15%;
  color: #409eff;
}

.name {
  flex-basis: 45%;
}

.room {
  flex-basis: 20%;
  text-align: right;
  color: #888;
}
</style>
