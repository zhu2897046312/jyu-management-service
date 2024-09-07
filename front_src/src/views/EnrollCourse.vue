<template>
  <div class="container">
    <!-- 选课列表 -->
    <ul class="enroll">
      <li v-for="course in courses" :key="course.course_code">
        <span>{{ course.course_name }}</span>
        <button @click="toggleDropdown(course.course_name)">
          <i :class="['dropdown-icon', { 'icon-up': activeDropdown === course.course_name, 'icon-down': activeDropdown !== course.course_name }]"></i>
        </button>
        <!-- 可选：展示课程详细信息的区域 -->
        <div v-if="activeDropdown === course.course_name" class="course-details">
          <span><strong>Course Code:</strong> {{ course.course_code }}</span>
          <span><strong>Academic Year:</strong> {{ course.academic_year }}</span>
          <span><strong>Semester:</strong> {{ course.semester }}</span>
          <span><strong>Commencement Academy:</strong> {{ course.commencement_academy }}</span>
          <span><strong>Course Type:</strong> {{ course.course_type }}</span>
          <span><strong>Course Nature:</strong> {{ course.course_nature }}</span>
          <span><strong>Credits:</strong> {{ course.credits }}</span>
          <span><strong>Class Name:</strong> {{ course.class_name }}</span>
          <span><strong>Teacher Name:</strong> {{ course.teacher_name }}</span>
          <span><strong>Class Time:</strong> {{ course.class_time }}</span>
          <span><strong>Class Address:</strong> {{ course.class_address }}</span>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      activeDropdown: null,
      courses: []
    };
  },
  mounted() {
    this.fetchCourses();
  },
  methods: {
    async fetchCourses() {
      try {
        const response = await axios.get('http://localhost:8081/admin/GetAll');
        this.courses = response.data;
      } catch (error) {
        console.error('获取课程数据失败:', error);
      }
    },
    toggleDropdown(courseName) {
      this.activeDropdown = this.activeDropdown === courseName ? null : courseName;
    }
  }
};
</script>

<style scoped>
/* 样式 */
.container {
  padding: 20px;
}

.enroll {
  list-style-type: none;
  padding: 0;
}

.enroll li {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.enroll span {
  flex: 1;
}

.dropdown-icon {
  cursor: pointer;
  transition: transform 0.3s ease;
}

.icon-up {
  transform: rotate(180deg);
}

.icon-down {
  transform: rotate(0deg);
}

.course-details {
  margin-top: 10px;
}
</style>
