<template>
    <div class="user-info-management">
      <el-button type="primary" @click="showAddDialog">添加用户基本信息</el-button>
      <el-table :data="users" stripe>
        <el-table-column prop="account" label="学号" width="120"></el-table-column>
        <el-table-column prop="name" label="姓名" width="150"></el-table-column>
        <el-table-column prop="old_name" label="曾用名" width="200"></el-table-column>
        <el-table-column prop="sex" label="性别" width="100"></el-table-column>
        <el-table-column prop="identification_type" label="证件类型" width="200"></el-table-column>
        <el-table-column prop="identification_number" label="身份证号" width="200"></el-table-column>
        <el-table-column prop="birthday" label="出生日期" width="200"></el-table-column>
        <el-table-column prop="ethnic_group" label="民族" width="200"></el-table-column>
        <el-table-column prop="political_outlook" label="政治面貌" width="200"></el-table-column>
        <el-table-column prop="enrollment_dates" label="入学日期" width="200"></el-table-column>
        <el-table-column label="操作" width="180">
          <template #default="{ row }">
            <el-button size="mini" @click="editUser(row)">编辑</el-button>
            <el-button size="mini" type="danger" @click="deleteUser(row.account)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
  
      <el-dialog title="用户基本信息" v-model="dialogVisible" width="50%">
        <el-form :model="form" ref="form" label-width="120px">
          <el-form-item label="学号">
            <el-input v-model="form.account" :disabled="isEditing"></el-input>
          </el-form-item>
          <el-form-item label="姓名">
            <el-input v-model="form.name"></el-input>
          </el-form-item>
          <el-form-item label="性别">
            <el-input v-model.number="form.sex"></el-input>
          </el-form-item>
          <el-form-item label="身份证号">
            <el-input v-model="form.identification_number"></el-input>
          </el-form-item>
          <!-- 其他字段 -->
          <el-form-item label="曾用名">
            <el-input v-model="form.old_name"></el-input>
          </el-form-item>
          <el-form-item label="证件类型">
            <el-input v-model="form.identification_type"></el-input>
          </el-form-item>
          <el-form-item label="政治面貌">
            <el-input v-model="form.political_outlook"></el-input>
          </el-form-item>
          <el-form-item label="出生日期">
            <el-input v-model="form.birthday"></el-input>
          </el-form-item>
          <el-form-item label="民族">
            <el-input v-model="form.ethnic_group"></el-input>
          </el-form-item>
          <el-form-item label="入学日期">
            <el-input v-model="form.enrollment_dates"></el-input>
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
      users: [],
      dialogVisible: false,
      form: {
        account: '',
        name: '',
        sex: '',
        identification_number: '',
        birthday: '',
        ethnic_group: '',
        identification_type: '',
        old_name: '',
        political_outlook: '',
        enrollment_dates: ''
      },
      isEditing: false
    };
  },
  methods: {
    async fetchUsers() {
      try {
        const response = await axios.get('http://localhost:8081/admin/GetUserBasicInfo');
        this.users = response.data.UserBasic;
      } catch (error) {
        console.error('Failed to fetch users:', error);
      }
    },
    showAddDialog() {
      this.isEditing = false;
      this.form = {
        account: '',
        name: '',
        sex: '',
        identification_number: '',
        birthday: '',
        ethnic_group: '',
        identification_type: '',
        old_name: '',
        political_outlook: '',
        enrollment_dates: ''
      };
      this.dialogVisible = true;
    },
    editUser(user) {
      this.isEditing = true;
      this.form = { ...user };
      this.dialogVisible = true;
    },
    async deleteUser(account) {
      try {
        await axios.delete(`http://localhost:8081/admin/DeleteUserBasicInfo/${account}`);
        this.fetchUsers();
      } catch (error) {
        console.error('Failed to delete user:', error);
      }
    },
    async submitForm() {
      try {
        const method = this.isEditing ? 'put' : 'post';
        const url = this.isEditing
          ? `http://localhost:8081/admin/UpdateUserBasicInfo/${this.form.account}`
          : 'http://localhost:8081/admin/AddUserBasicInfo';
        await axios({ method, url, data: this.form });
        this.dialogVisible = false;
        this.fetchUsers();
      } catch (error) {
        console.error('Failed to submit form:', error);
      }
    },
    cancelForm() {
      this.dialogVisible = false;
    }
  },
  mounted() {
    this.fetchUsers();
  }
};
</script>

<style scoped>
.user-info-management {
padding: 20px;
}

.el-table {
margin-bottom: 20px;
}
</style>