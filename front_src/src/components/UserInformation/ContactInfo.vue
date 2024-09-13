<template>
    <div class="container">
      <div class="left">
        <table>
          <tr>
            <td class="label">电子邮箱：</td>
            <td>{{ ContactInformation.email }}</td>
            <td class="label">手机号码：</td>
            <td>{{ContactInformation.phone}}</td>
          </tr>
          <tr>
            <td class="label">固定电话：</td>
            <td>{{ContactInformation.landline}}</td>
            <td class="label">家庭地址：</td>
            <td>{{ContactInformation.home_address}}</td>
          </tr>
          <tr>
            <td class="label">通信地址：</td>
            <td>{{ContactInformation.ccorrespondence_address}}</td>
            <td class="label">邮政编码：</td>
            <td>{{ContactInformation.post_code}}</td>
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
            ContactInformation:{
                account: "",
                ccorrespondence_address: "",
                phone: "",
                email: "",
                landline: "",
                post_code: "",
                home_address:"",
            }
        }
    },
    mounted() {
        this.fetchContactInformations();
    },
    methods: {
      fetchContactInformations() {
        const account = this.$store.getters.getLoginData.account;  // 当前登录用户的学号
        
        axios({
            method: 'GET',
            url: 'http://localhost:8081/admin/GetContactInformation',
            params: {
                account: account  // 将学号作为查询参数
            },
            headers: {
                'Content-Type': 'application/json'
            }
        })
        .then(res => {
            console.log('成功', res.data);
            this.ContactInformation = res.data;
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
