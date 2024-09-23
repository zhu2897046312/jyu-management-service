<template>
  <div>
      <!-- 添加下载和上传按钮 -->
    <el-button type="primary" @click="downloadTemplate">下载Excel模板</el-button>
    <el-button type="primary" @click="download">下载已有的所有账户信息</el-button>
    <el-upload
      class="upload-demo"
      drag
      action="http://localhost:8081/admin/uploadExcel"
      :on-success="handleUploadSuccess"
      :show-file-list="false"
    >
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">拖拽文件或点击上传</div>
      <div class="el-upload__text">新用户由系统生成账号不用填account列/批量更新用户信息先下载execl文件再填写好要更新的内容</div>
      <div class="el-upload__tip" slot="tip">上传Excel文件</div>
    </el-upload>

    <el-button type="primary" @click="showAddDialog">添加账号</el-button>
    <el-table :data="userAccounts" stripe>
      <el-table-column prop="account" label="学号" width="150"></el-table-column>
      <el-table-column prop="password" label="密码" width="150"></el-table-column>
      <el-table-column prop="chat_type" label="账户类型" width="150"></el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="{ row }">
          <el-button size="mini" @click="editUserAccount(row)">编辑</el-button>
          <el-button size="mini" type="danger" @click="deleteUserAccount(row.account)">删除</el-button>
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

    <el-dialog title="账号信息" v-model="dialogVisible" width="50%">
      <el-form :model="form" ref="form" label-width="120px">
        <!-- 账号在编辑时显示，新增时由后端生成 -->
        <el-form-item label="学号" v-if="isEditing">
          <el-input v-model="form.account" disabled></el-input>
        </el-form-item>

        <el-form-item label="密码">
          <el-input v-model="form.password" type="password"></el-input>
        </el-form-item>

        <el-form-item label="确认密码">
          <el-input v-model="form.confirmPassword" type="password"></el-input>
        </el-form-item>

        <el-form-item label="账户类型">
          <el-select v-model="form.chat_type">
            <el-option label="学生" :value="1"></el-option>
            <el-option label="教师" :value="2"></el-option>
            <el-option label="管理员" :value="0"></el-option>
          </el-select>
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
    userAccounts: [],
    dialogVisible: false,
    form: {
        account: '',
        password: '',
        confirmPassword: '', // 增加确认密码字段
        chat_type: ''
    },
    isEditing: false,
    currentPage: 1,
    pageSize: 10,
    total: 0
    };
},
methods: {
    async fetchUserAccounts(page = 1) {
      try {
          const response = await axios.get('http://localhost:8081/admin/GetUserAccount', {
          params: {
              page: page,
              pageSize: this.pageSize
          }
          });
          this.userAccounts = response.data.UserAccount;
          this.total = response.data.total;
      } catch (error) {
          console.error('Failed to fetch user accounts:', error);
      }
    },
    showAddDialog() {
      this.isEditing = false;
      this.form = {
          account: '',
          password: '',
          confirmPassword: '',
          chat_type: ''
      };
      this.dialogVisible = true;
    },
    editUserAccount(row) {
      this.isEditing = true;
      this.form = { ...row, confirmPassword: row.password }; // 确认密码字段
      this.dialogVisible = true;
    },
    async deleteUserAccount(account) {
      try {
          await axios.delete(`http://localhost:8081/admin/DeleteUserAccount/${account}`);
          await axios.delete(`http://localhost:8081/admin/DeleteUserBasicInfo/${account}`);
          await axios.delete(`http://localhost:8081/admin/DeleteContacts/${account}`);
          await axios.delete(`http://localhost:8081/admin/DeleteStudentStatusInfo/${account}`);
          this.fetchUserAccounts(this.currentPage);
      } catch (error) {
          console.error('Failed to delete user account:', error);
      }
    },
    async submitForm() {
      if (this.form.password !== this.form.confirmPassword) {
          this.$message.error('密码和确认密码不一致');
          return;
      }

      try {
          const url = this.isEditing
          ? `http://localhost:8081/admin/UpdateUserAccount/${this.form.account}`
          : 'http://localhost:8081/admin/AddUserAccount';
          const method = this.isEditing ? 'PUT' : 'POST';

          await axios({
          method,
          url,
          data: this.form
          });

          this.fetchUserAccounts(this.currentPage);
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
      this.fetchUserAccounts(page);
    },
    // 下载Excel模板
    downloadTemplate() {
      axios({
        url: 'http://localhost:8081/generate_excel_template', // 后端生成模板的接口
        method: 'GET',
        responseType: 'blob' // 确保返回的是文件流
      }).then((response) => {
        const blob = new Blob([response.data], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' });
        const link = document.createElement('a');
        link.href = URL.createObjectURL(blob);
        link.download = 'user_account_template.xlsx'; // 下载的文件名
        link.click();
      }).catch((error) => {
        console.error('下载模板失败', error);
      });
    },
    handleUploadSuccess(response) {
      // 确保正确处理二进制数据
      console.log(response.message);
    },
    download() {
      axios({
        url: 'http://localhost:8081/all_execl', // 后端生成模板的接口
        method: 'GET',
        responseType: 'blob' // 确保返回的是文件流
      }).then((response) => {
        const blob = new Blob([response.data], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' });
        const link = document.createElement('a');
        link.href = URL.createObjectURL(blob);
        link.download = 'all.xlsx'; // 下载的文件名
        link.click();
      }).catch((error) => {
        console.error('下载模板失败', error);
      });
    }
},
mounted() {
    this.fetchUserAccounts();
}
};
</script>

<style scoped>
.el-table {
margin-bottom: 20px;
}
</style>
  