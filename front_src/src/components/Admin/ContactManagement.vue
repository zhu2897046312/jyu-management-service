<template>
    <div class="contact-management">
      <el-button type="primary" @click="showAddContactDialog">添加联系方式</el-button>
      <el-table :data="contacts" stripe>
        <el-table-column prop="account" label="学号" width="150"></el-table-column>
        <el-table-column prop="phone" label="手机号码" width="150"></el-table-column>
        <el-table-column prop="email" label="电子邮箱" width="200"></el-table-column>
        <el-table-column prop="correspondence_address" label="通讯地址" width="200"></el-table-column>
        <el-table-column prop="landline" label="固定电话" width="200"></el-table-column>
        <el-table-column prop="home_address" label="家庭地址" width="200"></el-table-column>
        <el-table-column prop="post_code" label="邮政编码" width="200"></el-table-column>
        <el-table-column label="操作" width="180">
          <template #default="{ row }">
            <el-button size="mini" @click="editContact(row)">编辑</el-button>
            <el-button size="mini" type="danger" @click="deleteContact(row.account)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
  
      <el-dialog title="联系方式信息" v-model="dialogVisible" width="50%" @close="resetForm">
        <el-form :model="form" ref="form" label-width="120px">
          <el-form-item label="学号" :rules="[{ required: true, message: '学号不能为空', trigger: 'blur' }]">
            <el-input v-model="form.account" :disabled="isEditing"></el-input>
          </el-form-item>
          <el-form-item label="通讯地址">
            <el-input v-model="form.correspondence_address"></el-input>
          </el-form-item>
          <el-form-item label="手机号码">
            <el-input v-model="form.phone"></el-input>
          </el-form-item>
          <el-form-item label="电子邮箱">
            <el-input v-model="form.email"></el-input>
          </el-form-item>
          <el-form-item label="固定电话">
            <el-input v-model="form.landline"></el-input>
          </el-form-item>
          <el-form-item label="邮政编码">
            <el-input v-model="form.post_code"></el-input>
          </el-form-item>
          <el-form-item label="家庭地址">
            <el-input v-model="form.home_address"></el-input>
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
    contacts: [],
    dialogVisible: false,
    form: {
        account: '',
        correspondence_address: '',
        phone: '',
        email: '',
        landline: '',
        post_code: '',
        home_address: '',
    },
    isEditing: false,
    };
},
methods: {
    async fetchContacts() {
    try {
        const response = await axios.get('http://localhost:8081/admin/GetContacts');
        this.contacts = response.data.Contacts;
    } catch (error) {
        console.error('Failed to fetch contacts:', error);
    }
    },
    showAddContactDialog() {
    this.isEditing = false;
    this.form = {
        account: '',
        correspondence_address: '',
        phone: '',
        email: '',
        landline: '',
        post_code: '',
        home_address: '',
    };
    this.dialogVisible = true;
    },
    editContact(contact) {
    this.isEditing = true;
    this.form = { ...contact };
    this.dialogVisible = true;
    },
    async deleteContact(account) {
    try {
        await axios.delete(`http://localhost:8081/admin/DeleteContacts/${account}`);
        this.fetchContacts();
    } catch (error) {
        console.error('Failed to delete contact:', error);
    }
    },
    async submitForm() {
    const url = this.isEditing
        ? `http://localhost:8081/admin/UpdateContacts/${this.form.account}`
        : 'http://localhost:8081/admin/AddContacts';
    const method = this.isEditing ? 'PUT' : 'POST';

    try {
        await axios({ method, url, data: this.form });
        this.fetchContacts();
        this.dialogVisible = false;
    } catch (error) {
        console.error('Failed to submit form:', error);
    }
    },
    cancelForm() {
    this.dialogVisible = false;
    },
},
mounted() {
    this.fetchContacts();
},
};
</script>

<style scoped>
.contact-management {
padding: 20px;
}

.el-table {
margin-bottom: 20px;
}
</style>
  