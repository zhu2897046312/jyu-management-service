<template>
    <div class="common-layout">
      <el-container>
        <el-header>Header</el-header>
        <el-main>
          <el-table :data="courseInfo" @row-click="editStudentScores">
            <el-table-column prop="course_name" label="课程名称"></el-table-column>
            <el-table-column prop="class_name" label="教学班级"></el-table-column>
          </el-table>
        </el-main>
        <el-footer>Footer</el-footer>
      </el-container>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        courseInfo: [] // 用于存储从后端获取的课程信息
      };
    },
    methods: {
      fetchCourseInfo() {
        // 这里调用后端接口获取课程信息
        this.axios.get('/api/courseInfo').then(response => {
          this.courseInfo = response.courses;
        });
      },
      editStudentScores(course) {
        // 跳转到编辑成绩界面，传递课程信息
        this.$router.push({ name: 'EditScores', params: { courseCode: course.course_code } });
      }
    },
    mounted() {
      this.fetchCourseInfo();
    }
  };
  </script>
  
  <style scoped>
  .common-layout {
    padding: 20px;
  }
  </style>
  