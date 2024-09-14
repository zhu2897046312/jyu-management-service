<template>
    <div class="container">
      <div class="left">
        <table>
          <tr>
            <td class="label">年级：</td>
            <td>{{ StudentStatusInformation.grade }}</td>
            <td class="label">学院名称：</td>
            <td>{{StudentStatusInformation.academy_name}}</td>
            <td class="label">专业名称：</td>
            <td>{{StudentStatusInformation.professional_name}}</td>
          </tr>
          <tr>
            <td class="label">班级名称：</td>
            <td>{{ StudentStatusInformation.class_name }}</td>
            <td class="label">学制：</td>
            <td>{{StudentStatusInformation.academic}}</td>
            <td class="label">学籍状态：</td>
            <td>{{StudentStatusInformation.status}}</td>
          </tr>
          <tr>
            <td class="label">是否在校：</td>
            <td>{{ StudentStatusInformation.is_in_School }}</td>
            <td class="label">报道注册状态：</td>
            <td>{{StudentStatusInformation.registration_status}}</td>
            <td class="label">注册时间：</td>
            <td>{{StudentStatusInformation.registration_time}}</td>
          </tr>
          <tr>
            <td class="label">学历层次：</td>
            <td>{{ StudentStatusInformation.educational_level }}</td>
            <td class="label">培养方式：</td>
            <td>{{StudentStatusInformation.cultivation_method}}</td>
            <td class="label">培养层次：</td>
            <td>{{StudentStatusInformation.cultivation_level}}</td>
          </tr>
          <tr>
            <td class="label">学生类别：</td>
            <td>{{ StudentStatusInformation.student_type }}</td>
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
          StudentStatusInformation:{
                account: "",
                grade: "",
                academy_name: "",
                class_name: "",
                professional_name: "",
                academic: "",
                status: "",
                is_in_School: "",
                registration_status:"",
                educational_level:"",
                cultivation_method:"",
                cultivation_level:"",
                student_type:"",
                check_in_time:"",
                registration_time:"",
            }
        }
    },
    mounted() {
        this.fetchStudentStatusInformations();
    },
    methods: {
      fetchStudentStatusInformations() {
        const account = this.$store.getters.getLoginData.account;  // 当前登录用户的学号
        
        axios({
            method: 'GET',
            url: 'http://localhost:8081/admin/GetStudentStatusInformation',
            params: {
                account: account  // 将学号作为查询参数
            },
            headers: {
                'Content-Type': 'application/json'
            }
        })
        .then(res => {
            console.log('成功', res.data);
            this.StudentStatusInformation = res.data;
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
