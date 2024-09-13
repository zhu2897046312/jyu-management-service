<template>
    <div class="container">
      <div class="left">
        <table>
          <tr>
            <td class="label">姓名：</td>
            <td>{{ UserBasicInformation.name }}</td>
            <td class="label">曾用名：</td>
            <td>{{UserBasicInformation.old_name}}</td>
          </tr>
          <tr>
            <td class="label">性别：</td>
            <td>{{UserBasicInformation.sex}}</td>
            <td class="label">证件类型：</td>
            <td>{{UserBasicInformation.identification_type}}</td>
          </tr>
          <tr>
            <td class="label">证件号码：</td>
            <td>{{UserBasicInformation.identification_number}}</td>
            <td class="label">出生日期：</td>
            <td>{{UserBasicInformation.birthday}}</td>
          </tr>
          <tr>
            <td class="label">民族：</td>
            <td>{{UserBasicInformation.ethnic_group}}</td>
            <td class="label">籍贯：</td>
            <td>广东河源</td>
          </tr>
          <tr>
            <td class="label">政治面貌：</td>
            <td>群众</td>
            <td class="label">入学日期：</td>
            <td>2022-09-12</td>
          </tr>
        </table>
      </div>
      <div class="right">
        <table>
          <tr>
            <td class="label">照片1：</td>
            <td class="photo">入学前</td>
          </tr>
          <tr>
            <td class="label">照片2：</td>
            <td class="photo">入学后</td>
          </tr>
        </table>
      </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data(){
        return {
            UserBasicInformation:{
                account: "",
                name: "",
                sex: "",
                identification_number: "",
                birthday: "",
                ethnic_group: "",
                identification_type: "",
                old_name:"",
                political_outlook:"",
                enrollment_dates:"",
            }
        }
    },
    mounted() {
        this.fetchUserBasicInformations();
    },
    methods: {
      fetchUserBasicInformations() {
        const account = this.$store.getters.getLoginData.account;  // 当前登录用户的学号
        
        axios({
            method: 'GET',
            url: 'http://localhost:8081/admin/GetUserInformation',
            params: {
                account: account  // 将学号作为查询参数
            },
            headers: {
                'Content-Type': 'application/json'
            }
        })
        .then(res => {
            console.log('成功', res.data);
            this.UserBasicInformation = res.data;
        })
        .catch(err => {
            console.error('失败', err);
        });
      }

    }
}
</script>
  
<style scoped>
.container {
display: flex;
justify-content: space-between;
padding: 20px;
}

.left, .right {
width: 48%;
}

table {
width: 100%;
border-collapse: collapse;
}

td {
padding: 10px;
border: 1px solid #ddd;
}

.label {
font-weight: bold;
width: 20%;
}

.photo {
text-align: center;
background-color: #f5f5f5;
padding: 20px;
border-radius: 8px;
}

.right .photo {
width: 100px;
height: 100px;
display: flex;
align-items: center;
justify-content: center;
background-color: #eee;
font-size: 14px;
}

.right tr {
margin-bottom: 20px;
}

.right td {
padding: 10px;
}

td {
vertical-align: middle;
}
</style>
