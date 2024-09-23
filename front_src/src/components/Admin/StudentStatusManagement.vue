<template>
    <div>
      <el-button type="primary" @click="showAddDialog">添加学籍信息</el-button>
      <el-table :data="studentStatuses" stripe>
        <el-table-column prop="account" label="学号" width="150"></el-table-column>
        <el-table-column prop="user_name" label="姓名" width="150"></el-table-column>
        <el-table-column prop="academic_year" label="年级" width="150"></el-table-column>
        <el-table-column prop="academy_name" label="学院名称" width="150"></el-table-column>
        <el-table-column prop="class_name" label="班级名称" width="150"></el-table-column>
        <el-table-column prop="status" label="学籍状态" width="150"></el-table-column>
        <el-table-column label="操作" width="180">
          <template #default="{ row }">
            <el-button size="mini" @click="editStudentStatus(row)">编辑</el-button>
            <el-button size="mini" type="danger" @click="deleteStudentStatus(row.account)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
  
      <el-pagination
        @current-change="handlePageChange"
        :current-page="currentPage"
        :page-size="pageSize"
        :total="total"
        layout="total, prev, pager, next, jumper"
      ></el-pagination>
  
      <el-dialog title="学籍信息" v-model="dialogVisible" width="50%">
        <el-form :model="form" ref="form" label-width="120px">
          <el-form-item label="学号">
            <el-input v-model="form.account" :disabled="isEditing"></el-input>
          </el-form-item>
          <el-form-item label="年级">
            <el-input v-model="form.academic_year"></el-input>
          </el-form-item>
          <el-form-item label="学院名称">
            <el-input v-model="form.academy_name"></el-input>
          </el-form-item>
          <el-form-item label="班级名称">
            <el-input v-model="form.class_name"></el-input>
          </el-form-item>
          <el-form-item label="专业名称">
            <el-input v-model="form.professional_name"></el-input>
          </el-form-item>
          <el-form-item label="学籍状态">
            <el-input v-model="form.status"></el-input>
          </el-form-item>
          <el-form-item label="是否在校">
            <el-input v-model.number="form.is_in_School"></el-input>
          </el-form-item>
          <el-form-item label="报到注册状态">
            <el-input v-model="form.registration_status"></el-input>
          </el-form-item>
          <el-form-item label="学历层次">
            <el-input v-model="form.educational_level"></el-input>
          </el-form-item>
          <el-form-item label="培养方式">
            <el-input v-model="form.cultivation_method"></el-input>
          </el-form-item>
          <el-form-item label="培养层次">
            <el-input v-model.number="form.cultivation_level"></el-input>
          </el-form-item>
          <el-form-item label="学生类别">
            <el-input v-model.number="form.student_type"></el-input>
          </el-form-item>
          <el-form-item label="报到时间">
            <el-input v-model="form.check_in_time"></el-input>
          </el-form-item>
          <el-form-item label="注册时间">
            <el-input v-model="form.registration_time"></el-input>
          </el-form-item>
          <el-form-item label="学制">
            <el-input v-model.number="form.academic"></el-input>
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
    studentStatuses: [],
    dialogVisible: false,
    form: {
        account: '',
        academic_year: '',
        academy_name: '',
        class_name: '',
        professional_name:'',
        status: '',
        is_in_School:'',
        registration_status:'',
        educational_level:'',
        cultivation_method:'',
        cultivation_level:'',
        student_type:'',
        check_in_time:'',
        registration_time:'',
        academic:'',
        user_name:'',
    },
    isEditing: false,
    currentPage: 1,
    pageSize: 10,
    total: 0,
    };
},
methods: {
    async fetchStudentStatuses(page = 1) {
    try {
        const response = await axios.get('http://localhost:8081/admin/GetStudentStatusInfo', {
        params: {
            page: page,
            pageSize: this.pageSize,
        },
        });
        this.studentStatuses = response.data.StudentStatusInfo;
        this.total = response.data.total;
    } catch (error) {
        console.error('Failed to fetch student statuses:', error);
    }
    },
    showAddDialog() {
    this.isEditing = false;
    this.form = {
        account: '',
        academic_year: '',
        academy_name: '',
        class_name: '',
        status: ''
    };
    this.dialogVisible = true;
    },
    editStudentStatus(row) {
    this.isEditing = true;
    this.form = { ...row };
    this.dialogVisible = true;
    },
    async deleteStudentStatus(account) {
    try {
        await axios.delete(`http://localhost:8081/admin/DeleteStudentStatusInfo/${account}`);
        this.fetchStudentStatuses(this.currentPage);
    } catch (error) {
        console.error('Failed to delete student status:', error);
    }
    },
    async submitForm() {
    try {
        const url = this.isEditing 
        ? `http://localhost:8081/admin/UpdateStudentStatusInfo/${this.form.account}` 
        : 'http://localhost:8081/admin/AddStudentStatusInfo';
        const method = this.isEditing ? 'PUT' : 'POST';

        await axios({
        method,
        url,
        data: this.form
        });

        this.fetchStudentStatuses(this.currentPage);
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
    this.fetchStudentStatuses(page);
    },
},
mounted() {
    this.fetchStudentStatuses();
},
};
</script>

<style scoped>
.el-table {
margin-bottom: 20px;
}
</style>
  