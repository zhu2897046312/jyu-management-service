<template>
    <div class="container">
      <form @submit.prevent="login">
        <div class="form-group">
          <label for="account">账号</label>
          <input 
            type="text" 
            id="account" 
            v-model="loginData.account" 
            placeholder="请输入账号" 
          />
          <!-- 更新account到Vuex -->
        </div>
        <div class="form-group">
          <label for="password">密码</label>
          <input 
            type="password" 
            id="password" 
            v-model="loginData.password" 
            placeholder="请输入密码" 
          />
          <!-- 更新password到Vuex -->
        </div>
        <button type="submit">登录</button>
      </form>
    </div> 
</template>

<script>

import axios from 'axios';
import { mapActions } from 'vuex';


export default {
  data() {
    return {
      loginData: {
        account: "",
        password: ""
      }
    };
  },
  methods: {
    ...mapActions(['updateAccount', 'updatePassword']),
    login() {
      // 发送登录请求
      // 调用 Vuex action 更新全局状态
      this.updateAccount(this.loginData.account);
      this.updatePassword(this.loginData.password);
      
      let submit = {
          "account": this.loginData.account,
          "password": this.loginData.password
      }
      axios({
        method: "POST",
        url: "http://localhost:8081/admin/Login",
        data: submit,
        headers: {
          'Content-Type': 'application/json'
        }
      }).then(res => {
        console.log(res.data); // 打印后端返回的数据
        const chat_type = res.data.data
        if (chat_type === 0) {
          this.$router.push({ path: '/admin' })
        } else if (chat_type === 1) {
          this.$router.push({ path: '/index' })
        }else{
          this.$router.push({ path: '/teacher' })
        }
        console.log("登录成功");
      }).catch(err => {
        console.log("登录失败");
      });
    },
  }
};
</script>

<style scoped>
/* 让容器高度占满整个屏幕，并使用Flexbox居中 */
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh; /* 使容器高度占满屏幕 */
  background-color: #f5f5f5; /* 可选：设置背景色 */
}

form {
  padding: 20px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 300px; /* 设置表单的宽度 */
}

.form-group {
  display: flex;
  margin-bottom: 15px;
}

label {
  margin-bottom: 5px;
}

input {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>
