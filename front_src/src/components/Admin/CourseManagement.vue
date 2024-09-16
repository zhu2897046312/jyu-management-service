<template>
    <div class="course-management">
      <el-button type="primary" @click="showAddCourseDialog">添加课程</el-button>
      <el-table :data="courses" stripe>
        <el-table-column prop="course_code" label="课程代码" width="150"></el-table-column>
        <el-table-column prop="course_name" label="课程名称" width="200"></el-table-column>
        <el-table-column prop="teacher_name" label="教师名称" width="150"></el-table-column>
        <el-table-column prop="credits" label="学分" width="100"></el-table-column>
        <el-table-column label="操作" width="180">
          <template #default="{ row }">
            <el-button size="mini" @click="editCourse(row)">编辑</el-button>
            <el-button size="mini" type="danger" @click="deleteCourse(row.course_code)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
  
      <el-pagination
        @current-change="handlePageChange"
        :current-page="currentPage"
        :page-size="pageSize"
        :total="totalCourses"
        layout="total, prev, pager, next, jumper"
        class="pagination"
      ></el-pagination>
  
      <el-dialog
        title="课程信息"
        v-model="dialogVisible"
        width="70%"
        @close="resetForm"
      >
        <el-form :model="form" ref="form" label-width="150px">
          <el-form-item label="课程代码" :rules="[{ required: true, message: '课程代码不能为空', trigger: 'blur' }]">
            <el-input v-model="form.course_code" :disabled="isEditing"></el-input>
          </el-form-item>
          <el-form-item label="课程名称" :rules="[{ required: true, message: '课程名称不能为空', trigger: 'blur' }]">
            <el-input v-model="form.course_name"></el-input>
          </el-form-item>
          <el-form-item label="教师名称" :rules="[{ required: true, message: '教师名称不能为空', trigger: 'blur' }]">
            <el-input v-model="form.teacher_name"></el-input>
          </el-form-item>
          <el-form-item label="学分" :rules="[{ required: true, message: '学分不能为空', trigger: 'blur' }]">
            <el-input type="number" v-model.number="form.credits"></el-input>
          </el-form-item>
          <el-form-item label="年级">
            <el-input v-model="form.academic_year"></el-input>
          </el-form-item>
          <el-form-item label="学期">
            <el-input type="number" v-model.number="form.semester"></el-input>
          </el-form-item>
          <el-form-item label="开课学院">
            <el-input v-model="form.commencement_academy"></el-input>
          </el-form-item>
          <el-form-item label="课程归属">
            <el-input v-model="form.course_affiliation"></el-input>
          </el-form-item>
          <el-form-item label="课程类别">
            <el-input type="number" v-model.number="form.course_type"></el-input>
          </el-form-item>
          <el-form-item label="课程性质">
            <el-input type="number" v-model.number="form.course_nature"></el-input>
          </el-form-item>
          <el-form-item label="教学班名称">
            <el-input v-model="form.class_name"></el-input>
          </el-form-item>
          <el-form-item label="上课时间">
            <el-input v-model="form.class_time"></el-input>
          </el-form-item>
          <el-form-item label="上课地点">
            <el-input v-model="form.class_address"></el-input>
          </el-form-item>
          <el-form-item label="人数">
            <el-input type="number" v-model.number="form.max_student_number"></el-input>
          </el-form-item>
          <el-form-item label="教学模式">
            <el-input type="number" v-model.number="form.teaching_mode"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm">提交</el-button>
            <el-button @click="cancelForm">取消</el-button>
          </el-form-item>
        </el-form>
      </el-dialog>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        courses: [],
        dialogVisible: false,
        form: {
          course_code: '',
          course_name: '',
          teacher_name: '',
          credits: null,
          academic_year: '',
          semester: null,
          commencement_academy: '',
          course_affiliation: '',
          course_type: null,
          course_nature: null,
          class_name: '',
          class_time: '',
          class_address: '',
          max_student_number: null,
          teaching_mode: null,
        },
        isEditing: false,
        currentPage: 1,  // 当前页码
        pageSize: 10,    // 每页条数
        totalCourses: 0, // 总课程数
      };
    },
    methods: {
      async fetchCourses(page = 1) {
        try {
          const response = await axios.get(`http://localhost:8081/admin/GetCourses`, {
            params: {
              page: page,
              pageSize: this.pageSize,
            },
          });
          this.courses = response.data.courses;
          this.totalCourses = response.data.total;
        } catch (error) {
          console.error('Failed to fetch courses:', error);
        }
      },
      showAddCourseDialog() {
        this.isEditing = false;
        this.form = {
          course_code: '',
          course_name: '',
          teacher_name: '',
          credits: null,
          academic_year: '',
          semester: null,
          commencement_academy: '',
          course_affiliation: '',
          course_type: null,
          course_nature: null,
          class_name: '',
          class_time: '',
          class_address: '',
          max_student_number: null,
          teaching_mode: null,
        };
        this.dialogVisible = true;
      },
      editCourse(course) {
        this.isEditing = true;
        this.form = { ...course };
        this.dialogVisible = true;
      },
      async deleteCourse(course_code) {
        try {
          await axios.delete(`http://localhost:8081/admin/DeleteCourses/${course_code}`);
          this.fetchCourses(this.currentPage);
        } catch (error) {
          console.error('Failed to delete course:', error);
        }
      },
      async submitForm() {
          try {
              const url = this.isEditing 
              ? `http://localhost:8081/admin/UpdateCourses/${this.form.course_code}` 
              : `http://localhost:8081/admin/AddCourse`;
              const method = this.isEditing ? 'PUT' : 'POST';
  
              await axios({ 
                  method,
                  url,
                  data: this.form
              });
              
              this.fetchCourses(this.currentPage);  // 刷新课程列表
              this.dialogVisible = false;
          } catch (error) {
              console.error('Failed to submit form:', error);
          }
      },
      cancelForm() {
        this.dialogVisible = false;
      },
      handlePageChange(page) {
        this.currentPage = page;
        this.fetchCourses(page);
      },
    },
    mounted() {
      this.fetchCourses();
    },
  };
  </script>
  
  <style scoped>
  .course-management {
    padding: 20px;
  }
  
  .el-table {
    margin-bottom: 20px;
  }
  
  .pagination {
    margin-top: 20px;
  }
  </style>
  