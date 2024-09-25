<template>
    <div class="edit-scores">
      <el-form :model="form" ref="formRef" label-width="120px">
        <el-form-item label="课程名称">
          <el-input v-model="form.courseName" disabled></el-input>
        </el-form-item>
        <el-form-item label="教学班级">
          <el-input v-model="form.className" disabled></el-input>
        </el-form-item>
        <el-form-item label="选择学生">
          <el-select v-model="selectedStudent" @change="loadStudentScores" placeholder="请选择学生">
            <el-option
              v-for="student in students"
              :key="student.id"
              :label="student.name"
              :value="student.id"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="成绩">
          <el-input v-model="form.score" type="number"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submit">保存</el-button>
          <el-button @click="cancel">取消</el-button>
        </el-form-item>
      </el-form>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    props: {
      course: Object,
    },
    data() {
      return {
        form: {
          courseName: this.course.course_name,
          className: this.course.class_name,
          score: null,
        },
        students: [],
        selectedStudent: null,
      };
    },
    mounted() {
      this.loadStudents();
    },
    methods: {
      async loadStudents() {
        try {
          const response = await axios.get('http://localhost:8081/teacher/GetStudentsByCourse', {
            params: { courseId: this.course.id },
          });
          this.students = response.data.students || [];
        } catch (error) {
          console.error('获取学生数据失败:', error);
        }
      },
      async loadStudentScores() {
        if (this.selectedStudent) {
          try {
            const response = await axios.get('http://localhost:8081/teacher/GetStudentScore', {
              params: { studentId: this.selectedStudent, courseId: this.course.id },
            });
            this.form.score = response.data.score || null;
          } catch (error) {
            console.error('获取学生成绩失败:', error);
          }
        }
      },
      async submit() {
        try {
          const response = await axios.post('http://localhost:8081/teacher/UpdateStudentScore', {
            courseId: this.course.id,
            studentId: this.selectedStudent,
            score: this.form.score,
          });
          if (response.data.success) {
            this.$message.success('成绩更新成功');
            this.$emit('refresh');
          } else {
            this.$message.error('成绩更新失败');
          }
        } catch (error) {
          console.error('更新成绩时发生错误:', error);
          this.$message.error('更新成绩时发生错误');
        }
      },
      cancel() {
        this.$emit('close');
      },
    },
  };
  </script>
  
  <style scoped>
  .edit-scores {
    padding: 20px;
  }
  </style>
  